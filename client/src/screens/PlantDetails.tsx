import React from 'react';
import { useQuery } from '@apollo/react-hooks';
import gql from 'graphql-tag';

import { H1, Section } from '../components';

const GET_PLANT = gql`
  query GetPlant($name: String!) {
    getPlant(name: $name) {
      _id
      name
    }
  }
`;

export const PlantDetails: React.FunctionComponent = () => {
  const dirtyPath = location.pathname;
  const cleanPath = dirtyPath.replace(/%20/g, ' ').replace(/\//g, '');

  const { data, loading, error } = useQuery(GET_PLANT, {
    variables: { name: cleanPath },
  });

  return (
    <>
      <Section>
        {loading ? (
          <p>Loading...</p>
        ) : error ? (
          <p>ERROR</p>
        ) : (
          <>
            <H1>{data.getPlant.name}</H1>
            <p>{data.getPlant._id}</p>
          </>
        )}
      </Section>
    </>
  );
};
