import React from 'react';

interface Props {
  name: string;
}

export const PlantCard: React.FunctionComponent<Props> = (props: Props) => {
  return (
    <div>
      <h2>{props.name}</h2>
    </div>
  );
};
