import React from 'react';
import { NavLink } from 'react-router-dom';

const Navigation: React.FunctionComponent = () => {
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
            <NavLink to="/plants">Search Plants</NavLink>
          </li>
        </ul>
      </nav>
    </header>
  );
};

export default Navigation;
