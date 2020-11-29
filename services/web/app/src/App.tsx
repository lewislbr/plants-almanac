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
              <Route path="/add-plant" component={AddPlant} />
              <Route path="/:id" component={PlantDetails} />
            </Switch>
          </main>
        </Container>
      </BrowserRouter>
    </StrictMode>
  )
}
