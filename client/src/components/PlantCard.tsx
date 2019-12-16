import React from 'react';

interface Props {
  name: string;
}

export function PlantCard(props: Props): JSX.Element {
  return (
    <div>
      <h2>{props.name}</h2>
    </div>
  );
}
