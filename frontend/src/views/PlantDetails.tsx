import React from 'react';
import {gql, useMutation, useQuery} from '@apollo/client';

import {Button, DataFieldBody, DataFieldTitle, PageTitle} from '../components';

const GET_PLANT = gql`
  query GetPlant($name: String!) {
    getPlant(name: $name) {
      _id
      name
      otherNames
      description
      plantSeason
      harvestSeason
      pruneSeason
      tips
    }
  }
`;

const DELETE_PLANT = gql`
  mutation DeletePlant($_id: ID!) {
    deletePlant(_id: $_id) {
      _id
      name
      otherNames
      description
      plantSeason
      harvestSeason
      pruneSeason
      tips
    }
  }
`;

export function PlantDetails(props: {history: any; match: any}): JSX.Element {
  const {data, loading, error} = useQuery(GET_PLANT, {
    variables: {name: props.match.params.plantname},
  });
  const [deletePlant] = useMutation(DELETE_PLANT);

  async function confirmDeletePlant(
    event: React.FormEvent<HTMLFormElement>,
  ): Promise<void> {
    event.preventDefault();
    await deletePlant({variables: {_id: data.getPlant._id}});
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
            <PageTitle>{data.getPlant.name}</PageTitle>
          </section>
          <section>
            <DataFieldTitle>{'Other Names:'}</DataFieldTitle>
            <DataFieldBody>
              {data.getPlant.otherNames || 'No data yet'}
            </DataFieldBody>
            <DataFieldTitle>{'Description:'}</DataFieldTitle>
            <DataFieldBody>
              {data.getPlant.description || 'No data yet'}
            </DataFieldBody>
            <DataFieldTitle>{'Plant Season:'}</DataFieldTitle>
            <DataFieldBody>
              {data.getPlant.plantSeason || 'No data yet'}
            </DataFieldBody>
            <DataFieldTitle>{'Harvest Season:'}</DataFieldTitle>
            <DataFieldBody>
              {data.getPlant.harvestSeason || 'No data yet'}
            </DataFieldBody>
            <DataFieldTitle>{'Prune Season:'}</DataFieldTitle>
            <DataFieldBody>
              {data.getPlant.pruneSeason || 'No data yet'}
            </DataFieldBody>
            <DataFieldTitle>{'Tips:'}</DataFieldTitle>
            <DataFieldBody>{data.getPlant.tips || 'No data yet'}</DataFieldBody>
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
