import { React } from "react";
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
    Flex,
    Heading,
    Link,
} from "@chakra-ui/react";

import LoadingScreen from "../components/LoadingScreen";
import Navbar from "../components/Navbar";
import AuthProvider from "../components/AuthProvider";
import DeleteDialog from "../components/DeleteDialog";
import { fetchData } from "../utils/fetchData";
import { useQuery } from "@tanstack/react-query";
import CreateDialog from "../components/CreateDialog";
import { api, app } from "../utils/urls";

// Dashboard page
const Dashboard = () => {
    const fetchKeypairs = async () => {
        try {
            const options = { withCredentials: true };
            const response = await fetchData(
                api.retrieve_all_ssh_keypairs,
                options
            );
            const json = await response.data;
            // console.log(json);
            return json.data;
        } catch (err) {
            console.error(err)
        }
    };
    const { isLoading, data } = useQuery(["keypairs"], fetchKeypairs);

    if (isLoading) return "Loading";

    const labels = data.map((keypair) => {
        return (
            <Tr>
                <Td>{keypair.label}</Td>
                <Td>
                    <Stack spacing={2} direction="row">
                        <Link href={app.view_ssh_keypair + keypair.id}>
                            <Button variant="solid" colorScheme="blue">
                                View
                            </Button>
                        </Link>
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
                <Container maxW="6xl" mt={6}>
                    <Flex justifyContent="space-between" mb={4}>
                        <Heading as="h2">SSH Keys</Heading>
                        <CreateDialog />
                    </Flex>
                    <TableContainer>
                        <Table size="lg" variant="simple">
                            <Thead>
                                <Tr>
                                    <Th>Label</Th>
                                    <Th>Actions</Th>
                                </Tr>
                            </Thead>
                            <Tbody>{labels}</Tbody>
                        </Table>
                    </TableContainer>
                </Container>
            </AuthProvider>
        </LoadingScreen>
    );
};

export default Dashboard;
