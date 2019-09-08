import React from 'react';
import { NavLink } from 'react-router-dom';

import {
  Header_,
  Li,
  Input,
  NavIcon,
  NavIconLine1,
  NavIconLine2,
  NavIconLine3,
  NavIconLine4,
  NavModal,
  Ul,
} from './HeaderStyles';

export const Header: React.FunctionComponent = () => {
  const closeNavModal = (): void => {
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
          <nav>
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
          </nav>
        </>
      ) : (
        <>
          <div>
            <h2>Plants Almanac</h2>
          </div>
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
                <NavLink to="/" onClick={closeNavModal}>
                  Home
                </NavLink>
              </Li>
              <Li>
                <NavLink to="/search" onClick={closeNavModal}>
                  Search Plants
                </NavLink>
              </Li>
              <Li>
                <NavLink to="/addplant" onClick={closeNavModal}>
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
