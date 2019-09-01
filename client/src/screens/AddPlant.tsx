import React, { useRef } from 'react';

export const AddPlant: React.FunctionComponent = () => {
  const nameElement = useRef<HTMLInputElement>(null!);
  const descriptionElement = useRef<HTMLTextAreaElement>(null!);
  const plantSeasonElement = useRef<HTMLInputElement>(null!);
  const harvestSeasonElement = useRef<HTMLInputElement>(null!);
  const pruneSeasonElement = useRef<HTMLInputElement>(null!);
  const tipsElement = useRef<HTMLTextAreaElement>(null!);

  const cancelAddPlant = (): void => {
    console.log('Canceled');
  };

  const confirmAddPlant = (event: React.FormEvent<HTMLFormElement>): void => {
    event.preventDefault();

    const name = nameElement.current.value;
    const description = descriptionElement.current.value;
    const plantSeason = plantSeasonElement.current.value;
    const harvestSeason = harvestSeasonElement.current.value;
    const pruneSeason = pruneSeasonElement.current.value;
    const tips = tipsElement.current.value;

    if (!name) return event.preventDefault();

    const query = {
      query: `
        mutation {
          createPlant(name: "${name}", description: "${description}", plantSeason: ["${plantSeason}"], harvestSeason: ["${harvestSeason}"], pruneSeason: ["${pruneSeason}"], tips: "${tips}") {
            _id
            name
            description
            plantSeason
            harvestSeason
            pruneSeason
            tips
          }
        }
      `,
    };

    (async (): Promise<void> => {
      try {
        const response = await fetch('http://localhost:4000/graphql', {
          method: 'POST',
          body: JSON.stringify(query),
          headers: {
            'Content-Type': 'application/json',
          },
        });
        if (response.status !== 200 && response.status !== 201) {
          throw new Error('Error');
        }
        const data = await response.json();
        console.log(data);
      } catch (error) {
        console.log(error);
      }
    })();
  };

  return (
    <>
      <div>
        <h2>Add Plant</h2>
      </div>
      <form onSubmit={confirmAddPlant}>
        <div>
          <label>Name</label>
          <input type="text" ref={nameElement} />
        </div>
        <div>
          <label>Description</label>
          <textarea rows={4} ref={descriptionElement} />
        </div>
        <div>
          <label>Plant Season</label>
          <input type="text" ref={plantSeasonElement} />
        </div>
        <div>
          <label>Harvest Season</label>
          <input type="text" ref={harvestSeasonElement} />
        </div>
        <div>
          <label>Prune Season</label>
          <input type="text" ref={pruneSeasonElement} />
        </div>
        <div>
          <label>Tips</label>
          <textarea rows={4} ref={tipsElement} />
        </div>
        <div>
          <button type="button" onClick={cancelAddPlant}>
            Cancel
          </button>
          <button type="submit">Confirm</button>
        </div>
      </form>
    </>
  );
};
