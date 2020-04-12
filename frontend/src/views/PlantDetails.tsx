import React from 'react';

import {Alert} from '../components';
import {useGetPlantQuery} from '../graphql/queries/getPlant.graphql';
import {useDeletePlantMutation} from '../graphql/mutations/deletePlant.graphql';

export function PlantDetails(props: {history: any; match: any}): JSX.Element {
  const {data, loading, error} = useGetPlantQuery({
    variables: {name: props.match.params.plantname},
  });
  const [deletePlantMutation] = useDeletePlantMutation();

  const [alertOpen, setAlertOpen] = React.useState(false);

  function openAlert(): void {
    setAlertOpen(true);
  }

  async function deletePlant(event: React.SyntheticEvent): Promise<void> {
    event.preventDefault();
    await deletePlantMutation({
      variables: {_id: data?.getPlant?._id as string},
    });
    props.history.push('/');
  }

  return (
    <>
      {loading ? (
        <div>{'Loading...'}</div>
      ) : error ? (
        <p>{'ERROR'}</p>
      ) : (
        <>
          <section>
            <h1 className="page-title">{data?.getPlant?.name}</h1>
          </section>
          <section className="mb-12">
            <h5 className="data-title">{'Other Names:'}</h5>
            <p className="data-body">
              {data?.getPlant?.otherNames || 'No data yet'}
            </p>
            <h5 className="data-title">{'Description:'}</h5>
            <p className="data-body">
              {data?.getPlant?.description || 'No data yet'}
            </p>
            <h5 className="data-title">{'Plant Season:'}</h5>
            <p className="data-body">
              {data?.getPlant?.plantSeason || 'No data yet'}
            </p>
            <h5 className="data-title">{'Harvest Season:'}</h5>
            <p className="data-body">
              {data?.getPlant?.harvestSeason || 'No data yet'}
            </p>
            <h5 className="data-title">{'Prune Season:'}</h5>
            <p className="data-body">
              {data?.getPlant?.pruneSeason || 'No data yet'}
            </p>
            <h5 className="data-title">{'Tips:'}</h5>
            <p className="data-body">{data?.getPlant?.tips || 'No data yet'}</p>
          </section>
          <div className="flex justify-center">
            <button
              className="button button-danger"
              type="button"
              onClick={openAlert}
            >
              {'Delete plant'}
            </button>
          </div>
          {alertOpen ? (
            <Alert deletePlant={deletePlant} setAlertOpen={setAlertOpen} />
          ) : null}
        </>
      )}
    </>
  );
}
