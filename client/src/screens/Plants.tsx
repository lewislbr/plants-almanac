import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import 'regenerator-runtime/runtime';

import { PlantCard } from '../components';

const Plants: React.FunctionComponent = () => {
  const [plants, setPlants] = useState([]);
  const [search, setSearch] = useState('');
  const [query, setQuery] = useState('rosemary');

  useEffect(() => {
    const getData = async (): Promise<void> => {
      const TREFLE_TOKEN = process.env.TREFLE_TOKEN;

      const response = await fetch(
        `https://trefle.io/api/plants?q=${query}?token=${TREFLE_TOKEN}`
      );
      const data = await response.json();
      setPlants(data);
      console.log(data);
    };
    getData();
    console.log('useEffect here');
  }, [query]);

  const updateSearch = (event: React.FormEvent<HTMLInputElement>): void => {
    const target = event.target as HTMLInputElement;
    setSearch(target.value);
    console.log('updateSearch here');
  };

  const updateQuery = (event: React.FormEvent<HTMLFormElement>): void => {
    event.preventDefault();
    setQuery(search);
  };

  return (
    <>
      <div>
        <h1>Search</h1>
      </div>
      <form onSubmit={updateQuery}>
        <input type="text" value={search} onChange={updateSearch} />
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

export default Plants;
