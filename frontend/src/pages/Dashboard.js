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
    const [keypairs, setKeypairs] = useState([]);
    useEffect(() => {
        async function fetchKeypairLabels() {
            try {
                const options = { withCredentials: true };
                const response = await fetchData(
                    "http://127.0.0.1:8080/api/v1/key",
                    options
                );
                const json = await response.data;
                setKeypairs(json.data);
            } catch (err) {}
        }
        fetchKeypairLabels();
    }, []);

    const labels = keypairs.map((keypair) => {
        return (
            <Tr>
                <Td>{keypair.label}</Td>
                <Td>
                    <Stack spacing={2} direction="row">
                        <Button variant="solid" colorScheme="blue">
                            View
                        </Button>
                        <DeleteDialog labelID={keypair.id} />
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
                            </Tbody>
                        </Table>
                    </TableContainer>
                </Container>
            </AuthProvider>
        </LoadingScreen>
    );
};

export default Dashboard;
