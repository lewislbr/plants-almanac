import React from 'react';
import { Link } from 'react-router-dom';
import { useQuery } from '@apollo/react-hooks';
import gql from 'graphql-tag';

import { H1, Loading, PlantCard, Section } from '../components';
import { Plant } from '../utils/plantInterface';

const GET_PLANTS = gql`
  query GetPlants {
    getPlants {
      _id
      name
    }
  }
`;

export const Home: React.FunctionComponent = () => {
  const { data, loading, error } = useQuery(GET_PLANTS);

  return (
    <>
      <Section>
        <H1>Home</H1>
      </Section>
      <Section>
        {loading ? (
          <Loading />
        ) : error ? (
          <p>ERROR</p>
        ) : (
          <div>
            {data.getPlants.map((plant: Plant) => (
              <Link to={`/${plant.name}`} key={plant._id}>
                <PlantCard name={plant.name} />
              </Link>
            ))}
          </div>
        )}
      </Section>
    </>
  );
};
