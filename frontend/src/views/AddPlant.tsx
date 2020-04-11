import React, {useRef} from 'react';
import {gql, useMutation} from '@apollo/client';

import {Button, PageTitle} from '../components';

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
    addPlant(
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

export function AddPlant(props: {history: any}): JSX.Element {
  const [addPlant] = useMutation(ADD_PLANT);

  const nameElement = useRef<HTMLInputElement>(null);
  const otherNamesElement = useRef<HTMLInputElement>(null);
  const descriptionElement = useRef<HTMLTextAreaElement>(null);
  const plantSeasonElement = useRef<HTMLInputElement>(null);
  const harvestSeasonElement = useRef<HTMLInputElement>(null);
  const pruneSeasonElement = useRef<HTMLInputElement>(null);
  const tipsElement = useRef<HTMLTextAreaElement>(null);

  async function confirmAddPlant(
    event: React.FormEvent<HTMLFormElement>,
  ): Promise<void> {
    event.preventDefault();
    const name = nameElement.current?.value;
    const otherNames = otherNamesElement.current?.value || null;
    const description = descriptionElement.current?.value || null;
    const plantSeason = plantSeasonElement.current?.value || null;
    const harvestSeason = harvestSeasonElement.current?.value || null;
    const pruneSeason = pruneSeasonElement.current?.value || null;
    const tips = tipsElement.current?.value || null;
    if (!name) return event.preventDefault();
    await addPlant({
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
    props.history.push('/');
  }

  return (
    <>
      <section>
        <div>
          <PageTitle>{'Add Plant'}</PageTitle>
        </div>
      </section>
      <section>
        <form onSubmit={confirmAddPlant}>
          <div>
            <label className="label">
              {'Name'}{' '}
              <span className="text-gray-500 text-xs">{'(Required)'}</span>
            </label>
            <input className="input" type="text" ref={nameElement} required />
          </div>
          <div>
            <label className="label">{'Other Names'}</label>
            <input className="input" type="text" ref={otherNamesElement} />
          </div>
          <div>
            <label className="label">{'Description'}</label>
            <textarea className="input" rows={4} ref={descriptionElement} />
          </div>
          <div>
            <label className="label">{'Plant Season'}</label>
            <input className="input" type="text" ref={plantSeasonElement} />
          </div>
          <div>
            <label className="label">{'Harvest Season'}</label>
            <input className="input" type="text" ref={harvestSeasonElement} />
          </div>
          <div>
            <label className="label">{'Prune Season'}</label>
            <input className="input" type="text" ref={pruneSeasonElement} />
          </div>
          <div>
            <label className="label">{'Tips'}</label>
            <textarea className="input" rows={4} ref={tipsElement} />
          </div>
          <div className="flex justify-center">
            <Button style="primary" type="submit">
              {'Add plant'}
            </Button>
          </div>
        </form>
      </section>
    </>
  );
}
