import {
  ApolloClient,
  ApolloQueryResult,
  createHttpLink,
  gql,
  InMemoryCache,
} from "@apollo/client"
import {persistCache, LocalStorageWrapper} from "apollo3-cache-persist"
import {AddVariables} from "../interfaces/Add"
import {DeleteVariables} from "../interfaces/Delete"
import {EditVariables} from "../interfaces/Edit"
import {Plant} from "../interfaces/Plant"
import {Plants} from "../interfaces/Plants"

const cache = new InMemoryCache()

await persistCache({
  cache,
  storage: new LocalStorageWrapper(window.localStorage),
})

const plantsClient = new ApolloClient({
  cache,
  link: createHttpLink({
    credentials: "include",
    uri: process.env.PLANTS_URL,
  }),
})

export async function addOne(plant: Record<string, string>): Promise<void> {
  const ADD = gql`
    mutation Add(
      $name: String!
      $other_names: String
      $description: String
      $plant_season: String
      $harvest_season: String
      $prune_season: String
      $tips: String
    ) {
      add(
        name: $name
        other_names: $other_names
        description: $description
        plant_season: $plant_season
        harvest_season: $harvest_season
        prune_season: $prune_season
        tips: $tips
      )
    }
  `
  const newPlantDTO: AddVariables = {
    name: plant.name,
    other_names: plant.otherNames || null,
    description: plant.description || null,
    plant_season: plant.plantSeason || null,
    harvest_season: plant.harvestSeason || null,
    prune_season: plant.pruneSeason || null,
    tips: plant.tips || null,
  }

  await plantsClient.mutate({
    mutation: ADD,
    update(cache, {data: {add}}) {
      cache.modify({
        fields: {
          plants(existingPlantRefs = []): unknown[] {
            const newPlantRef = cache.writeQuery({
              data: add,
              query: ADD,
            })

            return [...existingPlantRefs, newPlantRef]
          },
        },
      })
    },
    variables: newPlantDTO,
  })
}

export async function listAll(): Promise<ApolloQueryResult<Plants>> {
  const PLANTS = gql`
    query Plants {
      plants {
        id
        created_at
        edited_at
        name
      }
    }
  `

  return plantsClient.query({query: PLANTS})
}

export async function listOne(id: string): Promise<ApolloQueryResult<Plant>> {
  const PLANT = gql`
    query Plant($id: ID!) {
      plant(id: $id) {
        id
        created_at
        edited_at
        name
        other_names
        description
        plant_season
        harvest_season
        prune_season
        tips
      }
    }
  `

  return plantsClient.query({query: PLANT, variables: {id: id}})
}

export async function editOne(
  id: EditVariables["id"],
  plant: Record<string, string>,
): Promise<void> {
  const EDIT = gql`
    mutation Edit(
      $id: ID!
      $name: String!
      $other_names: String
      $description: String
      $plant_season: String
      $harvest_season: String
      $prune_season: String
      $tips: String
    ) {
      edit(
        id: $id
        name: $name
        other_names: $other_names
        description: $description
        plant_season: $plant_season
        harvest_season: $harvest_season
        prune_season: $prune_season
        tips: $tips
      )
    }
  `
  const editedPlantDTO: EditVariables = {
    id: id,
    name: plant.name,
    other_names: plant.otherNames ?? null,
    description: plant.description ?? null,
    plant_season: plant.plantSeason ?? null,
    harvest_season: plant.harvestSeason ?? null,
    prune_season: plant.pruneSeason ?? null,
    tips: plant.tips ?? null,
  }

  await plantsClient.mutate({
    mutation: EDIT,
    update(cache, {data: {edit}}) {
      cache.modify({
        fields: {
          plant(): unknown {
            const newPlantRef = cache.writeQuery({
              data: edit,
              query: EDIT,
            })

            return newPlantRef
          },
          plants(existingPlantRefs = []): unknown[] {
            const newPlantRef = cache.writeQuery({
              data: edit,
              query: EDIT,
            })

            return [...existingPlantRefs, newPlantRef]
          },
        },
      })
    },
    variables: editedPlantDTO,
  })
}

export async function deleteOne(id: DeleteVariables["id"]): Promise<void> {
  const DELETE = gql`
    mutation Delete($id: ID!) {
      delete(id: $id)
    }
  `

  await plantsClient.mutate({
    mutation: DELETE,
    update(cache) {
      cache.modify({
        fields: {
          plants(_, {DELETE}): unknown {
            return DELETE
          },
        },
      })
    },
    variables: {id: id},
  })
}
