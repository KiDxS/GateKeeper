import { React, useEffect, useState } from "react";
import {
    Stack,
    Container,
    Table,
    Thead,
    Tbody,
    Tr,
    Th,
    Td,
    TableContainer,
    Button,
} from "@chakra-ui/react";

import LoadingScreen from "../components/LoadingScreen";
import Navbar from "../components/Navbar";
import AuthProvider from "../components/AuthProvider";
import DeleteDialog from "../components/DeleteDialog";
import { fetchData } from "../utils/fetchData";

// Dashboard page
const Dashboard = () => {
    const [keypairLabel, setKeypairLabel] = useState([]);
    useEffect(() => {
        async function fetchKeypairLabels() {
            try {
                const options = { withCredentials: true };
                const response = await fetchData(
                    "http://127.0.0.1:8080/api/v1/key",
                    options
                );
                const json = await response.data;
                setKeypairLabel(json.data);
            } catch (err) {}
        }
        fetchKeypairLabels();
    }, []);

    const labels = keypairLabel.map((label, id) => {
        return (
            <Tr>
                <Td>{label}</Td>
                <Td>
                    <Stack spacing={2} direction="row">
                        <Button variant="solid" colorScheme="blue">
                            View
                        </Button>
                        <DeleteDialog />
                    </Stack>
                </Td>
            </Tr>
        );
    });

    return (
        <LoadingScreen>
            <AuthProvider>
                <Navbar />
                <Container maxW="6xl" mt={8}>
                    <TableContainer>
                        <Table size="lg" variant="simple">
                            <Thead>
                                <Tr>
                                    <Th>Label</Th>
                                    <Th>Actions</Th>
                                </Tr>
                            </Thead>
                            <Tbody>
                                {labels}
                                {/* <Tr>
                                    <Td>This is just a test</Td>
                                    <Td>
                                        <Stack spacing={2} direction="row">
                                            <Button
                                                variant="solid"
                                                colorScheme="blue"
                                            >
                                                View
                                            </Button>
                                            <DeleteDialog />
                                        </Stack>
                                    </Td>
                                </Tr>
                                <Tr>
                                    <Td>This is just a test</Td>
                                    <Td>
                                        <Stack spacing={2} direction="row">
                                            <Button
                                                variant="solid"
                                                colorScheme="blue"
                                            >
                                                View
                                            </Button>
                                            <DeleteDialog />
                                        </Stack>
                                    </Td>
                                </Tr>  */}
                            </Tbody>
                        </Table>
                    </TableContainer>
                </Container>
            </AuthProvider>
        </LoadingScreen>
    );
};

export default Dashboard;
