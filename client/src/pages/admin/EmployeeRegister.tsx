import * as React from 'react';
import Avatar from '@mui/material/Avatar';
import Button from '@mui/material/Button';
import CssBaseline from '@mui/material/CssBaseline';
import TextField from '@mui/material/TextField';
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import PersonAddIcon from '@mui/icons-material/PersonAdd';
import {useCallback, useState} from "react";
import {Alert, FormControl, FormHelperText, InputLabel, Modal, OutlinedInput, Select} from "@mui/material";
import MenuItem from "@mui/material/MenuItem";
import {UserService} from "../../services/user/User.ts";
import {useNavigate} from "react-router-dom";

const style = {
    position: 'absolute' as 'absolute',
    top: '50%',
    left: '50%',
    transform: 'translate(-50%, -50%)',
    width: 400,
    bgcolor: 'background.paper',
    boxShadow: 24,
    pt: 2,
    px: 4,
    pb: 3,
};

type RegisterFormFields = {
    firstname: string;
    lastname: string;
    nationalIdentificationNumber: string;
    email: string;
    role: string;
    team: string | undefined;
}

const registerForm = () => {
    const [form, setForm] = useState<RegisterFormFields>({
        firstname: "",
        lastname: "",
        nationalIdentificationNumber: "",
        email: "",
        role: "",
        team: undefined,
    });

    const clearField = (fieldName: string) => {
        setForm({
            ...form,
            [fieldName]: ''
        });
    };

    const updateFormField = (fieldName: string, value: string) => {
        setForm({
            ...form,
            [fieldName]: value
        });
    };

    return { form, updateFormField, clearField };
}

const roles: Map<string, string> = new Map([
    ["ADMIN", "Administrator"],
    ["ATTENDING", "Attending Physician"],
    ["RESIDENT", "Resident Doctor"],
    ["NURSE", "Nurse"],
    ["TECHNICIAN", "Technician"],
]);

const teams: Map<string, string> = new Map([
    ["GENERAL", "Internal Medicine"],
    ["CARDIO", "Cardiology"],
    ["NEURO", "Neurology"],
    ["ORTHO", "Orthopedics"],
    ["OB-GYN", "Obstetrics and Gynecology"],
    ["PEDS", "Pediatrics"],
    ["ONCOLOGY", "Oncology"],
    ["PSYCH", "Psychiatry"],
    ["UROLOGY", "Urology"],
    ["RADIOLOGY", "Radiology"]
]);

