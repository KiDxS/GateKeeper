// theme.js

// 1. import `extendTheme` function
import { extendTheme } from "@chakra-ui/react";

const theme = extendTheme({
    semanticTokens: {
        colors: {
            navbg: {
                default: "blackAlpha.900",
                _dark: "blackAlpha.600",
            },
            body: {
                default: "black",
                _dark: "white",
            },
        },
    },
});

export default theme;
