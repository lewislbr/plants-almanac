import styled from 'styled-components';

export const Header_ = styled.header`
  align-items: center;
  backdrop-filter: blur(5px);
  background: var(--color-light-translucent);
  display: flex;
  flex-direction: row;
  height: 60px;
  justify-content: space-between;
  left: 0%;
  padding: 0px var(--padding-desktop);
  position: fixed;
  right: 0;
  top: 0;
`;

export const Li = styled.li`
  margin-left: 30px;

  @media (max-width: 600px) {
    display: block;
    font-size: var(--font-size-l);
    margin: var(--spacing-s) 0;
    width: 100%;
  }
`;

export const Nav = styled.nav`
  align-items: center;
  display: flex;
  flex-direction: row;
`;

export const NavIcon = styled.label`
  align-items: center;
  cursor: pointer;
  display: flex;
  height: 75%;
  justify-content: center;
  padding: 10px;
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
  top: 10px;
`;

export const NavIconLine2 = styled(NavIconLine)`
  top: 18px;
  width: 70%;
`;

export const NavIconLine3 = styled(NavIconLine)`
  top: 18px;
  width: 70%;
`;

export const NavIconLine4 = styled(NavIconLine)`
  top: 26px;
`;

export const Input = styled.input`
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

export const NavModal = styled.nav`
  display: none;
  ${Input}:checked ~ & {
    backdrop-filter: blur(5px);
    background: var(--color-light-translucent);
    display: block;
    height: 100vh;
    left: 0;
    position: absolute;
    right: 0;
    top: 100%;
    width: 100vw;
  }
`;

export const Ul = styled.ul`
  align-items: center;
  display: flex;
  flex-direction: row;
  list-style: none;

  @media (max-width: 600px) {
    flex-direction: column;
  }
`;