const EmployeeRegister = ()=> {
    const navigate = useNavigate();
    const userService: UserService = new UserService();
    const [submitAttempted, setSubmitAttempted] = useState(false);
    const { form, updateFormField } = registerForm();
    const [role, setRole] = useState('');
    const [team, setTeam] = useState('');
    const [success, setSuccess] = useState(false);
    const [error, setError] = useState(false);

    const handleRoleChange = (role: string) => {
        setRole(role);
        updateFormField("role", role);
    }

    const handleTeamChange = (team: string) => {
        setTeam(team);
        updateFormField("team", team);
    }

    const isValidEmail = (email: string): boolean => {
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

        return emailRegex.test(email);
    }

    const handleSubmit = useCallback((event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();

        if (form.firstname === "" || form.lastname === "" || form.nationalIdentificationNumber === "" || form.email === "" || form.role === "" || form.team == "") {
            setSubmitAttempted(true);
            return;
        }

        if (!isValidEmail(form.email)) {
            setSubmitAttempted(true);
            return;
        }

        setSubmitAttempted(false);
        userService.Register({
            firstname: form.firstname,
            lastname: form.lastname,
            national_identification_number: form.nationalIdentificationNumber,
            email: form.email,
            role: form.role,
            team: form.team,
            joining_date: new Date(),
        }).then(r => {
            if (!r) {
                setError(true);
            }
            setSuccess(r);
        });
    }, [form]);

    return (
            <Container component="main" maxWidth="sm">
                <CssBaseline />
                <Box
                    sx={{
                        marginTop: 8,
                        display: 'flex',
                        flexDirection: 'column',
                        alignItems: 'center',
                    }}
                >
                    {error && <Alert variant="filled" severity="error" sx={{ m: 3 }}>Register failed. Check your data and try again!</Alert>}
                    <Avatar sx={{ m: 3, bgcolor: 'secondary.main' }}>
                        <PersonAddIcon />
                    </Avatar>
                    <Typography component="h1" variant="h5">
                        Add New Employee
                    </Typography>
                    <Box component="form" noValidate onSubmit={handleSubmit} sx={{ mt: 4 }}>
                        <Grid container spacing={2}>
                            <Grid item xs={12} sm={6}>
                                <TextField
                                    id="firstName"
                                    name="firstName"
                                    type="text"
                                    label="First Name"
                                    autoComplete="given-name"
                                    value={form.firstname}
                                    onChange={(e) => updateFormField("firstname", e.target.value)}
                                    required
                                    fullWidth
                                    autoFocus
                                    error={submitAttempted && form.firstname === ""}
                                    helperText={submitAttempted && form.firstname === "" ? "This field is required" : ""}
                                />
                            </Grid>
                            <Grid item xs={12} sm={6}>
                                <TextField
                                    id="lastName"
                                    name="lastName"
                                    label="Last Name"
                                    autoComplete="family-name"
                                    value={form.lastname}
                                    onChange={(e) => updateFormField("lastname", e.target.value)}
                                    required
                                    fullWidth
                                    error={submitAttempted && form.lastname === ""}
                                    helperText={submitAttempted && form.lastname === "" ? "This field is required" : ""}
                                />
                            </Grid>
                            <Grid item xs={12} sm={6}>
                                <TextField
                                    id="nid"
                                    name="nid"
                                    type="text"
                                    label="National ID"
                                    value={form.nationalIdentificationNumber}
                                    onChange={(e) => updateFormField("nationalIdentificationNumber", e.target.value)}
                                    required
                                    fullWidth
                                    error={submitAttempted && form.nationalIdentificationNumber === ""}
                                    helperText={submitAttempted && form.nationalIdentificationNumber === "" ? "This field is required" : ""}
                                />
                            </Grid>
                            <Grid item xs={12} sm={6}>
                                <TextField
                                    id="email"
                                    name="email"
                                    type="email"
                                    label="Email Address"
                                    value={form.email}
                                    onChange={(e) => updateFormField("email", e.target.value)}
                                    required
                                    fullWidth
                                    error={submitAttempted && (form.email === ""  || !isValidEmail(form.email))}
                                    helperText={submitAttempted && form.email === "" ? "This field is required" : !isValidEmail(form.email) ? "Invalid email address" : ""}
                                />
                            </Grid>
                            <Grid item xs={12} sm={6}>
                                <FormControl fullWidth required>
                                    <InputLabel id="role-label">Role</InputLabel>
                                    <Select
                                        labelId="role-label"
                                        id="role"
                                        value={role}
                                        onChange={(e) => handleRoleChange(e.target.value)}
                                        input={<OutlinedInput label="Name" />}
                                        required
                                        fullWidth
                                        error={submitAttempted && form.role === ""}
                                    >
                                        {Array.from(roles.entries()).map(([key, value]) => (
                                            <MenuItem key={value} value={key}>{value}</MenuItem>
                                        ))}
                                    </Select>
                                    {submitAttempted && form.role === "" && <FormHelperText error>This field is required</FormHelperText>}
                                </FormControl>
                            </Grid>
                            <Grid item xs={12} sm={6}>
                                <FormControl fullWidth>
                                    <InputLabel id="team-label">Team</InputLabel>
                                    <Select
                                        labelId="team-label"
                                        id="team"
                                        value={team}
                                        onChange={(e) => handleTeamChange(e.target.value)}
                                        input={<OutlinedInput label="Name" />}
                                        fullWidth
                                    >
                                        {Array.from(teams.entries()).map(([key, value]) => (
                                            <MenuItem key={value} value={key}>{value}</MenuItem>
                                        ))}
                                    </Select>
                                </FormControl>
                            </Grid>
                        </Grid>
                        <Button
                            type="submit"
                            fullWidth
                            variant="contained"
                            sx={{ mt: 3, mb: 2 }}
                        >
                            Sign Up
                        </Button>
                    </Box>
                </Box>

                <Modal
                    open={success}
                    onClose={() => navigate("/admin/employees")}
                    aria-labelledby="parent-modal-title"
                    aria-describedby="parent-modal-description"
                >
                    <Box sx={{ ...style, width: 400 }}>
                        <h2 id="parent-modal-title">Success!</h2>
                        <p id="parent-modal-description">
                            You have successfully registered a new employee.
                        </p>
                        <Button
                            variant="text"
                            onClick={() => navigate("/admin/employees")}
                        >
                            Go Back to Employees page
                        </Button>
                    </Box>
                </Modal>
            </Container>
    );
}

export default EmployeeRegister;