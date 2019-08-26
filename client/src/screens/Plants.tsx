import React, { useState } from 'react';
import { Link } from 'react-router-dom';

import { PlantCard } from '../components';

export const Plants: React.FunctionComponent = () => {
  const [plants, setPlants] = useState([]);
  const [input, setInput] = useState('');
  const [query, setQuery] = useState('');

  const getData = async (): Promise<void> => {
    const TREFLE_TOKEN = process.env.TREFLE_TOKEN;

    const response = await fetch(
      `https://trefle.io/api/plants?q=${query}?token=${TREFLE_TOKEN}`
    );
    const data = await response.json();
    setPlants(data);
    console.log(data);
  };

  const updateInput = (event: React.FormEvent<HTMLInputElement>): void => {
    const target = event.target as HTMLInputElement;
    setInput(target.value);
    console.log('updateInput here');
  };

  const makeQuery = (event: React.FormEvent<HTMLFormElement>): void => {
    if (!input) return event.preventDefault();
    event.preventDefault();
    setQuery(input);
    getData();
  };

  return (
    <>
      <div>
        <h1>Search</h1>
      </div>
      <form onSubmit={makeQuery}>
        <input type="text" value={input} onChange={updateInput} />
        <button type="submit">Search</button>
      </form>
      <div>
        {plants.map((plant) => (
          <Link to={`/plants/${name}`} key={plant}>
            <PlantCard />
          </Link>
        ))}
      </div>
    </>
  );
};
