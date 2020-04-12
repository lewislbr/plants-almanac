import React from 'react';

export function Alert(props: {
  deletePlant: (event: React.SyntheticEvent) => Promise<void>;
  setAlertOpen: React.Dispatch<React.SetStateAction<boolean>>;
}): JSX.Element {
  function closeAlert(): void {
    props.setAlertOpen(false);
  }

  return (
    <div className="items-center bg-backdrop fixed flex h-full justify-center left-0 right-0 top-0 w-full z-20">
      <div className="content-center bg-gray-100 fixed flex flex-col m-auto p-5 rounded-lg z-30">
        <p className="block mb-6">
          {'The plant will be removed from the database.'}
        </p>
        <button
          className="button button-danger mb-4"
          onClick={props.deletePlant}
          type="button"
        >
          {'Delete'}
        </button>
        <button
          className="button button-secondary"
          onClick={closeAlert}
          type="button"
        >
          {'Cancel'}
        </button>
      </div>
    </div>
  );
}
