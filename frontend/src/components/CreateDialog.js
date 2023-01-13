import React from "react";
import {
    useDisclosure,
    Button,
    AlertDialog,
    AlertDialogOverlay,
    AlertDialogContent,
    AlertDialogHeader,
    AlertDialogBody,
    AlertDialogFooter,
    FormControl,
    FormLabel,
    Input,
    Stack,
    FormErrorMessage,
} from "@chakra-ui/react";
import { useForm } from "react-hook-form";
import { postData } from "../utils/postData";
import { api } from "../utils/urls";

const CreateDialog = () => {
    const { isOpen, onOpen, onClose } = useDisclosure();
    const cancelRef = React.useRef();
    const {
        register,
        handleSubmit,
        formState: { errors, isSubmitting, isValid },
    } = useForm({
        mode: "onChange",
    });

    const onSubmit = async (data) => {
        try {
            const options = { withCredentials: true };
            await postData(api.create_ssh_keypair, data, options);
        } catch (err) {
            console.error(err)
        }
    };
    const onClosing = () => {
        if (isValid) {
            onClose();
        }
        console.log(isValid);
        console.log(errors);
    };
    return (
        <>
            <Button onClick={onOpen} colorScheme="blue" size="lg">
                Create new
            </Button>
            <AlertDialog
                isOpen={isOpen}
                leastDestructiveRef={cancelRef}
                onClose={onClose}
            >
                <AlertDialogOverlay>
                    <AlertDialogContent>
                        <AlertDialogHeader fontSize="lg" fontWeight="bold">
                            Create a new SSH keypair
                        </AlertDialogHeader>

                        <form onSubmit={handleSubmit(onSubmit)}>
                            <AlertDialogBody>
                                <Stack spacing={4}>
                                    <FormControl isInvalid={errors.label}>
                                        <FormLabel htmlFor="label">
                                            Label
                                        </FormLabel>
                                        <Input
                                            id="label"
                                            type="text"
                                            placeholder="Label"
                                            {...register("label", {
                                                required:
                                                    "This field is required.",
                                            })}
                                        />
                                        <FormErrorMessage>
                                            {errors.label &&
                                                errors.label.message}
                                        </FormErrorMessage>
                                    </FormControl>
                                    <FormControl>
                                        <FormLabel htmlFor="password">
                                            Password
                                        </FormLabel>
                                        <Input
                                            id="password"
                                            type="password"
                                            placeholder="Password (leave it if blank)"
                                            {...register("password")}
                                        />
                                    </FormControl>
                                </Stack>
                            </AlertDialogBody>
                            <AlertDialogFooter>
                                <Button ref={cancelRef} onClick={onClose}>
                                    Cancel
                                </Button>
                                <Button
                                    type="submit"
                                    colorScheme="blue"
                                    ml={3}
                                    onClick={onClosing}
                                    isLoading={isSubmitting}
                                    is
                                >
                                    Create
                                </Button>
                            </AlertDialogFooter>
                        </form>
                    </AlertDialogContent>
                </AlertDialogOverlay>
            </AlertDialog>
        </>
    );
};

export default CreateDialog;
