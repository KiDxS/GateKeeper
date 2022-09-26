import React, { useState } from "react";
import {
    Alert,
    Button,
    Center,
    Container,
    Paper,
    PasswordInput,
    Title,
} from "@mantine/core";
import { useForm } from "@mantine/form";
import axios from "axios";
import { IconAlertCircle } from "@tabler/icons";

function ChangePassword() {
    const [successStatus, setSuccessStatus] = useState();
    const [failedStatus, setFailedStatus] = useState();
    const form = useForm({
        initialValues: {
            current_password: "",
            new_password: "",
            confirm_new_password: "",
        },
        validate: {
            current_password: (value) =>
                value.length < 1 ? "Please enter your password." : null,
            new_password: (value) =>
                value.length < 12 ? "Must be more than 12 characters." : null,
            confirm_new_password: (value, values) =>
                value !== values.new_password
                    ? "The passwords you entered do not match."
                    : null,
        },
    });
    // Renders the flash messages
    const FlashMessage = () => {
        if (successStatus) {
            return (
                <Alert
                    icon={<IconAlertCircle size={16} />}
                    title="Done!"
                    color="teal"
                >
                    Your password has been changed.
                </Alert>
            );
        } if (failedStatus) {
            return (
                <Alert
                    icon={<IconAlertCircle size={16} />}
                    title="Bummer!"
                    color="red"
                >
                    Something has went wrong while changing your password.
                </Alert>
            );
        }
    };
    // Sends a POST request to the API.
    const changePasswordRequest = form.onSubmit((values) => {
        axios
            .post(
                "http://localhost:8080/api/v1/user/change-password",
                {
                    current_password: values.current_password,
                    new_password: values.new_password,
                    confirm_new_password: values.confirm_new_password,
                },
                {
                    withCredentials: true,
                }
            )
            .then((res) => {
                console.log(res.data);
                setSuccessStatus(true);
            })
            .catch((err) => {
                setFailedStatus(true);
            });
    });
    return (
        <div>
            <Container size={460} my={200}>
                <Paper radius="md" p="xl" withBorder>
                    <Title order={1} mb="md" align="center">
                        Change Password
                    </Title>
                    <FlashMessage />
                    <form onSubmit={changePasswordRequest}>
                        <PasswordInput
                            label="Current Password"
                            placeholder="Current Password"
                            withAsterisk
                            {...form.getInputProps("current_password")}
                        ></PasswordInput>
                        <PasswordInput
                            label="New Password"
                            placeholder="New Password"
                            mt={10}
                            withAsterisk
                            {...form.getInputProps("new_password")}
                        ></PasswordInput>
                        <PasswordInput
                            label="Confirm Password"
                            placeholder="Confirm Password"
                            mt={10}
                            withAsterisk
                            {...form.getInputProps("confirm_new_password")}
                        ></PasswordInput>
                        <Center>
                            <Button type="submit" size="md" mt="xl">
                                Submit
                            </Button>
                        </Center>
                    </form>
                </Paper>
            </Container>
        </div>
    );
}

export default ChangePassword;
