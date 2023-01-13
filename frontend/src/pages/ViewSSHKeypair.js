import {
    Button,
    Container,
    Heading,
    Icon,
    SimpleGrid,
    Stack,
    Flex,
} from "@chakra-ui/react";
import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import AuthProvider from "../components/AuthProvider";
import AutoResizeTextArea from "../components/AutoResizeTextArea";
import Navbar from "../components/Navbar";
import { IoClipboard } from "react-icons/io5";
import { fetchData } from "../utils/fetchData";

const ViewSSHKeypair = () => {
    const { id } = useParams();
    const [pubKey, setPubKey] = useState();
    const [privKey, setPrivKey] = useState();
    const [hasCopiedPrivKey, setHasCopiedPrivKey] = useState(false);
    const [hasCopiedPubKey, setHasCopiedPubKey] = useState(false);

    useEffect(() => {
        const fetchKeypair = async () => {
            try {
                const url = `http://127.0.0.1:8080/api/v1/key/${id}`;
                const options = { withCredentials: true };
                const response = await fetchData(url, options);
                const json = response.data;
                const data = json.data;
                setPubKey(data.pubKey);
                setPrivKey(data.privKey);
            } catch (err) {
                console.error(err)
            }
        };
        fetchKeypair();
    });

    const onCopyPrivKey = () => {
        // Writes the privKey to the clipboard
        navigator.clipboard.writeText(privKey);
        setHasCopiedPrivKey(true);
        // Sets the value to false after 1 second
        setTimeout(() => {
            setHasCopiedPrivKey(false);
        }, 1000);
    };
    const onCopyPubKey = () => {
        // Writes the pubKey to the clipboard
        navigator.clipboard.writeText(pubKey);
        setHasCopiedPubKey(true);

        // Sets the value to false after 1 second
        setTimeout(() => {
            setHasCopiedPubKey(false);
        }, 1000);
    };

    return (
        <AuthProvider>
            <Navbar />
            <Container maxW="container.xl">
                <SimpleGrid
                    minChildWidth="200px"
                    mt={8}
                    columns={2}
                    spacing={10}
                >
                    <Stack spacing={2}>
                        <Flex direction="row" justifyContent="space-between">
                            <Heading>Private Key</Heading>
                            <Button onClick={onCopyPrivKey} size="sm">
                                <Icon as={IoClipboard} />{" "}
                                {hasCopiedPrivKey
                                    ? "Copied"
                                    : "Copy to clipboard"}
                            </Button>
                        </Flex>
                        <AutoResizeTextArea value={privKey} />
                    </Stack>

                    <Stack spacing={2}>
                        <Flex direction="row" justifyContent="space-between">
                            <Heading>Public Key</Heading>
                            <Button onClick={onCopyPubKey} size="sm">
                                <Icon as={IoClipboard} />{" "}
                                {hasCopiedPubKey
                                    ? "Copied"
                                    : "Copy to clipboard"}
                            </Button>
                        </Flex>
                        <AutoResizeTextArea value={pubKey} />
                    </Stack>
                </SimpleGrid>
            </Container>
        </AuthProvider>
    );
};

export default ViewSSHKeypair;
