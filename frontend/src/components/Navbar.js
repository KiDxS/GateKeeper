import React from "react";
import {
    Heading,
    useColorModeValue,

    Flex,
    Box,
} from "@chakra-ui/react";
import ThemeModeSwitch from "./ThemeModeSwitch";

const Navbar = () => {

    return (
        <Flex
            h={16}
            px={4}
            alignItems="center"
            justifyContent="space-between"
            bgColor={useColorModeValue("blackAlpha.900", "blackAlpha.600")}
        >
            <Box color={useColorModeValue("white")}>
                <Heading size={{ base: "md", sm: "xl" }}>GateKeeper</Heading>
            </Box>
            <ThemeModeSwitch />
        </Flex>
    );
};

export default Navbar;
