/* eslint-disable react/button-has-type */
import React from 'react';

export function Button(props: {
  children: string;
  onClick?: any;
  style: 'primary' | 'secondary' | 'danger';
  type: 'button' | 'submit' | 'reset';
}): JSX.Element {
  return (
    <button
      className={`${
        props.style === 'primary'
          ? 'bg-green-400 hover:bg-green-500'
          : props.style === 'secondary'
          ? 'bg-gray-300 hover:bg-gray-00'
          : props.style === 'danger'
          ? 'bg-red-400 hover:bg-red-500'
          : ''
      } font-semibold mt-8 px-16 py-3 rounded-lg text-white transition duration-500 ease-in-out`}
      type={props.type}
      onClick={props.onClick}
    >
      {props.children}
    </button>
  );
}
