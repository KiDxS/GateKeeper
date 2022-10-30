import React, { useEffect, useState } from "react";
import { useCookies } from "react-cookie";
import { useNavigate } from "react-router-dom";
const AuthProvider = (props) => {
    const [cookies] = useCookies();
    const [loggedIn, setLoggedIn] = useState(false);
    const navigate = useNavigate();
    useEffect(() => {
        if (cookies.authToken && cookies.authToken.length !== 0) {
            setLoggedIn(true);
            console.log(cookies.authToken.length);
            return;
        }

        navigate("/");
    }, [loggedIn, cookies, navigate]);
    return <>{loggedIn && props.children}</>;
};

export default AuthProvider;