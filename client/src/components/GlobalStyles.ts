import { createGlobalStyle } from 'styled-components';

export const GlobalStyles = createGlobalStyle`
  :root {
    --color-accent: hsl(300, 93%, 42%);
    --color-accent-hover: hsl(300, 93%, 35%);
    --color-dark: hsl(0, 0%, 10%);
    --color-dark-light: hsl(0, 0%, 30%);
    --color-dark-lighter: hsl(0, 0%, 50%);
    --color-dark-translucent: hsla(0, 0%, 10%, 0.85);
    --color-error: hsl(343, 100%, 45%);
    --color-light: hsl(70, 0%, 95%);
    --color-light-translucent: hsla(70, 0%, 95%, 0.85);
    --font-size-xs: 17px;
    --font-size-s: 18px;
    --font-size-m: 22px;
    --font-size-l: 30px;
    --font-size-xl: 46px;
    --font-size-xxl: 80px;
    --radius-m: 6px;
    --radius-l: 10px;
  }

  html
  body {
    background-color: var(--color-light);
    color: var(--color-dark);
    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
    font-size: var(--font-size-s);
    padding: 50px;

    @media (max-width: 555px) {
      font-size: var(--font-size-xs);
    }
  }
`;
