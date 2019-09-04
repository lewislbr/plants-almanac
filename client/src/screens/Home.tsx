import React from 'react';
import { useQuery } from '@apollo/react-hooks';
import gql from 'graphql-tag';

import { H1, PlantCard } from '../components';

const GET_PLANTS = gql`
  query {
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
  if (loading) return <p>Loading...</p>;
  if (error) return <p>ERROR</p>;

  return (
    <>
      <H1>Home</H1>
      {data.getPlants.map((plant: Plant) => (
        <PlantCard key={plant._id} name={plant.name} />
      ))}
    </>
  );
};
