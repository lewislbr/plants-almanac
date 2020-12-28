import React, {StrictMode, useContext} from "react"
import {BrowserRouter, Redirect, Route, Switch} from "react-router-dom"
import {Container} from "@material-ui/core"
import {
  AddPlant,
  CreateAccount,
  LogIn,
  PlantDetails,
  PlantList,
  Welcome,
} from "./views"
import {AuthContext} from "./contexts/auth"

export function App(): JSX.Element {
  const {authenticatedUser} = useContext(AuthContext)

  return (
    <StrictMode>
      {authenticatedUser ? (
        <BrowserRouter>
          <Container maxWidth="md" style={{padding: "30px 5% 80px 5%"}}>
            <Switch>
              <Route exact path="/plants" component={PlantList} />
              <Route exact path="/add-plant" component={AddPlant} />
              <Route exact path="/plants/:id" component={PlantDetails} />
              <Route exact path="/edit/:id" component={AddPlant} />
              <Redirect from="/" to="/plants" />
            </Switch>
          </Container>
        </BrowserRouter>
      ) : (
        <BrowserRouter>
          <Container maxWidth="md" style={{padding: "30px 5% 80px 5%"}}>
            <Switch>
              <Route exact path="/welcome" component={Welcome} />
              <Route exact path="/log-in" component={LogIn} />
              <Route exact path="/create-account" component={CreateAccount} />
              <Redirect from="/" to="/welcome" />
            </Switch>
          </Container>
        </BrowserRouter>
      )}
    </StrictMode>
  )
}
