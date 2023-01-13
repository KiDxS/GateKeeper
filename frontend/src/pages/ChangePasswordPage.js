import React, { useRef } from "react";
import {
    Container,
    Box,
    Stack,
    Heading,
    Image,
    FormControl,
    FormLabel,
    Input,
    Button,
    FormErrorMessage,
    Alert,
    AlertIcon,
} from "@chakra-ui/react";
import Navbar from "../components/Navbar";
import padlockLogo from "../assets/padlock.png";
import { useForm } from "react-hook-form";
import AuthProvider from "../components/AuthProvider";
import { postData } from "../utils/postData";
import { api } from "../utils/urls";

// Change password page
const ChangePasswordPage = () => {
    // Imports necessary properties from react-hook-form
    const {
        register,
        handleSubmit,
        getValues,
        formState: { errors, isSubmitting },
    } = useForm();
    const status = {
        success: useRef(),
        error: useRef(false),
    };

    const renderAlert = () => {
        if (status.success.current === true) {
            return (
                <Alert status="success">
                    <AlertIcon />
                    Your password has been changed.
                </Alert>
            );
        }
        if (status.success.current === false) {
            return (
                <Alert status="error">
                    <AlertIcon />
                    The password put as the current password is wrong.
                </Alert>
            );
        }
        if (status.error.current === true) {
            return (
                <Alert status="error">
                    <AlertIcon />
                    There was an error processing your request.
                </Alert>
            );
        }
    };

    // onSubmit function acts as a callback function that handles the behavioral aspect of the form if data are submitted.
    const onSubmit = async (data) => {
        try {
            const options = { withCredentials: true };
            await postData(api.change_password, data, options);
            status.success.current = true;
        } catch (err) {
            if (err.code === "ERR_NETWORK") {
                status.error.current = true;
                return;
            }
            switch (err.response.status) {
                case 400:
                    status.success.current = false;
                    break;
                default:
                    status.error.curent = true;
            }
        }
    };
    return (
        <AuthProvider>
            <Box>
                <Navbar />
                <Container>
                    <Stack mt={12} spacing={{ base: 2, sm: 10 }}>
                        <Image
                            src={padlockLogo}
                            alt="padlock logo"
                            boxSize={100}
                            mx="auto"
                        />
                        <Box
                            borderWidth={2}
                            borderRadius="md"
                            padding={7}
                            boxShadow="2xl"
                        >
                            <Heading size="lg" textAlign="center" mb={6}>
                                Change Password
                            </Heading>
                            <form onSubmit={handleSubmit(onSubmit)}>
                                <Stack spacing={5}>
                                    {renderAlert()}
                                    <FormControl
                                        isInvalid={errors.current_password}
                                    >
                                        <FormLabel htmlFor="current_password">
                                            Current Password
                                        </FormLabel>
                                        <Input
                                            type="password"
                                            // Registers the following input to react-hook-form
                                            {...register("current_password", {
                                                required:
                                                    "This field is required.",
                                            })}
                                        />
                                        <FormErrorMessage>
                                            {errors.current_password &&
                                                errors.current_password.message}
                                        </FormErrorMessage>
                                    </FormControl>
                                    <FormControl
                                        isInvalid={errors.new_password}
                                    >
                                        <FormLabel>New Password</FormLabel>
                                        <Input
                                            type="password"
                                            // Registers the following input to react-hook-form
                                            {...register("new_password", {
                                                required:
                                                    "This field is required.",
                                                minLength: {
                                                    value: 12,
                                                    message:
                                                        "This field must have a minimum length of 12 characters.",
                                                },
                                            })}
                                        />
                                        <FormErrorMessage>
                                            {errors.new_password &&
                                                errors.new_password.message}
                                        </FormErrorMessage>
                                    </FormControl>
                                    <FormControl
                                        isInvalid={errors.confirm_new_password}
                                    >
                                        <FormLabel>
                                            Confirm New Password
                                        </FormLabel>
                                        <Input
                                            type="password"
                                            // Registers the following input to react-hook-form
                                            {...register(
                                                "confirm_new_password",
                                                {
                                                    required:
                                                        "This field is required.",
                                                    minLength: {
                                                        value: 12,
                                                        message:
                                                            "This field must have a minimum length of 12 characters.",
                                                    },
                                                    validate: {
                                                        equals: (value) =>
                                                            getValues(
                                                                "new_password"
                                                            ) === value ||
                                                            "Passwords does not match",
                                                    },
                                                }
                                            )}
                                        />
                                        <FormErrorMessage>
                                            {errors.confirm_new_password &&
                                                errors.confirm_new_password
                                                    .message}
                                        </FormErrorMessage>
                                    </FormControl>
                                    <Button
                                        type="submit"
                                        colorScheme="blue"
                                        variant="solid"
                                        size={{ base: "md", sm: "lg" }}
                                        isLoading={isSubmitting}
                                    >
                                        Submit
                                    </Button>
                                </Stack>
                            </form>
                        </Box>
                    </Stack>
                </Container>
            </Box>
        </AuthProvider>
    );
};

export default ChangePasswordPage;
