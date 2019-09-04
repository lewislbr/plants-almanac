import React from 'react';
import { useQuery } from '@apollo/react-hooks';
import gql from 'graphql-tag';

import { H1, PlantCard, Section } from '../components';

const GET_PLANTS = gql`
  query GetPlants {
    getPlants {
      _id
      name
    }
  }
`;

interface Plant {
  _id: string;
  name: string;
  description: string;
  plantSeason: string[];
  harvestSeason: string[];
  pruneSeason: string[];
  tips: string;
}

export const Home: React.FunctionComponent = () => {
  const { data, loading, error } = useQuery(GET_PLANTS);

  return (
    <>
      <Section>
        <H1>Home</H1>
      </Section>
      <Section>
        {loading ? (
          <p>Loading...</p>
        ) : error ? (
          <p>ERROR</p>
        ) : (
          <div>
            {data.getPlants.map((plant: Plant) => (
              <PlantCard key={plant._id} name={plant.name} />
            ))}
          </div>
        )}
      </Section>
    </>
  );
};
