import React from "react"
import ReactDOM from "react-dom"
import {Router} from "react-router"
import {createBrowserHistory} from "history"
import {
  ApolloClient,
  ApolloProvider,
  HttpLink,
  InMemoryCache,
} from "@apollo/client"
import {App} from "./App"
import "./styles.css"

export const history = createBrowserHistory()
export const client = new ApolloClient({
  cache: new InMemoryCache(),
  link: new HttpLink({
    uri:
      process.env.NODE_ENV === "production"
        ? process.env.PLANTS_PRODUCTION_URL
        : process.env.PLANTS_DEVELOPMENT_URL,
  }),
})

ReactDOM.render(
  <Router history={history}>
    <ApolloProvider client={client}>
      <App />
    </ApolloProvider>
  </Router>,
  document.getElementById("root"),
)

if (module.hot) {
  module.hot.accept()
}
