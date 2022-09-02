import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import Login from "./component/Login";
import Index from "./component/Index";

function App() {
    return (
        <BrowserRouter>
            <Routes>
                <Route exact path="/" element={<Index />}></Route>
                <Route exact path="/login" element={<Login />}></Route>
                <Route
                    path="*"
                    element={<Navigate to="/not-found" replace />}
                />
            </Routes>
        </BrowserRouter>
    );
}

export default App;
