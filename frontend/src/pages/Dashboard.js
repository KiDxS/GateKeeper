import { React } from "react";
import { Box, Stack, Heading, Text, Container } from "@chakra-ui/react";
import LoadingScreen from "../components/LoadingScreen";
import Navbar from "../components/Navbar";
import AuthProvider from "../components/AuthProvider";

const Dashboard = () => {
  return (
    <LoadingScreen>
      <AuthProvider>
        <Navbar />
        <Container>
          <Heading>Dashboard!!</Heading>
        </Container>
      </AuthProvider>
    </LoadingScreen>
  );
};

export default Dashboard;
