import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import Login from './components/Login';
import AdminConsole from './components/AdminConsole';
import KidsConsole from './components/KidsConsole';

function App() {
  return (
    <Router>
      <Switch>
        <Route exact path="/" component={Login} />
        <Route path="/admin" component={AdminConsole} />
        <Route path="/kids" component={KidsConsole} />
      </Switch>
    </Router>
  );
}

export default App;