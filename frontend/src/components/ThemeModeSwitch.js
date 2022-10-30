import React from "react";
import { useColorMode, IconButton, useColorModeValue } from "@chakra-ui/react";
import { MoonIcon, SunIcon } from "@chakra-ui/icons";
import { AnimatePresence, motion } from "framer-motion";

// ThemeModeSwitch is a function that is used to switch  between two color mode which are : white & dark.
const ThemeModeSwitch = () => {
  const { toggleColorMode } = useColorMode();
  return (
    <AnimatePresence exitBeforeEnter initial={false}>
      <motion.div
        style={{ display: "inline-block" }}
        key={useColorModeValue("light", "dark")}
        initial={{ y: -20, opacity: 0 }}
        animate={{ y: 0, opacity: 1 }}
        exit={{ y: 20, opacity: 0 }}
        transition={{ duration: 0.2 }}
      >
        <IconButton
          aria-label="Toggle theme"
          icon={useColorModeValue(<MoonIcon />, <SunIcon />)}
          color={useColorModeValue("black", "white")}
          onClick={toggleColorMode}
        ></IconButton>
      </motion.div>
    </AnimatePresence>
  );
};

export default ThemeModeSwitch;
