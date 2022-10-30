import { React, useEffect, useState } from "react";
import {
    BrowserRouter as Router,
    Routes,
    Route,
    Navigate,
} from "react-router-dom";
import Dashboard from "./pages/Dashboard";
import LoginPage from "./pages/LoginPage";
import { useCookies } from "react-cookie";

const App = () => {
    const [loggedIn, setLoggedIn] = useState(false);
    const [cookies] = useCookies();
    
    useEffect(() => {
        if (cookies.authToken && cookies.authToken.length !== 0) {
            setLoggedIn(true);
        }
        setLoggedIn(false);
    },[cookies])
    return (
        <Router>
            <Routes>
                <Route
                    exact
                    path="/"
                    element={
                        loggedIn ? <Navigate to="/dashboard" /> : <LoginPage />
                    }
                ></Route>
                <Route exact path="/dashboard" element={<Dashboard />}></Route>
                <Route exact path="/change-password" element=""></Route>
            </Routes>
        </Router>
    );
};

export default App;
