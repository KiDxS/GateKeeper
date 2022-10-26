import React from "react";
import {
    Heading,
    useColorModeValue,
    Flex,
    Box,
    Link,
    Stack,
} from "@chakra-ui/react";
import ThemeModeSwitch from "./ThemeModeSwitch";

const Navbar = () => {
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
                <Link href="/dashboard">Dashboard</Link>
                <Link href="/change-password">Change Password</Link>
                <Link href="">Logout</Link>

                <ThemeModeSwitch />
            </Stack>
        </Flex>
    );
};

export default Navbar;
