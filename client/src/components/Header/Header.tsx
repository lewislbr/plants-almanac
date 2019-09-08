import React from 'react';
import { NavLink } from 'react-router-dom';

import {
  Header_,
  Li,
  Input,
  Nav,
  NavIcon,
  NavIconLine1,
  NavIconLine2,
  NavIconLine3,
  NavIconLine4,
  NavModal,
  Ul,
} from './HeaderStyles';

export const Header: React.FunctionComponent = () => {
  const navHandler = (): void => {
    const target = document.getElementById('icon') as HTMLInputElement;
    target.checked = false;
  };

  return (
    <Header_>
      {window.innerWidth >= 600 ? (
        <>
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
        </>
      ) : (
        <>
          <Input type="checkbox" id="icon" />
          <NavIcon htmlFor="icon">
            <NavIconLine1 />
            <NavIconLine2 />
            <NavIconLine3 />
            <NavIconLine4 />
          </NavIcon>
          <NavModal id="nav">
            <Ul>
              <Li>
                <NavLink to="/" onClick={navHandler}>
                  Home
                </NavLink>
              </Li>
              <Li>
                <NavLink to="/search" onClick={navHandler}>
                  Search Plants
                </NavLink>
              </Li>
              <Li>
                <NavLink to="/addplant" onClick={navHandler}>
                  Add Plant
                </NavLink>
              </Li>
            </Ul>
          </NavModal>
        </>
      )}
    </Header_>
  );
};
