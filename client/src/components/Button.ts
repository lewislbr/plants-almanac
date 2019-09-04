import styled from 'styled-components';

interface Props {
  primary?: boolean;
  secondary?: boolean;
}

export const Button = styled.button`
  background-color: ${(props: Props): string =>
    props.primary ? 'var(--color-accent-primary)' : 'var(--color-light)'};
  border-radius: var(--radius-m);
  color: ${(props: Props): string =>
    props.primary ? 'var(--color-light)' : 'var(--color-accent-primary)'};
  cursor: pointer;
  margin: var(--spacing-xs);
  padding: var(--spacing-xxs) var(--spacing-s);
`;
