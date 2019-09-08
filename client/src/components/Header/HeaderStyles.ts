import styled from 'styled-components';
import { NavLink } from 'react-router-dom';

export const Header_ = styled.header`
  align-items: center;
  backdrop-filter: blur(5px);
  background: var(--color-light-translucent);
  display: flex;
  flex-direction: row;
  height: 60px;
  justify-content: space-between;
  left: 0;
  padding: 0px var(--padding-desktop);
  position: fixed;
  right: 0;
  top: 0;
  z-index: 2;
`;

export const StyledNavLink = styled(NavLink)`
  margin-left: var(--spacing-s);

  @media (max-width: 600px) {
    display: block;
    font-size: var(--font-size-l);
    margin: var(--spacing-m) 0;
    width: 100%;
  }
`;

export const NavIcon = styled.label`
  cursor: pointer;
  height: 26px;
  transform: rotate(0deg);
  transition: 0.25s ease-in-out;
  -webkit-tap-highlight-color: hsla(0, 0%, 0%, 0);
  -webkit-tap-highlight-color: transparent;
  width: 30px;
`;

export const NavIconLine = styled.span`
  background: var(--color-dark);
  border-radius: 50px;
  display: block;
  height: 3px;
  left: 0;
  position: absolute;
  transform: rotate(0deg);
  transition: 0.15s ease-in-out;
  width: 100%;
`;

export const NavIconLine1 = styled(NavIconLine)`
  top: 3px;
`;

export const NavIconLine2 = styled(NavIconLine)`
  top: 11px;
  width: 70%;
`;

export const NavIconLine3 = styled(NavIconLine)`
  top: 11px;
  width: 70%;
`;

export const NavIconLine4 = styled(NavIconLine)`
  top: 19px;
`;

export const NavIconInput = styled.input`
  display: none;
  :checked + ${NavIcon} {
    ${NavIconLine1} {
      left: 50%;
      top: 10px;
      width: 0%;
    }
    ${NavIconLine2} {
      transform: rotate(45deg);
      width: 100%;
    }
    ${NavIconLine3} {
      transform: rotate(-45deg);
      width: 100%;
    }
    ${NavIconLine4} {
      left: 50%;
      top: 26px;
      width: 0%;
    }
  }
`;

export const Nav = styled.nav`
  align-items: center;
  display: flex;
  flex-direction: row;
  list-style: none;

  @media (max-width: 600px) {
    display: none;
    ${NavIconInput}:checked ~ & {
      backdrop-filter: blur(5px);
      background: var(--color-light-translucent);
      display: block;
      flex-direction: column;
      height: 100vh;
      left: 0;
      padding: 0 var(--padding-desktop);
      position: absolute;
      top: 100%;
      width: 100vw;
    }
  }
`;
