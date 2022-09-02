import React from "react";
import {
    Paper,
    createStyles,
    TextInput,
    PasswordInput,
    Button,
    Title,
    Text,
} from "@mantine/core";
import { useForm } from "@mantine/form";
import axios from "axios";

const useStyles = createStyles((theme) => ({
    wrapper: {
        minHeight: 900,
        backgroundSize: "cover",
        backgroundRepeat: "no-repeat",
        backgroundPosition: "0% 80%",

        backgroundImage:
            "linear-gradient(rgba(0, 0, 0, 0.5), rgba(0, 0, 0, 0.5)), url(https://images.unsplash.com/photo-1612206984652-9c468787ee11?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1920&q=80);",
    },

    form: {
        borderRight: `1px solid ${
            theme.colorScheme === "dark"
                ? theme.colors.dark[0]
                : theme.colors.dark[0]
        }`,
        minHeight: 900,
        maxWidth: 400,
        paddingTop: 80,

        [`@media (max-width: ${theme.breakpoints.sm}px)`]: {
            maxWidth: "100%",
        },
    },

    title: {
        color: theme.colorScheme === "dark" ? theme.white : theme.black,
        fontFamily: `Greycliff CF, ${theme.fontFamily}`,
    },

    logo: {
        color: theme.colorScheme === "dark" ? theme.white : theme.black,
        width: 120,
        display: "block",
        marginLeft: "auto",
        marginRight: "auto",
    },
}));

function Login() {
    const form = useForm({
        initialValues: {
            username: "",
            password: "",
        },
        validate: {
            username: (value) =>
                value.length < 1 ? "Username must not be empty" : null,
            password: (value) =>
                value.length < 1 ? "Password must not be empty" : null,
        },
    });
    const sendLoginRequest = form.onSubmit((values) => {
        axios
            .post(
                "http://localhost:8080/api/v1/user/login",
                {
                    username: values.username,
                    password: values.password,
                },
                {
                    withCredentials: true,
                }
            )
            .then((res) => {
                console.log(res.status);
            })
            .catch((err) => {});
    });
    const { classes } = useStyles();
    return (
        <div className={classes.wrapper}>
            <form onSubmit={sendLoginRequest}>
                <Paper className={classes.form} radius={0} p={30}>
                    <Title
                        order={1}
                        className={classes.title}
                        align="center"
                        mt="md"
                        mb={50}
                    >
                        GateKeeper Login Page
                    </Title>
                    <TextInput
                        label="Username"
                        placeholder="Your username"
                        size="md"
                        {...form.getInputProps("username")}
                    />
                    <PasswordInput
                        label="Password"
                        placeholder="Your password"
                        mt="md"
                        size="md"
                        {...form.getInputProps("password")}
                    />
                    <Button
                        fullWidth
                        mt="xl"
                        size="md"

                        type="submit"
                    >
                        {" "}
                        Login
                    </Button>
                </Paper>
            </form>
        </div>
    );
}

export default Login;
