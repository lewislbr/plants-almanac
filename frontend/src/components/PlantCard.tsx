import React from 'react';

export function PlantCard(props: {name: string}): JSX.Element {
  return (
    <div>
      <h2>{props.name}</h2>
    </div>
  );
}
