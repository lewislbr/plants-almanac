import React from "react"
import ReactDOM from "react-dom"
import {CssBaseline, ThemeProvider} from "@material-ui/core"
import {App} from "./App"
import {theme} from "./shared/theme"
import {AuthProvider} from "./user/contexts/auth"

ReactDOM.render(
  <AuthProvider>
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <App />
    </ThemeProvider>
  </AuthProvider>,
  document.getElementById("root"),
)

if (process.env.NODE_ENV === "development") {
  if (module.hot) {
    module.hot.accept()
  }
}
