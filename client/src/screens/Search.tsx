import React, { useState } from 'react';
import { Link } from 'react-router-dom';

import { Button, H1, Input, PlantCard, Section } from '../components';
import { Plant } from '../utils/plantInterface';

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
      <Section>
        <div>
          <H1>Search</H1>
        </div>
      </Section>
      <Section>
        <form onSubmit={makeQuery}>
          <Input type="text" value={query} onChange={updateQuery} />
          <Button type="submit" primary>
            Search
          </Button>
        </form>
      </Section>
      <Section>
        <div>
          {plants.map((plant: Plant) => (
            <Link to={`/${name}`} key={plant._id}>
              <PlantCard key={plant._id} name={plant.name} />
            </Link>
          ))}
        </div>
      </Section>
    </>
  );
};
