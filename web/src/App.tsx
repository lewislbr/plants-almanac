import React, {StrictMode} from "react"
import {BrowserRouter, Route, Switch} from "react-router-dom"
import {Container} from "@material-ui/core"
import {Header} from "./components"
import {AddPlant, PlantDetails, PlantList} from "./views"

export function App(): JSX.Element {
  return (
    <StrictMode>
      <BrowserRouter>
        <Container maxWidth="md" style={{padding: "80px 5%"}}>
          <Header />
          <main>
            <Switch>
              <Route exact path="/" component={PlantList} />
              <Route exact path="/add-plant" component={AddPlant} />
              <Route exact path="/:id" component={PlantDetails} />
              <Route exact path="/edit/:id" component={AddPlant} />
            </Switch>
          </main>
        </Container>
      </BrowserRouter>
    </StrictMode>
  )
}
