import React from "react";
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
} from "@chakra-ui/react";
import Navbar from "../components/Navbar";
import padlockLogo from "../assets/padlock.png";
import { useForm } from "react-hook-form";
import AuthProvider from "../components/AuthProvider";

// Change password page
const ChangePasswordPage = () => {
    // Imports necessary properties from react-hook-form
    const {
        register,
        handleSubmit,
        getValues,
        formState: { errors, isSubmitting },
    } = useForm();

    // onSubmit function acts as a callback function that handles the behavioral aspect of the form if data are submitted.
    const onSubmit = async (data) => console.log(data);
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
