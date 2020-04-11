import React from 'react';
import {NavLink} from 'react-router-dom';

export function Header(): JSX.Element {
  return (
    <header>
      <div>
        <h2>{'Plants Almanac'}</h2>
      </div>
      <nav>
        <NavLink to="/">{'Home'}</NavLink>
        <NavLink to="/search">{'Search Plants'}</NavLink>
        <NavLink to="/addplant">{'Add Plant'}</NavLink>
      </nav>
    </header>
  );
}
