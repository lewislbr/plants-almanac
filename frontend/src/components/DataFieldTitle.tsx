import React from 'react';

export function DataFieldTitle(props: {children: string}): JSX.Element {
  return <h3 className="block font-semibold mb-1 text-xl">{props.children}</h3>;
}
