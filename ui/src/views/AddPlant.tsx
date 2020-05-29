import * as React from 'react';
import {useAddPlantMutation} from '../graphql/types';

export function AddPlant({history}: {history: any}): JSX.Element {
  const [addPlant] = useAddPlantMutation();
  const nameElement = React.useRef<HTMLInputElement>(null);
  const otherNamesElement = React.useRef<HTMLInputElement>(null);
  const descriptionElement = React.useRef<HTMLTextAreaElement>(null);
  const plantSeasonElement = React.useRef<HTMLInputElement>(null);
  const harvestSeasonElement = React.useRef<HTMLInputElement>(null);
  const pruneSeasonElement = React.useRef<HTMLInputElement>(null);
  const tipsElement = React.useRef<HTMLTextAreaElement>(null);

  async function submitAddPlant(
    event: React.FormEvent<HTMLFormElement>,
  ): Promise<void> {
    event.preventDefault();

    const newPlant = {
      name: nameElement.current?.value || '',
      otherNames: otherNamesElement.current?.value || null,
      description: descriptionElement.current?.value || null,
      plantSeason: plantSeasonElement.current?.value || null,
      harvestSeason: harvestSeasonElement.current?.value || null,
      pruneSeason: pruneSeasonElement.current?.value || null,
      tips: tipsElement.current?.value || null,
    };

    await addPlant({variables: newPlant});

    history.push('/');
  }

  return (
    <>
      <section>
        <div>
          <h1 className="page-title">{'Add Plant'}</h1>
        </div>
      </section>
      <section>
        <form onSubmit={submitAddPlant}>
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
