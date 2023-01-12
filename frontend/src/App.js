import { React } from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

import Dashboard from "./pages/Dashboard";
import LoginPage from "./pages/LoginPage";
import ChangePasswordPage from "./pages/ChangePasswordPage";
import ViewSSHKeypair from "./pages/ViewSSHKeypair";

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
                <Route exact path="/view/:id" element={<ViewSSHKeypair />} />
            </Routes>
        </Router>
    );
};

export default App;
