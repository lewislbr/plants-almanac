import React, { useEffect, useState } from 'react';
import 'regenerator-runtime/runtime';

import { PlantCard } from '../components';

const Search = () => {
  const API_KEY = process.env.API_KEY;

  const [plants, setPlants] = useState([]);
  const [search, setSearch] = useState('');
  const [query, setQuery] = useState('rosemary');

  const getData = async () => {
    const response = await fetch(
      `https://trefle.io/api/plants?q=${query}?token=${API_KEY}`
    );
    const data = response.json();
    setPlants(data);
    console.log(data);
  };

  useEffect(() => {
    getData();
    console.log('useEffect here');
  }, [query]);

  const updateSearch = (event: any) => {
    setSearch(event.target.value);
    console.log('updateSearch here');
  };

  const updateQuery = (event: any) => {
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
          <PlantCard />
        ))}
      </div>
    </>
  );
};

export default Search;
