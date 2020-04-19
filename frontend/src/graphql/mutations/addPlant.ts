import {gql} from '@apollo/client';

export const ADD_PLANT = gql`
  mutation AddPlant(
    $_id: ID!
    $name: String!
    $otherNames: String
    $description: String
    $plantSeason: String
    $harvestSeason: String
    $pruneSeason: String
    $tips: String
  ) {
    addPlant(
      _id: $_id
      name: $name
      otherNames: $otherNames
      description: $description
      plantSeason: $plantSeason
      harvestSeason: $harvestSeason
      pruneSeason: $pruneSeason
      tips: $tips
    ) {
      insertedId
    }
  }
`;