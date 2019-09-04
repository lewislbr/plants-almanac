import React from 'react';
import { NavLink } from 'react-router-dom';

import { Header_, Li, Nav, Ul } from './HeaderStyles';

export const Header: React.FunctionComponent = () => {
  return (
    <Header_>
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
    </Header_>
  );
};
