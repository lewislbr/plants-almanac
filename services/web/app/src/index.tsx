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
import {CssBaseline, ThemeProvider} from "@material-ui/core"
import {App} from "./App"
import {theme} from "./theme"

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
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <App />
      </ThemeProvider>
    </ApolloProvider>
  </Router>,
  document.getElementById("root"),
)

if (module.hot) {
  module.hot.accept()
}
