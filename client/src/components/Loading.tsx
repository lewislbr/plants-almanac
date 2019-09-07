import React from 'react';
import styled, { keyframes } from 'styled-components';

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
  position: relative;
  width: 64px;
  height: 64px;
`;

const Div2 = styled.div`
  box-sizing: border-box;
  display: block;
  position: absolute;
  width: 51px;
  height: 51px;
  margin: 6px;
  border: 6px solid var(--color-accent-primary);
  border-radius: 50%;
  animation: ${spin} 1.2s cubic-bezier(0.5, 0, 0.5, 1) infinite;
  border-color: var(--color-accent-primary) transparent transparent transparent;
`;

const Div3 = styled(Div2)`
  animation-delay: -0.45s;
`;

const Div4 = styled(Div2)`
  animation-delay: -0.3s;
`;

export const Loading: React.FunctionComponent = () => {
  return (
    <Div1>
      <Div2></Div2>
      <Div3></Div3>
      <Div4></Div4>
      <Div2></Div2>
    </Div1>
  );
};
