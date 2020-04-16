import React from 'react';
import ReactDOM from 'react-dom';
import ApolloClient from 'apollo-boost';
import {ApolloProvider} from '@apollo/react-hooks';

import {App} from './App';
import './styles.css';

const client = new ApolloClient({
  uri:
    process.env.NODE_ENV === 'production'
      ? 'https://plants-almanac.herokuapp.com/'
      : 'http://localhost:4000/graphql',
});

ReactDOM.render(
  <ApolloProvider client={client}>
    <App />
  </ApolloProvider>,
  document.getElementById('root'),
);

if (module.hot) {
  module.hot.accept();
}
