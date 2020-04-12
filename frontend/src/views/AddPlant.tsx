import React, {useRef} from 'react';

import {useAddPlantMutation} from '../graphql/mutations/addPlant.graphql';

export function AddPlant(props: {history: any}): JSX.Element {
  const [addPlantMutation] = useAddPlantMutation();

  const nameElement = useRef<HTMLInputElement>(null);
  const otherNamesElement = useRef<HTMLInputElement>(null);
  const descriptionElement = useRef<HTMLTextAreaElement>(null);
  const plantSeasonElement = useRef<HTMLInputElement>(null);
  const harvestSeasonElement = useRef<HTMLInputElement>(null);
  const pruneSeasonElement = useRef<HTMLInputElement>(null);
  const tipsElement = useRef<HTMLTextAreaElement>(null);

  async function addPlant(
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
    await addPlantMutation({
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
          <h1 className="page-title">{'Add Plant'}</h1>
        </div>
      </section>
      <section>
        <form onSubmit={addPlant}>
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
          <div className="flex justify-center pt-6">
            <button className="button button-primary" type="submit">
              {'Add plant'}
            </button>
          </div>
        </form>
      </section>
    </>
  );
}
