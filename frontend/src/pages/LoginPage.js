import { React, useEffect, useRef } from "react";
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
import Navbar from "../components/Navbar";
import { postData } from "../utils/postData";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import { useCookies } from "react-cookie";
import LoadingScreen from "../components/LoadingScreen";

const LoginPage = () => {
    const [cookies] = useCookies();
    const navigate = useNavigate();
    const {
        register,
        handleSubmit,
        formState: { errors, isSubmitting },
    } = useForm();
    
    const loginStatus = {
        success: useRef(false),
        wrong_creds: useRef(false),
        error: useRef(false),
    };
    useEffect(() => {
        if (cookies.authToken && cookies.authToken.length !== 0) {
            navigate("/dashboard");
        }
    }, [navigate, cookies]);

    // handleLogin is a function that sends a login request to the login API
    const handleLogin = async (data) => {
        try {
            const url = "http://127.0.0.1:8080/api/v1/user/login";
            const options = { withCredentials: true };
            const response = await postData(url, data, options);
            loginStatus.success.current = true;
            navigate("/dashboard");
        } catch (err) {
            if (err.code === "ERR_NETWORK") {
                loginStatus.error.current = true;
                return;
            }
            switch (err.response.status) {
                case 401:
                    loginStatus.wrong_creds.current = true;
                    break;
                default:
                    loginStatus.error.current = true;
            }
        }
    };
    // renderAlert is a function that is used to render alerts if needed.
    const renderAlert = () => {
        if (loginStatus.success.current === true) {
            return (
                <Alert status="success">
                    <AlertIcon />
                    Successfully logged on.
                </Alert>
            );
        }
        if (loginStatus.wrong_creds.current === true) {
            return (
                <Alert status="error">
                    <AlertIcon />
                    Username/Password is incorrect.
                </Alert>
            );
        }
        if (loginStatus.error.current === true) {
            return (
                <Alert status="error">
                    <AlertIcon />
                    There was an error processing your request.
                </Alert>
            );
        }
    };
    return (
        <LoadingScreen>
            <Navbar />

            <Container>
                <Stack spacing={{ base: 2, sm: 4 }} mt="12">
                    <Heading
                        size={{ base: "2xl", sm: "4xl" }}
                        textAlign="center"
                    >
                        GateKeeper
                    </Heading>
                    <Heading size={{ base: "sm", sm: "md" }} textAlign="center">
                        Login to your account
                    </Heading>
                </Stack>
                <Box my="7" borderWidth="1px" borderRadius="md" boxShadow="2xl">
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
        </LoadingScreen>
    );
};

export default LoginPage;
