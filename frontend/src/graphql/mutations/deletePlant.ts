import {gql} from '@apollo/client';

export const DELETE_PLANT = gql`
  mutation DeletePlant($_id: ID!) {
    deletePlant(_id: $_id) {
      deletedCount
    }
  }
`;
