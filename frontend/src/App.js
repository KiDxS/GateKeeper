import { React } from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";


import Dashboard from "./pages/Dashboard";
import LoginPage from "./pages/LoginPage";
import ChangePasswordPage from "./pages/ChangePasswordPage";

// App component
const App = () => {
    return (
        <Router>
            <Routes>
                <Route exact path="/" element={<LoginPage />} />
                <Route exact path="/dashboard" element={<Dashboard />} />
                <Route
                    exact
                    path="/change-password"
                    element={<ChangePasswordPage />}
                />
            </Routes>
        </Router>
    );
};

export default App;
