import { React, useEffect, useState } from "react";
import { Spinner, Flex } from "@chakra-ui/react";
const LoadingScreen = (props) => {
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    setTimeout(() => setLoading(false), 1500);
  }, []);
  return (
    <>
      {loading ? (
        <Flex
          alignItems="center"
          justifyContent="center"
          minH="100vh"
          textAlign="center"
        >
          <Spinner color="body" bgColor="bg" verticalAlign="middle" size="xl" />
        </Flex>
      ) : (
        props.children
      )}
    </>
  );
};

export default LoadingScreen;
