import {Grid} from "@mui/material";
import {Patient} from "../../types/Patient.ts";

export interface PatientInformationProps {
    patient: Patient;
}

const PatientInformation = ({ patient }: PatientInformationProps) => {
    return (
        <Grid item xs={12} sx={{mt: 2}}>
            <Grid container sx={{pl: 3}}>
                <Grid item xs={3}>
                    <p><b>Patient: </b></p>
                </Grid>
                <Grid item xs={8}>
                    <p>{patient.firstname} {patient.lastname}</p>
                </Grid>
                <Grid item xs={3}>
                    <p><b>National ID: </b></p>
                </Grid>
                <Grid item xs={8}>
                    <p>{patient.nationalIdentificationNumber}</p>
                </Grid>
                <Grid item xs={3}>
                    <p><b>Medical Record Number: </b></p>
                </Grid>
                <Grid item xs={8}>
                    <p>{patient.medicalRecordNumber}</p>
                </Grid>
                <Grid item xs={3}>
                    <p><b>Contact Phone: </b></p>
                </Grid>
                <Grid item xs={8}>
                    <p>{patient.phoneNumber}</p>
                </Grid>
                <Grid item xs={3}>
                    <p><b>Contact Email: </b></p>
                </Grid>
                <Grid item xs={8}>
                    <p>{patient.email}</p>
                </Grid>
            </Grid>
        </Grid>
    );
}

export default PatientInformation;