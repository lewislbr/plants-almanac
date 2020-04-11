/* eslint-disable @typescript-eslint/no-non-null-assertion */
import React, {useRef} from 'react';
import {useMutation} from '@apollo/react-hooks';
import gql from 'graphql-tag';

const ADD_PLANT = gql`
  mutation AddPlant(
    $name: String!
    $otherNames: String
    $description: String
    $plantSeason: String
    $harvestSeason: String
    $pruneSeason: String
    $tips: String
  ) {
    createPlant(
      name: $name
      otherNames: $otherNames
      description: $description
      plantSeason: $plantSeason
      harvestSeason: $harvestSeason
      pruneSeason: $pruneSeason
      tips: $tips
    ) {
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

export function AddPlant(): JSX.Element {
  const [addPlant] = useMutation(ADD_PLANT);

  const nameElement = useRef<HTMLInputElement>(null!);
  const otherNamesElement = useRef<HTMLInputElement>(null!);
  const descriptionElement = useRef<HTMLTextAreaElement>(null!);
  const plantSeasonElement = useRef<HTMLInputElement>(null!);
  const harvestSeasonElement = useRef<HTMLInputElement>(null!);
  const pruneSeasonElement = useRef<HTMLInputElement>(null!);
  const tipsElement = useRef<HTMLTextAreaElement>(null!);

  function cancelAddPlant(): void {
    console.log('Canceled');
  }

  function confirmAddPlant(event: React.FormEvent<HTMLFormElement>): void {
    event.preventDefault();

    const name = nameElement.current.value;
    const otherNames =
      otherNamesElement.current.value == ''
        ? null
        : otherNamesElement.current.value;
    const description =
      descriptionElement.current.value == ''
        ? null
        : descriptionElement.current.value;
    const plantSeason =
      plantSeasonElement.current.value == ''
        ? null
        : plantSeasonElement.current.value;
    const harvestSeason =
      harvestSeasonElement.current.value == ''
        ? null
        : harvestSeasonElement.current.value;
    const pruneSeason =
      pruneSeasonElement.current.value == ''
        ? null
        : pruneSeasonElement.current.value;
    const tips =
      tipsElement.current.value == '' ? null : tipsElement.current.value;

    if (!name) return event.preventDefault();

    addPlant({
      variables: {
        name: name,
        otherNames: otherNames,
        description: description,
        plantSeason: plantSeason,
        harvestSeason: harvestSeason,
        pruneSeason: pruneSeason,
        tips: tips,
      },
    });
  }

  return (
    <>
      <section>
        <div>
          <h1>{'Add Plant'}</h1>
        </div>
      </section>
      <section>
        <form onSubmit={confirmAddPlant}>
          <div>
            <label>
              {'Name'} <span>{'(Required)'}</span>
            </label>
            <input type="text" ref={nameElement} />
          </div>
          <div>
            <label>{'Other Names'}</label>
            <input type="text" ref={otherNamesElement} />
          </div>
          <div>
            <label>{'Description'}</label>
            <textarea rows={4} ref={descriptionElement} />
          </div>
          <div>
            <label>{'Plant Season'}</label>
            <input type="text" ref={plantSeasonElement} />
          </div>
          <div>
            <label>{'Harvest Season'}</label>
            <input type="text" ref={harvestSeasonElement} />
          </div>
          <div>
            <label>{'Prune Season'}</label>
            <input type="text" ref={pruneSeasonElement} />
          </div>
          <div>
            <label>{'Tips'}</label>
            <textarea rows={4} ref={tipsElement} />
          </div>
          <div>
            <button type="button" onClick={cancelAddPlant}>
              {'Cancel'}
            </button>
            <button type="submit">{'Confirm'}</button>
          </div>
        </form>
      </section>
    </>
  );
}
