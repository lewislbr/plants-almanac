import React from 'react';

import {Button, DataFieldBody, DataFieldTitle, PageTitle} from '../components';
import {useGetPlantQuery} from '../graphql/queries/getPlant.graphql';
import {useDeletePlantMutation} from '../graphql/mutations/deletePlant.graphql';

export function PlantDetails(props: {history: any; match: any}): JSX.Element {
  const {data, loading, error} = useGetPlantQuery({
    variables: {name: props.match.params.plantname},
  });
  const [deletePlant] = useDeletePlantMutation();

  async function confirmDeletePlant(
    event: React.FormEvent<HTMLFormElement>,
  ): Promise<void> {
    event.preventDefault();
    await deletePlant({variables: {_id: data?.getPlant?._id as string}});
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
            <PageTitle>{data?.getPlant?.name}</PageTitle>
          </section>
          <section>
            <DataFieldTitle>{'Other Names:'}</DataFieldTitle>
            <DataFieldBody>
              {data?.getPlant?.otherNames || 'No data yet'}
            </DataFieldBody>
            <DataFieldTitle>{'Description:'}</DataFieldTitle>
            <DataFieldBody>
              {data?.getPlant?.description || 'No data yet'}
            </DataFieldBody>
            <DataFieldTitle>{'Plant Season:'}</DataFieldTitle>
            <DataFieldBody>
              {data?.getPlant?.plantSeason || 'No data yet'}
            </DataFieldBody>
            <DataFieldTitle>{'Harvest Season:'}</DataFieldTitle>
            <DataFieldBody>
              {data?.getPlant?.harvestSeason || 'No data yet'}
            </DataFieldBody>
            <DataFieldTitle>{'Prune Season:'}</DataFieldTitle>
            <DataFieldBody>
              {data?.getPlant?.pruneSeason || 'No data yet'}
            </DataFieldBody>
            <DataFieldTitle>{'Tips:'}</DataFieldTitle>
            <DataFieldBody>
              {data?.getPlant?.tips || 'No data yet'}
            </DataFieldBody>
          </section>
          <div className="flex justify-center">
            <Button style="danger" type="button" onClick={confirmDeletePlant}>
              {'Delete plant'}
            </Button>
          </div>
        </>
      )}
    </>
  );
}
