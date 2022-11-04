import React from "react";
import {
  Heading,
  useColorModeValue,
  Flex,
  Box,
  Link,
  Stack,
  Button,
} from "@chakra-ui/react";
import ThemeModeSwitch from "./ThemeModeSwitch";
import { fetchData } from "../utils/fetchData";
import { useNavigate } from "react-router-dom";
import AuthProvider from "./AuthProvider";

const Navbar = () => {
  const navigate = useNavigate();
  const handleLogout = async () => {
    try {
      const options = { withCredentials: true };
      const response = await fetchData(
        "http://127.0.0.1:8080/api/v1/user/logout",
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
      color={useColorModeValue("white")}
      bgColor={useColorModeValue("blackAlpha.900", "blackAlpha.600")}
    >
      <Box color={useColorModeValue("white")}>
        <Heading size={{ base: "md", sm: "xl" }}>GateKeeper</Heading>
      </Box>
      <Stack direction="row" spacing={7} alignItems="center">
        <AuthProvider>
          <Link href="/dashboard">Dashboard</Link>
          <Link href="/change-password">Change Password</Link>
          <Button variant="link" onClick={handleLogout} color="body">
            Logout
          </Button>
        </AuthProvider>

        <ThemeModeSwitch />
      </Stack>
    </Flex>
  );
};

export default Navbar;
