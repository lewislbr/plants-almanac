import React from 'react';

import { PlantCard } from '../components';

const Search = () => {
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
