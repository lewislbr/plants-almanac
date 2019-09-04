import React, { useState } from 'react';
import { Link } from 'react-router-dom';

import { H1, PlantCard } from '../components';

export const Search: React.FunctionComponent = () => {
  const [plants, setPlants] = useState([]);
  const [query, setQuery] = useState('');

  const getData = async (): Promise<void> => {
    const JWT_TOKEN = process.env.JWT_TOKEN;

    const response = await fetch(
      `https://trefle.io/api/plants?q=${query}?token=${JWT_TOKEN}`
    );
    const data = await response.json();
    setPlants(data);
    console.log(data);
  };

  const updateQuery = (event: React.FormEvent<HTMLInputElement>): void => {
    const target = event.target as HTMLInputElement;
    setQuery(target.value);
    console.log('updateQuery here');
  };

  const makeQuery = (event: React.FormEvent<HTMLFormElement>): void => {
    if (!query) return event.preventDefault();
    event.preventDefault();
    getData();
  };

  return (
    <>
      <div>
        <H1>Search</H1>
      </div>
      <form onSubmit={makeQuery}>
        <input type="text" value={query} onChange={updateQuery} />
        <button type="submit">Search</button>
      </form>
      <div>
        {plants.map((plant: any) => (
          <Link to={`/${name}`} key={plant.id}>
            <PlantCard key={plant.id} name={plant.name} />
          </Link>
        ))}
      </div>
    </>
  );
};
