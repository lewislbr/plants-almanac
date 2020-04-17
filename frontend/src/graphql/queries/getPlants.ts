import {gql} from '@apollo/client';

export const GET_PLANTS = gql`
  query GetPlants {
    getPlants {
      _id
      name
    }
  }
`;
