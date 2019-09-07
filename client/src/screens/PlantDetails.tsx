import React from 'react';
import { useQuery } from '@apollo/react-hooks';
import gql from 'graphql-tag';

import { H1, Loading, Section } from '../components';

const GET_PLANT = gql`
  query GetPlant($name: String!) {
    getPlant(name: $name) {
      _id
      name
      description
      plantSeason
      harvestSeason
      pruneSeason
      tips
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
          <Loading />
        ) : error ? (
          <p>ERROR</p>
        ) : (
          <>
            <H1>{data.getPlant.name}</H1>
          </>
        )}
      </Section>
      <Section>
        {loading ? (
          <Loading />
        ) : error ? (
          <p>ERROR</p>
        ) : (
          <>
            <h3>Plant ID:</h3>
            <p>{data.getPlant._id}</p>
            <h3>Description:</h3>
            <p>
              {data.getPlant.description == null
                ? 'No data yet'
                : data.getPlant.description}
            </p>
            <h3>Plant Season:</h3>
            <p>
              {data.getPlant.plantSeason == null
                ? 'No data yet'
                : data.getPlant.plantSeason}
            </p>
            <h3>Harvest Season:</h3>
            <p>
              {data.getPlant.harvestSeason == null
                ? 'No data yet'
                : data.getPlant.harvestSeason}
            </p>
            <h3>Prune Season</h3>
            <p>
              {data.getPlant.pruneSeason == null
                ? 'No data yet'
                : data.getPlant.pruneSeason}
            </p>
            <h3>Tips:</h3>
            <p>
              {data.getPlant.tips == null ? 'No data yet' : data.getPlant.tips}
            </p>
          </>
        )}
      </Section>
    </>
  );
};
