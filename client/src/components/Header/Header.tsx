import React from 'react';
import { NavLink } from 'react-router-dom';

import { Header, Li, Nav, Ul } from './HeaderStyles';

export const HeaderComponent: React.FunctionComponent = () => {
  return (
    <Header>
      <div>
        <h2>Plants Almanac</h2>
      </div>
      <Nav>
        <Ul>
          <Li>
            <NavLink to="/">Home</NavLink>
          </Li>
          <Li>
            <NavLink to="/search">Search Plants</NavLink>
          </Li>
          <Li>
            <NavLink to="/addplant">Add Plant</NavLink>
          </Li>
        </Ul>
      </Nav>
    </Header>
  );
};
