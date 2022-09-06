import React from "react";
import {
    Button,
    Center,
    Container,
    Paper,
    PasswordInput,
    Title,
} from "@mantine/core";
import { useForm } from "@mantine/form";
function ChangePassword() {

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
                value !== values.new_password ? "The passwords you entered do not match." : null,
        },
    });

    const changePasswordRequest = form.onSubmit((values) => {

    })
    return (
        <div>
            <Container size={460} my={200}>
                <Paper radius="md" p="xl" withBorder>
                    <Title order={1} mb="md" align="center">
                        Change Password
                    </Title>
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
                            withAsterisk
                            {...form.getInputProps("new_password")}
                        ></PasswordInput>
                        <PasswordInput
                            label="New Current Password"
                            placeholder="New Current Password"
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
