import React from "react"
import RegisterPage from "./pages/register"
import LoginPage from "./pages/login";
import Password from "./pages/password";
import CreatePass from "./pages/create";
import UpdatePass from "./pages/update";

import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom";

function App() {
  return (
    <Router>
      <Switch>
        <Route path="/register">
          <RegisterPage />
        </Route>
        <Route path="/login">
          <LoginPage />
        </Route>
        <Route path="/site">
          <Password />
        </Route>
        <Route path="/createPass">
          <CreatePass />
        </Route>
        <Route path="/updatePass/:pass_id">
          <UpdatePass />
        </Route>
      </Switch>
    </Router>
  );
}

export default App;
