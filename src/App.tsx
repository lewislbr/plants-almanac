import React from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';

import { Home, PlantDetails, Search } from './screens';

const App = () => {
  return (
    <>
      <BrowserRouter>
        <Switch>
          <Route exact path="/" component={Home} />
          <Route path="/home" component={Home} />
          <Route path="/search" component={Search} />
          <Route path="/plants/:name" component={PlantDetails} />
        </Switch>
      </BrowserRouter>
    </>
  );
};

export default App;
