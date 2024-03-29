import { React, useState, useEffect } from "react";
import { Heading, Flex, Box, Link, Stack, Button } from "@chakra-ui/react";
import { useCookies } from "react-cookie";
import ThemeModeSwitch from "./ThemeModeSwitch";
import { fetchData } from "../utils/fetchData";
import { useNavigate } from "react-router-dom";
import { api } from "../utils/urls";

// Navbar component
const Navbar = () => {
    const [cookies] = useCookies();
    const [loggedIn, setLoggedIn] = useState(false);
    const navigate = useNavigate();

    useEffect(() => {
        if (cookies.authToken && cookies.authToken.length !== 0) {
            setLoggedIn(true);
            return;
        }
        setLoggedIn(false);
    }, [cookies]);

    // handleLogout is a function that sends a request to the logout API and redirects the user to the login page
    const handleLogout = async () => {
        try {
            const options = { withCredentials: true };
            await fetchData(
                api.logout,
                options
            );
            navigate("/");
        } catch (err) {
            console.log(err);
        }
    };
    return (
        <Flex
            h={16}
            px={4}
            alignItems="center"
            justifyContent="space-between"
            color="white"
            bgColor={"navbg"}
        >
            <Box color={"white"}>
                <Heading size={{ base: "md", sm: "xl" }}>gatekeeper</Heading>
            </Box>
            {loggedIn ? (
                <Stack direction="row" spacing={7} alignItems="center">
                    <Link href="/dashboard">Dashboard</Link>
                    <Link href="/change-password">Change Password</Link>
                    <Button variant="link" onClick={handleLogout} color="body">
                        Logout
                    </Button>

                    <ThemeModeSwitch />
                </Stack>
            ) : (
                <Stack direction="row" spacing={7} alignItems="center">
                    <ThemeModeSwitch />
                </Stack>
            )}
        </Flex>
    );
};

export default Navbar;
