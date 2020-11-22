import {ApolloQueryResult, gql} from "@apollo/client"
import {client} from "../index"

const PLANTS = gql`
  query Plants {
    plants {
      id
      name
    }
  }
`

export async function listAll(): Promise<ApolloQueryResult<unknown>> {
  return await client.query({query: PLANTS})
}

const PLANT = gql`
  query Plant($id: ID!) {
    plant(id: $id) {
      id
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

export async function listOne(id: string): Promise<ApolloQueryResult<unknown>> {
  return await client.query({query: PLANT, variables: {id: id}})
}
