import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import LoginPage from "./pages/LoginPage";

const App = () => {
    return (
        <Router>
            <Routes>
                <Route exact path="/login" element={<LoginPage />}></Route>
                <Route exact path="/dashboard" element=""></Route>
                <Route exact path="/change-password" element=""></Route>
            </Routes>
        </Router>
    );
};

export default App;
