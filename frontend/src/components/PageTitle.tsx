import React from 'react';

export function PageTitle(props: {children: string | undefined}): JSX.Element {
  return (
    <h1 className="font-black leading-none mb-6 tracking-tight text-5xl">
      {props.children}
    </h1>
  );
}
