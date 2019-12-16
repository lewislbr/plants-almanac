import React from 'react';
import styled, { keyframes } from 'styled-components';

const Wrapper = styled.div`
  align-items: center;
  display: flex;
  justify-content: center;
  padding-top: 10%;
`;

const spin = keyframes`
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
`;

const Div1 = styled.div`
  display: inline-block;
  height: 64px;
  width: 64px;
`;

const Div2 = styled.div`
  animation: ${spin} 1.2s cubic-bezier(0.5, 0, 0.5, 1) infinite;
  border: 6px solid var(--color-accent-primary);
  border-color: var(--color-accent-primary) transparent transparent transparent;
  border-radius: 50%;
  box-sizing: border-box;
  display: block;
  height: 51px;
  margin: 6px;
  position: absolute;
  width: 51px;
`;

const Div3 = styled(Div2)`
  animation-delay: -0.45s;
`;

const Div4 = styled(Div2)`
  animation-delay: -0.3s;
`;

export function Loading(): JSX.Element {
  return (
    <Wrapper>
      <Div1>
        <Div2></Div2>
        <Div3></Div3>
        <Div4></Div4>
        <Div2></Div2>
      </Div1>
    </Wrapper>
  );
}
