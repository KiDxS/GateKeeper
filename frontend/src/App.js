import { React, useEffect, useState } from "react";
import {
  BrowserRouter as Router,
  Routes,
  Route,
  Navigate,
} from "react-router-dom";
import { useCookies } from "react-cookie";

import Dashboard from "./pages/Dashboard";
import LoginPage from "./pages/LoginPage";
import ChangePasswordPage from "./pages/ChangePasswordPage";

const App = () => {
  const [loggedIn, setLoggedIn] = useState(false);
  const [cookies] = useCookies();

  useEffect(() => {
    if (cookies.authToken && cookies.authToken.length !== 0) {
      setLoggedIn(true);
    }
    setLoggedIn(false);
  }, [cookies]);
  return (
    <Router>
      <Routes>
        <Route
          exact
          path="/"
          element={loggedIn ? <Navigate to="/dashboard" /> : <LoginPage />}
        />
        <Route exact path="/dashboard" element={<Dashboard />} />
        <Route exact path="/change-password" element={<ChangePasswordPage />} />
      </Routes>
    </Router>
  );
};

export default App;
