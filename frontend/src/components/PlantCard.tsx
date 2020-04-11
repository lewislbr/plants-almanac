import React from 'react';

export function PlantCard(props: {name: string}): JSX.Element {
  return (
    <div>
      <h2 className="font-bold my-4 text-2xl">{props.name}</h2>
    </div>
  );
}
