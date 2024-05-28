import * as React from 'react';
import Button from '@mui/material/Button';
import CssBaseline from '@mui/material/CssBaseline';
import TextField from '@mui/material/TextField';
import Box from '@mui/material/Box';
import Card from '@mui/material/Card';
import CardMedia from '@mui/material/CardMedia';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import {useAuth} from "../../router/AuthProvider.tsx";
import {SessionService} from "../../services/session/Session.ts";
import {useCallback, useState} from "react";
import {Alert} from "@mui/material";

type LoginFormFields = {
    username: string;
    password: string;
};

const loginForm = () => {
    const [form, setForm] = useState<LoginFormFields>({
        username: "",
        password: "",
    });

    const clearField = (fieldName: string) => {
        setForm({
            ...form,
            [fieldName]: ''
        });
    };

    const updateFormField = useCallback((fieldName: string, value: string) => {
        setForm({
            ...form,
            [fieldName]: value
        });
    }, [form]);

    return { form, updateFormField, clearField };
}

const Login = () => {
    const { onLogin } = useAuth();
    const { form, updateFormField, clearField } = loginForm();
    const [loginAttempted, setLoginAttempted] = useState(false);
    const [submitAttempted, setSubmitAttempted] = useState(false);
    const sessionService = new SessionService();

    const handleSubmit = useCallback((event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();

        if (form.username === "" || form.password === "") {
            setSubmitAttempted(true);
            return;
        }

        setSubmitAttempted(false);
        sessionService.Login(form).then(r => {
            if (r == undefined) {
                setLoginAttempted(true);
                clearField("password");
                return;
            }

            onLogin(r);
        });
    }, [form, onLogin, clearField]);

    return (
        <Container component="main" maxWidth="xl">
            <CssBaseline />
            <Box
                sx={{
                    marginTop: 2,
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center',
                }}
                maxWidth="xl"
            >
                <Card>
                    <CardMedia
                        component="img"
                        height="250"
                        image="logo.png"
                        alt="Zmaj Medical Center"
                    />
                </Card>
            </Box>
            <Box
                sx={{
                    marginTop: 5,
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center',
                }}
                maxWidth="xs"
            >
                {loginAttempted && <Alert variant="filled" severity="error" sx={{ m: 3 }}>Login failed. Check your username and password!</Alert>}
                <Typography component="h1" variant="h5">
                    Log In
                </Typography>
                <Box component="form" onSubmit={handleSubmit} noValidate sx={{ mt: 3 }}>
                    <TextField
                        id="username"
                        name="username"
                        type="text"
                        label="Username"
                        value={form.username}
                        onChange={(e) => updateFormField("username", e.target.value)}
                        margin="normal"
                        required
                        fullWidth
                        autoFocus
                        error={submitAttempted && form.username === ""}
                        helperText={submitAttempted && form.username === "" ? "This field is required" : ""}
                    />
                    <TextField
                        id="password"
                        name="password"
                        type="password"
                        label="Password"
                        value={form.password}
                        onChange={(e) => updateFormField("password", e.target.value)}
                        margin="normal"
                        required
                        fullWidth
                        autoComplete="current-password"
                        error={submitAttempted && form.password === ""}
                        helperText={submitAttempted && form.username === "" ? "This field is required" : ""}
                    />
                    <Button
                        type="submit"
                        fullWidth
                        variant="contained"
                        sx={{ mt: 3, mb: 2 }}
                    >
                        LOG IN
                    </Button>
                </Box>
            </Box>
        </Container>
    );
};

export default Login;