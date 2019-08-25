import { buildSchema } from 'graphql';

export default buildSchema(`
type Plant {
  _id: ID!
  name: String!
  description: String
  plantSeason: [String!]
  harvestSeason: [String!]
  pruneSeason: [String!]
  tips: String
}

input PlantInput {
  name: String!
  description: String
  plantSeason: [String!]
  harvestSeason: [String!]
  pruneSeason: [String!]
  tips: String
}

type RootQuery {
  plants: [Plant!]!
}

type RootMutation {
  createPlant(input: PlantInput): Plant
}

schema {
  query: RootQuery
  mutation: RootMutation
}
`);
