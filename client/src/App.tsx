import React from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';

import { AddPlant, Home, PlantDetails, Search } from './screens';

import { GlobalStyles, HeaderComponent, Main } from './components';

export const App: React.FunctionComponent = () => {
  return (
    <BrowserRouter>
      <>
        <GlobalStyles />
        <HeaderComponent />
        <Main>
          <Switch>
            <Route exact path="/" component={Home} />
            <Route path="/home" component={Home} />
            <Route path="/addplant" component={AddPlant} />
            <Route exact path="/search" component={Search} />
            <Route path="/:name" component={PlantDetails} />
          </Switch>
        </Main>
      </>
    </BrowserRouter>
  );
};
