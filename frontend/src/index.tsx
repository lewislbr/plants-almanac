import React from 'react';
import ReactDOM from 'react-dom';
import {ApolloClient} from 'apollo-client';
import {InMemoryCache} from 'apollo-cache-inmemory';
import {HttpLink} from 'apollo-link-http';
import {ApolloProvider} from '@apollo/react-hooks';

import {App} from './App';
import './styles.css';

const cache = new InMemoryCache();
const link = new HttpLink({
  uri: 'http://localhost:4000/graphql',
});

const client = new ApolloClient({
  cache,
  link,
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
