import React from 'react';
import { NavLink } from 'react-router-dom';

export const Navigation: React.FunctionComponent = () => {
  return (
    <header>
      <div>
        <h1>Plants Almanac</h1>
      </div>
      <nav>
        <ul>
          <li>
            <NavLink to="/">Home</NavLink>
          </li>
          <li>
            <NavLink to="/search">Search Plants</NavLink>
          </li>
        </ul>
      </nav>
    </header>
  );
};
