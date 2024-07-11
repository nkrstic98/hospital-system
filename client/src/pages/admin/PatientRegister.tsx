import Grid from "@mui/material/Grid";
import TextField from "@mui/material/TextField";
import Box from "@mui/material/Box";
import {RegisterPatientFormFields} from "./PatientAdmissionStepper.tsx";

export interface PatientRegisterProps {
    form: RegisterPatientFormFields;
    updateFormField: (fieldName: string, value: string) => void;
    userRegisterAttempted: boolean;
}

const PatientRegister = ({ form, updateFormField, userRegisterAttempted }: PatientRegisterProps) => {
    return (
        <>
            <Box component="form" noValidate sx={{ mt: 4 }}>
                <Grid container spacing={2}>
                    <Grid item xs={12} sm={6}>
                        <TextField
                            id="firstname"
                            name="firstname"
                            type="text"
                            label="First Name"
                            autoComplete="given-name"
                            value={form.firstname}
                            onChange={(e) => updateFormField("firstname", e.target.value)}
                            required
                            fullWidth
                            autoFocus
                            error={userRegisterAttempted && form.firstname === ""}
                            helperText={userRegisterAttempted && form.firstname === "" ? "This field is required" : ""}
                        />
                    </Grid>
                    <Grid item xs={12} sm={6}>
                        <TextField
                            id="lastname"
                            name="lastname"
                            label="Last Name"
                            autoComplete="family-name"
                            value={form.lastname}
                            onChange={(e) => updateFormField("lastname", e.target.value)}
                            required
                            fullWidth
                            error={userRegisterAttempted && form.lastname === ""}
                            helperText={userRegisterAttempted && form.lastname === "" ? "This field is required" : ""}
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
                            error={userRegisterAttempted && form.nationalIdentificationNumber === ""}
                            helperText={userRegisterAttempted && form.nationalIdentificationNumber === "" ? "This field is required" : ""}
                        />
                    </Grid>
                    <Grid item xs={12} sm={6}>
                        <TextField
                            id="mid"
                            name="mid"
                            type="text"
                            label="Medical Record Number"
                            value={form.medicalRecordNumber}
                            onChange={(e) => updateFormField("medicalRecordNumber", e.target.value)}
                            required
                            fullWidth
                            error={userRegisterAttempted && form.medicalRecordNumber === ""}
                            helperText={userRegisterAttempted && form.medicalRecordNumber === "" ? "This field is required" : ""}
                        />
                    </Grid>
                    <Grid item xs={12} sm={6}>
                        <TextField
                            id="phoneNumber"
                            name="phoneNumber"
                            type="text"
                            label="Phone Number"
                            value={form.phoneNumber}
                            onChange={(e) => updateFormField("phoneNumber", e.target.value)}
                            required
                            fullWidth
                            error={userRegisterAttempted && form.phoneNumber === ""}
                            helperText={userRegisterAttempted && form.phoneNumber === "" ? "This field is required" : ""}
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
                            fullWidth
                        />
                    </Grid>
                </Grid>
            </Box>
        </>
    )
}

export default PatientRegister;