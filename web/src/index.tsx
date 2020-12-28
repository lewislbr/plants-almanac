import React from "react"
import ReactDOM from "react-dom"
import {
  ApolloClient,
  ApolloProvider,
  createHttpLink,
  InMemoryCache,
} from "@apollo/client"
import {setContext} from "@apollo/client/link/context"
import {CssBaseline, ThemeProvider} from "@material-ui/core"
import {App} from "./App"
import {theme} from "./theme"
import * as storageService from "./services/storage"
import {JWT} from "./constants/user"
import {AuthProvider} from "./contexts/auth"

const httpLink = createHttpLink({
  uri:
    process.env.NODE_ENV === "production"
      ? process.env.PLANTS_PRODUCTION_URL
      : process.env.PLANTS_DEVELOPMENT_URL,
})
const auth = setContext((_, {headers}) => {
  const jwt = storageService.retrieve(JWT)

  return {
    headers: {
      ...headers,
      ...(jwt && {authorization: `Bearer ${jwt}`}),
    },
  }
})

export const client = new ApolloClient({
  cache: new InMemoryCache(),
  link: auth.concat(httpLink),
})

ReactDOM.render(
  <AuthProvider>
    <ApolloProvider client={client}>
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <App />
      </ThemeProvider>
    </ApolloProvider>
  </AuthProvider>,
  document.getElementById("root"),
)

if (module.hot) {
  module.hot.accept()
}
