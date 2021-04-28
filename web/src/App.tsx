import React, {StrictMode, useContext, useEffect} from "react"
import {BrowserRouter, Redirect, Route, Switch} from "react-router-dom"
import {Container} from "@material-ui/core"
import {AddPlant, PlantDetails, PlantList} from "./plant/views"
import {Account, CreateAccount, LogIn, Welcome} from "./user/views"
import {AuthContext} from "./user/contexts/auth"
import {refreshToken} from "./user/services/user"

export function App(): JSX.Element {
  const {authenticatedUser} = useContext(AuthContext)

  useEffect(() => {
    function refresh(): void {
      if (authenticatedUser === true) {
        ;(async (): Promise<void> => {
          try {
            await refreshToken()
          } catch (err) {
            console.error(err)
          }
        })()
      }
    }

    window.addEventListener("beforeunload", refresh)

    return (): void => {
      window.removeEventListener("beforeunload", refresh)
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  return (
    <StrictMode>
      <Container maxWidth="md" style={{padding: "30px 5% 80px 5%"}}>
        {authenticatedUser ? (
          <BrowserRouter>
            <Switch>
              <Route exact path="/plants" component={PlantList} />
              <Route exact path="/add-plant" component={AddPlant} />
              <Route exact path="/plants/:id" component={PlantDetails} />
              <Route exact path="/edit/:id" component={AddPlant} />
              <Route exact path="/account" component={Account} />
              <Redirect from="/" to="/plants" />
            </Switch>
          </BrowserRouter>
        ) : (
          <BrowserRouter>
            <Switch>
              <Route exact path="/welcome" component={Welcome} />
              <Route exact path="/log-in" component={LogIn} />
              <Route exact path="/create-account" component={CreateAccount} />
              <Redirect from="/" to="/welcome" />
            </Switch>
          </BrowserRouter>
        )}
      </Container>
    </StrictMode>
  )
}
