import * as React from "react";
import {BrowserRouter, Route, Switch} from "react-router-dom";
import {Header} from "./components";
import {AddPlant, PlantDetails, Plants} from "./views";

export function App(): JSX.Element {
  return (
    <BrowserRouter>
      <Header />
      <main className="max-w-5xl mx-auto p-5pc">
        <Switch>
          <Route exact path="/" component={Plants} />
          <Route path="/addplant" component={AddPlant} />
          <Route path="/:_id" component={PlantDetails} />
        </Switch>
      </main>
    </BrowserRouter>
  );
}
