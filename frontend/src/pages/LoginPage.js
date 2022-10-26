import { React, useState } from "react";
import {
    Box,
    Heading,
    Stack,
    FormControl,
    FormLabel,
    Input,
    Container,
    Button,
    Alert,
    AlertIcon,
    FormErrorMessage,
} from "@chakra-ui/react";
import { postData } from "../utils/postData";
import { useForm } from "react-hook-form";

const LoginPage = () => {
    const {
        register,
        handleSubmit,
        formState: { errors, isSubmitting },
    } = useForm();
    const [authIssueStatus, setAuthIssueStatus] = useState(false);
    const [errorStatus, setErrorStatus] = useState(false);
    const handleLogin = async(data) => {
        try {
            const url ="http://127.0.0.1:8080/api/v1/user/login"
            const response = await postData(url, data, null);
            console.log(data);
        } catch(err) {
            switch(err.response.status) {
                case 401:
                    setAuthIssueStatus(true);
                    break;
                default:
                    setErrorStatus(true);
            }
        }
    }

    const renderAlert = () => {
        if (authIssueStatus === true) {
            return (
                <Alert status="error">
                    <AlertIcon />
                    Username/Password is incorrect.
                </Alert>
            );
        }
        if (errorStatus === true) {
            return (
                <Alert status="error">
                    <AlertIcon />
                    There was an error processing your request
                </Alert>
            );
        }
    };
    return (
        <Container mt="12">
            <Stack spacing={{ base: 2, sm: 4 }}>
                <Heading size={{ base: "2xl", sm: "4xl" }} textAlign="center">
                    GateKeeper
                </Heading>
                <Heading size={{ base: "sm", sm: "md" }} textAlign="center">
                    Login to your account
                </Heading>
            </Stack>
            <Box my="7" borderWidth="1px" borderRadius="md" boxShadow="lg">
                <form onSubmit={handleSubmit(handleLogin)}>
                    <Stack spacing="4" p={{ base: 10, sm: 20 }}>
                        {renderAlert()}
                        <FormControl isInvalid={errors.username}>
                            <FormLabel htmlFor="username" fontWeight="bold">
                                Username
                            </FormLabel>

                            <Input
                                id="username"
                                type="text"
                                placeholder="Username"
                                {...register("username", {
                                    required: "This field is required.",
                                })}
                            />
                            <FormErrorMessage>
                                {errors.username && errors.username.message}
                            </FormErrorMessage>
                        </FormControl>
                        <FormControl isInvalid={errors.password}>
                            <FormLabel htmlFor="password" fontWeight="bold">
                                Password
                            </FormLabel>
                            <Input
                                type="password"
                                placeholder="Password"
                                {...register("password", {
                                    required: "This field is required.",
                                })}
                            />
                            <FormErrorMessage>
                                {errors.password && errors.password.message}
                            </FormErrorMessage>
                        </FormControl>

                        <Button
                            type="submit"
                            colorScheme="blue"
                            variant="solid"
                            size={{ base: "md", sm: "lg" }}
                            isLoading={isSubmitting}
                        >
                            Login
                        </Button>
                    </Stack>
                </form>
            </Box>
        </Container>
    );
};

export default LoginPage;
