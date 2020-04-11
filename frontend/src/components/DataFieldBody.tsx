import React from 'react';

export function DataFieldBody(props: {children: string}): JSX.Element {
  return <p className="mb-4">{props.children}</p>;
}
