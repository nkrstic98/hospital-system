import {Accordion, AccordionDetails, AccordionSummary, Alert, Grid} from "@mui/material";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import Typography from "@mui/material/Typography";
import TableContainer from "@mui/material/TableContainer";
import Paper from "@mui/material/Paper";
import Table from "@mui/material/Table";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import TableCell from "@mui/material/TableCell";
import TableBody from "@mui/material/TableBody";
import {Patient} from "../../types/Patient.ts";

export interface PatientInformationProps {
    patient: Patient;
}

const PatientInformation = ({ patient }: PatientInformationProps) => {
    return (
        <Grid item xs={12} sx={{mt: 2}}>
            <Accordion>
                <AccordionSummary
                    expandIcon={<ExpandMoreIcon/>}
                    aria-controls="panel1-content"
                    id="panel1-header"
                >
                    <b>Patient
                        Info: &nbsp; {patient.medicalRecordNumber} - {patient.firstname} {patient.lastname}</b> &nbsp; (expand
                    for more details)
                </AccordionSummary>
                <AccordionDetails>
                    <Grid container sx={{pl: 3}}>
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
                    <Typography variant="h6" gutterBottom component="div" sx={{m: 4}}>
                        Past visits
                    </Typography>
                    {patient.admissions!== null && patient.admissions.length > 0 ? (
                        <TableContainer component={Paper}>
                            <Table aria-label="collapsible table">
                                <TableHead>
                                    <TableRow>
                                        <TableCell/>
                                        <TableCell align="center">Start</TableCell>
                                        <TableCell align="center">Finish</TableCell>
                                        <TableCell align="center">Department</TableCell>
                                        <TableCell align="center">Attending Physician</TableCell>
                                    </TableRow>
                                </TableHead>
                                <TableBody>
                                    {/*{users.map((user) => (*/}
                                    {/*    <UserRow key={user.username} user={user} />*/}
                                    {/*))}*/}
                                </TableBody>
                            </Table>
                        </TableContainer>
                    ) : (
                        <Alert severity="info">There are no prior visits recorded for this patient.</Alert>
                    )}
                </AccordionDetails>
            </Accordion>
        </Grid>
    );
}

export default PatientInformation;