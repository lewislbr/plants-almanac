import React from 'react';
import {BrowserRouter, Route, Switch} from 'react-router-dom';

import {Header} from './components';
import {AddPlant, Home, PlantDetails} from './views';

export function App(): JSX.Element {
  return (
    <BrowserRouter>
      <Header />
      <main>
        <Switch>
          <Route exact path="/" component={Home} />
          <Route path="/home" component={Home} />
          <Route path="/addplant" component={AddPlant} />
          <Route path="/:plantname" component={PlantDetails} />
        </Switch>
      </main>
    </BrowserRouter>
  );
}
