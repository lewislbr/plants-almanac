import React from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';

import { Home, PlantDetails, Plants } from './screens';

import { Navigation } from './components';

export const App: React.FunctionComponent = () => {
  return (
    <BrowserRouter>
      <>
        <Navigation />
        <main>
          <Switch>
            <Route exact path="/" component={Home} />
            <Route path="/home" component={Home} />
            <Route exact path="/plants" component={Plants} />
            <Route path="/plants/:name" component={PlantDetails} />
          </Switch>
        </main>
      </>
    </BrowserRouter>
  );
};
