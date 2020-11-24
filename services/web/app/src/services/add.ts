import {gql} from "@apollo/client"
import {client, history} from "../index"

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

export async function addOne(plant: Record<string, unknown>): Promise<void> {
  if (!plant.name) {
    throw new Error("Name is required.")
  }

  const plantDTO = {
    name: plant.name,
    other_names: plant.otherNames || null,
    description: plant.description || null,
    plant_season: plant.plantSeason || null,
    harvest_season: plant.harvestSeason || null,
    prune_season: plant.pruneSeason || null,
    tips: plant.tips || null,
  }

  await client.mutate({
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
    variables: plantDTO,
  })

  history.push("/")
}
