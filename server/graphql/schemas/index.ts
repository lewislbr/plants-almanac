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

type Query {
  plants: [Plant!]!
}

type Mutation {
  createPlant(input: PlantInput): Plant
}
`);
