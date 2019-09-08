import React from 'react';

import {
  Header_,
  Nav,
  NavIcon,
  NavIconInput,
  NavIconLine1,
  NavIconLine2,
  NavIconLine3,
  NavIconLine4,
  StyledNavLink,
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
          <Nav>
            <StyledNavLink to="/">Home</StyledNavLink>
            <StyledNavLink to="/search">Search Plants</StyledNavLink>
            <StyledNavLink to="/addplant">Add Plant</StyledNavLink>
          </Nav>
        </>
      ) : (
        <>
          <div>
            <h2>Plants Almanac</h2>
          </div>
          <NavIconInput type="checkbox" id="icon" />
          <NavIcon htmlFor="icon">
            <NavIconLine1 />
            <NavIconLine2 />
            <NavIconLine3 />
            <NavIconLine4 />
          </NavIcon>
          <Nav id="nav">
            <StyledNavLink to="/" onClick={closeNavModal}>
              Home
            </StyledNavLink>
            <StyledNavLink to="/search" onClick={closeNavModal}>
              Search Plants
            </StyledNavLink>
            <StyledNavLink to="/addplant" onClick={closeNavModal}>
              Add Plant
            </StyledNavLink>
          </Nav>
        </>
      )}
    </Header_>
  );
};
