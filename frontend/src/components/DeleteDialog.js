import React from "react";
import {
    AlertDialog,
    AlertDialogBody,
    AlertDialogFooter,
    AlertDialogHeader,
    AlertDialogContent,
    AlertDialogOverlay,
    Button,
    useDisclosure,
} from "@chakra-ui/react";
import axios from "axios";
import { useMutation, useQueryClient } from "@tanstack/react-query";
const DeleteDialog = (props) => {
    const { isOpen, onOpen, onClose } = useDisclosure();
    const cancelRef = React.useRef();
    const queryClient = useQueryClient();
    const onDelete = async (id) => {
        try {
            const options = { withCredentials: true };
            const url = `http://127.0.0.1:8080/api/v1/key/${id}`;
            await axios.delete(url, options);
            console.log("hello");
            onClose();
        } catch (err) {}
    };
    const mutation = useMutation(onDelete, {
        onSuccess: () => queryClient.invalidateQueries("keypairs"),
    });
    return (
        <>
            <Button onClick={onOpen} variant="solid" colorScheme="red">
                Delete
            </Button>

            <AlertDialog
                motionPreset="slideInBottom"
                isOpen={isOpen}
                leastDestructiveRef={cancelRef}
                onClose={onClose}
            >
                <AlertDialogOverlay>
                    <AlertDialogContent>
                        <AlertDialogHeader fontSize="lg" fontWeight="bold">
                            Delete SSH keypair
                        </AlertDialogHeader>

                        <AlertDialogBody>
                            Are you sure? You can't undo this action afterwards.
                        </AlertDialogBody>

                        <AlertDialogFooter>
                            <Button ref={cancelRef} onClick={onClose}>
                                Cancel
                            </Button>
                            <Button
                                colorScheme="red"
                                onClick={() => mutation.mutate(props.labelID)}
                                ml={3}
                            >
                                Delete
                            </Button>
                        </AlertDialogFooter>
                    </AlertDialogContent>
                </AlertDialogOverlay>
            </AlertDialog>
        </>
    );
};

export default DeleteDialog;
