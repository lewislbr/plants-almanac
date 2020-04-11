import React from 'react';
import {Link} from 'react-router-dom';
import {useQuery} from '@apollo/react-hooks';
import gql from 'graphql-tag';

import {PlantCard} from '../components';
import {Plant} from '../types/Plant';

const GET_PLANTS = gql`
  query GetPlants {
    getPlants {
      _id
      name
    }
  }
`;

export function Home(): JSX.Element {
  const {data, loading, error} = useQuery(GET_PLANTS);

  return (
    <>
      <section>
        <h1 className="font-bold">{'Home'}</h1>
      </section>
      <section>
        {loading ? (
          <div>{'Loading...'}</div>
        ) : error ? (
          <p>{'ERROR'}</p>
        ) : (
          <div>
            {data.getPlants.map((plant: Plant) => (
              <Link to={`/${plant.name}`} key={plant._id}>
                <PlantCard name={plant.name} />
              </Link>
            ))}
          </div>
        )}
      </section>
    </>
  );
}
