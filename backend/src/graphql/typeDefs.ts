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

  type AddPlantResponse {
    insertedId: ID!
  }

  type DeletePlantResponse {
    deletedCount: Int
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
    ): AddPlantResponse!
    deletePlant(_id: ID!): DeletePlantResponse
  }
`;
