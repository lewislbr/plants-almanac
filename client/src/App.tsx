import React from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';

import { Home, PlantDetails, Search } from './screens';

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
            <Route exact path="/search" component={Search} />
            <Route path="/:name" component={PlantDetails} />
          </Switch>
        </main>
      </>
    </BrowserRouter>
  );
};
