import React, { useEffect, useState } from 'react';
import 'regenerator-runtime/runtime';

import { PlantCard } from '../components';

const Search = () => {
  const API_KEY = process.env.API_KEY;

  const [query, setQuery] = useState('rosemary');

  const getData = async () => {
    const response = await fetch(
      `https://trefle.io/api/plants?q=${query}?token=${API_KEY}`
    );
    const data = response.json();
    console.log(data);
  };

  useEffect(() => {
    getData();
  }, []);

  return (
    <>
      <div>
        <h1>Search</h1>
      </div>
      <form>
        <input type="text" />
        <button type="submit">Search</button>
      </form>
      <div>
        <PlantCard />
      </div>
    </>
  );
};

export default Search;
