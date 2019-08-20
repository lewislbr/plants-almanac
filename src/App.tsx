import React from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';

import { Home, Search } from './screens';

const App = () => {
  return (
    <>
      <BrowserRouter>
        <Switch>
          <Route exact path="/" component={Home} />
          <Route path="/home" component={Home} />
          <Route path="/search" component={Search} />
        </Switch>
      </BrowserRouter>
    </>
  );
};

export default App;
