import {gql} from 'apollo-server';

export const typeDefs = gql`
  type Plant {
    _id: ID!
    name: String!
    otherNames: String
    description: String
    plantSeason: String
    harvestSeason: String
    pruneSeason: String
    tips: String
  }

  type Query {
    getPlants: [Plant!]
    getPlant(name: String!): Plant
  }

  type Mutation {
    addPlant(
      name: String!
      otherNames: String
      description: String
      plantSeason: String
      harvestSeason: String
      pruneSeason: String
      tips: String
    ): Plant!
    deletePlant(_id: ID!): Plant
  }
`;
