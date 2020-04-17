import {gql} from '@apollo/client';

export const GET_PLANT = gql`
  query GetPlant($name: String!) {
    getPlant(name: $name) {
      _id
      name
      otherNames
      description
      plantSeason
      harvestSeason
      pruneSeason
      tips
    }
  }
`;
