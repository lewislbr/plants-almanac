import React from 'react';
import {Link} from 'react-router-dom';

import {PlantCard} from '../components';
import {useGetPlantsQuery} from '../graphql/types';

export function Plants(): JSX.Element {
  const {data, loading, error, refetch} = useGetPlantsQuery();

  React.useEffect(() => {
    refetch();
  }, [refetch]);

  return (
    <>
      <section>
        <h1 className="page-title">{'Plants'}</h1>
      </section>
      <section className="mt-8">
        {loading ? (
          <div>{'Loading...'}</div>
        ) : error ? (
          <p>{'ERROR'}</p>
        ) : (
          <div>
            {data?.getPlants?.map(plant => (
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
