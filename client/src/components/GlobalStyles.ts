import { createGlobalStyle } from 'styled-components';

import IMBPlexSans300 from '../assets/fonts/IBMPlexSans300.woff2';
import IMBPlexSans400 from '../assets/fonts/IBMPlexSans400.woff2';
import IMBPlexSans700 from '../assets/fonts/IBMPlexSans700.woff2';

export const GlobalStyles = createGlobalStyle`
  :root {
    --color-accent-primary: hsl(0, 0%, 10%);
    --color-accent-primary-hover: hsl(0, 0%, 20%);
    --color-accent-secondary: hsl(0, 0%, 95%);
    --color-accent-secondary-hover: hsl(0, 0%, 85%);
    --color-dark: hsl(0, 0%, 10%);
    --color-dark-light: hsl(0, 0%, 30%);
    --color-dark-lighter: hsl(0, 0%, 50%);
    --color-dark-translucent: hsla(0, 0%, 10%, 0.85);
    --color-error: hsl(343, 100%, 45%);
    --color-light: hsl(0, 0%, 95%);
    --color-light-dark: hsl(0, 0%, 75%);
    --color-light-darker: hsl(0, 0%, 55%);
    --color-light-translucent: hsla(70, 0%, 95%, 0.85);
    --font-size-xxs: 15px;
    --font-size-xs: 17px;
    --font-size-s: 18px;
    --font-size-m: 22px;
    --font-size-l: 30px;
    --font-size-xl: 46px;
    --font-size-xxl: 80px;
    --padding-desktop: 5%;
    --padding-mobile: 3%;
    --page-width: 1200px;
    --spacing-xxs: 5px;
    --spacing-xs: 10px;
    --spacing-s: 20px;
    --spacing-m: 40px;
    --spacing-l: 80px;
    --spacing-xl: 160px;
    --radius-m: 6px;
    --radius-l: 10px;
  }

  @font-face {
    font-display: block;
	  font-family: IBM Plex Sans;
    font-style: normal;
	  font-weight: 300;
	  src: url(${IMBPlexSans300});
	}

  @font-face {
    font-display: block;
	  font-family: IBM Plex Sans;
    font-style: normal;
	  font-weight: 400;
	  src: url(${IMBPlexSans400});
	}

  @font-face {
    font-display: block;
	  font-family: IBM Plex Sans;
    font-style: normal;
	  font-weight: 700;
	  src: url(${IMBPlexSans700});
	}

  html
  body {
    background-color: var(--color-light);
    color: var(--color-dark);
    font-family: IBM Plex Sans, -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
    font-size: var(--font-size-s);

    @media (max-width: 555px) {
      font-size: var(--font-size-xs);
    }
  }

  /*! modern-normalize v0.5.0 | MIT License | https://github.com/sindresorhus/modern-normalize */

  /* Document
    ========================================================================== */

  /**
  * Use a better box model (opinionated).
  */

  html {
    box-sizing: border-box;
  }

  *,
  *::before,
  *::after {
    box-sizing: inherit;
  }

  /**
  * Use a more readable tab size (opinionated).
  */

  :root {
    -moz-tab-size: 4;
    tab-size: 4;
  }

  /**
  * 1. Correct the line height in all browsers.
  * 2. Prevent adjustments of font size after orientation changes in iOS.
  */

  html {
    line-height: 1.15; /* 1 */
    -webkit-text-size-adjust: 100%; /* 2 */
  }

  /* Sections
    ========================================================================== */

  /**
  * Remove the margin in all browsers.
  */

  body {
    margin: 0;
  }

  /* Grouping content
    ========================================================================== */

  /**
  * Add the correct height in Firefox.
  */

  hr {
    height: 0;
  }

  /* Text-level semantics
    ========================================================================== */

  /**
  * Add the correct text decoration in Chrome, Edge, and Safari.
  */

  abbr[title] {
    text-decoration: underline dotted;
  }

  /**
  * Add the correct font weight in Chrome, Edge, and Safari.
  */

  b,
  strong {
    font-weight: bolder;
  }

  /**
  * 1. Improve consistency of default fonts in all browsers. (https://github.com/sindresorhus/modern-normalize/issues/3)
  * 2. Correct the odd 'em' font sizing in all browsers.
  */

  body {
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
  }

  code,
  kbd,
  samp,
  pre {
    font-family: SFMono-Regular, Consolas, 'Liberation Mono', Menlo, Courier, monospace; /* 1 */
    font-size: 1em; /* 2 */
  }

  /**
  * Add the correct font size in all browsers.
  */

  small {
    font-size: 80%;
  }

  /**
  * Prevent 'sub' and 'sup' elements from affecting the line height in all browsers.
  */

  sub,
  sup {
    font-size: 75%;
    line-height: 0;
    position: relative;
    vertical-align: baseline;
  }

  sub {
    bottom: -0.25em;
  }

  sup {
    top: -0.5em;
  }

  /* Forms
    ========================================================================== */

  /**
  * 1. Change the font styles in all browsers.
  * 2. Remove the margin in Firefox and Safari.
  */

  button,
  input,
  optgroup,
  select,
  textarea {
    font-family: inherit; /* 1 */
    font-size: 100%; /* 1 */
    line-height: 1.15; /* 1 */
    margin: 0; /* 2 */
  }

  /**
  * Remove the inheritance of text transform in Edge and Firefox.
  * 1. Remove the inheritance of text transform in Firefox.
  */

  button,
  select { /* 1 */
    text-transform: none;
  }

  /**
  * Correct the inability to style clickable types in iOS and Safari.
  */

  button,
  [type='button'],
  [type='reset'],
  [type='submit'] {
    -webkit-appearance: button;
  }

  /**
  * Remove the inner border and padding in Firefox.
  */

  button::-moz-focus-inner,
  [type='button']::-moz-focus-inner,
  [type='reset']::-moz-focus-inner,
  [type='submit']::-moz-focus-inner {
    border-style: none;
    padding: 0;
  }

  /**
  * Restore the focus styles unset by the previous rule.
  */

  button:-moz-focusring,
  [type='button']:-moz-focusring,
  [type='reset']:-moz-focusring,
  [type='submit']:-moz-focusring {
    outline: 1px dotted ButtonText;
  }

  /**
  * Correct the padding in Firefox.
  */

  fieldset {
    padding: 0.35em 0.75em 0.625em;
  }

  /**
  * Remove the padding so developers are not caught out when they zero out 'fieldset' elements in all browsers.
  */

  legend {
    padding: 0;
  }

  /**
  * Add the correct vertical alignment in Chrome and Firefox.
  */

  progress {
    vertical-align: baseline;
  }

  /**
  * Correct the cursor style of increment and decrement buttons in Safari.
  */

  [type='number']::-webkit-inner-spin-button,
  [type='number']::-webkit-outer-spin-button {
    height: auto;
  }

  /**
  * 1. Correct the odd appearance in Chrome and Safari.
  * 2. Correct the outline style in Safari.
  */

  [type='search'] {
    -webkit-appearance: textfield; /* 1 */
    outline-offset: -2px; /* 2 */
  }

  /**
  * Remove the inner padding in Chrome and Safari on macOS.
  */

  [type='search']::-webkit-search-decoration {
    -webkit-appearance: none;
  }

  /**
  * 1. Correct the inability to style clickable types in iOS and Safari.
  * 2. Change font properties to 'inherit' in Safari.
  */

  ::-webkit-file-upload-button {
    -webkit-appearance: button; /* 1 */
    font: inherit; /* 2 */
  }

  /* Interactive
    ========================================================================== */

  /*
  * Add the correct display in Chrome and Safari.
  */

  summary {
    display: list-item;
  }
`;
