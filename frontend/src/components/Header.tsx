import * as React from 'react';
import {NavLink} from 'react-router-dom';

const ActiveNavLink = {
  boxShadow: '0px 3px 0 0 lightgreen',
};

export function Header(): JSX.Element {
  return (
    <header className="bg-gray-100 fixed h-16 left-0 right-0 top-0 z-10">
      <div className="flex h-full items-center justify-between mx-auto max-w-5xl p-5pc text-xl">
        <NavLink className="text-3xl" exact to="/">
          {'🌿'}
        </NavLink>
        <nav>
          <NavLink
            exact={true}
            to="/"
            activeStyle={ActiveNavLink}
            className="tracking-tight"
          >
            {'Plants'}
          </NavLink>
          <NavLink
            to="/addplant"
            activeStyle={ActiveNavLink}
            className="ml-3 tracking-tight"
          >
            {'Add Plant'}
          </NavLink>
        </nav>
      </div>
    </header>
  );
}
