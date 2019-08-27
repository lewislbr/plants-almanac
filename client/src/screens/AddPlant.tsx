import React from 'react';

const cancelAddPlant = (): void => {
  console.log('Canceled');
};

const confirmAddPlant = (): void => {
  console.log('Added plant');
};

export const AddPlant: React.FunctionComponent = () => {
  return (
    <>
      <div>
        <h2>Add Plant</h2>
      </div>
      <form>
        <div>
          <label>Name</label>
          <input type="text" />
        </div>
        <div>
          <label>Description</label>
          <textarea rows={4} />
        </div>
        <div>
          <label>Plant Season</label>
          <input type="text" />
        </div>
        <div>
          <label>Harvest Season</label>
          <input type="text" />
        </div>
        <div>
          <label>Prune Season</label>
          <input type="text" />
        </div>
        <div>
          <label>Tips</label>
          <textarea rows={4} />
        </div>
      </form>
      <div>
        <button onClick={cancelAddPlant}>Cancel</button>
        <button onClick={confirmAddPlant}>Confirm</button>
      </div>
    </>
  );
};
