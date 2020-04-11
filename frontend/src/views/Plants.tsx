import React from 'react';
import {Link} from 'react-router-dom';
import {gql, useQuery} from '@apollo/client';

import {PageTitle, PlantCard} from '../components';
import {Plant} from '../types/Plant';

const GET_PLANTS = gql`
  query GetPlants {
    getPlants {
      _id
      name
    }
  }
`;

export function Plants(): JSX.Element {
  const {data, loading, error, refetch} = useQuery(GET_PLANTS);

  React.useEffect(() => {
    refetch();
  }, [refetch]);

  return (
    <>
      <section>
        <PageTitle>{'Plants'}</PageTitle>
      </section>
      <section className="mt-8">
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
