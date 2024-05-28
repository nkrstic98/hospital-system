import Box from '@mui/material/Box';
import CssBaseline from '@mui/material/CssBaseline';
import Button from "@mui/material/Button";
import AddIcon from "@mui/icons-material/Add";
import TableContainer from "@mui/material/TableContainer";
import Paper from "@mui/material/Paper";
import Table from "@mui/material/Table";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import TableCell from "@mui/material/TableCell";
import TableBody from "@mui/material/TableBody";
import {useNavigate} from "react-router-dom";
import {PatientService} from "../../services/patient/Patient.ts";
import {useEffect, useState} from "react";
import {Admission} from "../../types/Admission.ts";
import {Chip, Modal} from "@mui/material";

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

const PatientIntake = () => {
    const navigate = useNavigate();
    const patientService = new PatientService();
    const [admissions, setAdmissions] = useState<Admission[]>([]);
    const [discharged, setDischarged] = useState<boolean>(false);

    const getAdmissionTime = (admissionTime: Date) => {
        const date = new Date(admissionTime);
        return date.toLocaleDateString() + " " + date.toLocaleTimeString();
    }

    const handleDischarge = (id: string) => {
        patientService.DischargePatient(id).then((success) => {
            if(success) {
                const updatedAdmissions = admissions.filter((a) => a.id !== id);
                setAdmissions(updatedAdmissions);

                setDischarged(true);
            }
        });
    }

    useEffect(() => {
        patientService.GetAdmissions({
            statuses: ["pending", "admitted"]
        }).then((admissions) => {
            if(admissions) {
                setAdmissions(admissions);
            }
        });
    }, []);

    return (
        <Box sx={{ display: 'flex', mt: 4 }}>
            <CssBaseline />
            <Box component="main" sx={{ flexGrow: 1, p: 3 }}>
                <Button
                    variant="contained"
                    sx={{ mt: 3, mb: 5 }}
                    startIcon={<AddIcon />}
                    onClick={() => navigate("/admin/patients/admissions")}
                >
                    Admit New Patient
                </Button>
                <TableContainer component={Paper}>
                    <Table aria-label="collapsible table">
                        <TableHead>
                            <TableRow>
                                <TableCell />
                                <TableCell align="center">Patient</TableCell>
                                <TableCell align="center">Department</TableCell>
                                <TableCell align="center">Physician</TableCell>
                                <TableCell align="center">Admission Time</TableCell>
                                <TableCell align="center">Admission Status</TableCell>
                                <TableCell align="center"></TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {admissions.map((a) => (
                                <TableRow>
                                    <TableCell />
                                    <TableCell align="center">{a.patient}</TableCell>
                                    <TableCell align="center">{a.department}</TableCell>
                                    <TableCell align="center">{a.physician}</TableCell>
                                    <TableCell align="center">{getAdmissionTime(a.admissionTime)}</TableCell>
                                    <TableCell align="center">
                                        {
                                            a.status === "pending" ?
                                                <Chip label="Pending" color="warning" />
                                                :
                                                <Chip label="Admitted" color="success" />
                                        }
                                    </TableCell>
                                    {/*<TableCell align="center">*/}
                                    {/*    {a.status === "pending" &&*/}
                                    {/*        <Button variant="outlined" color="error" onClick={() => handleDischarge(a.id)}>*/}
                                    {/*        Discharge*/}
                                    {/*    </Button>}*/}
                                    {/*</TableCell>*/}
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </TableContainer>
            </Box>

            <Modal
                open={discharged}
                onClose={() => setDischarged(false)}
                aria-labelledby="parent-modal-title"
                aria-describedby="parent-modal-description"
            >
                <Box sx={{...style, width: 400}}>
                    <h2 id="parent-modal-title">Success!</h2>
                    <p id="parent-modal-description">
                        Patient has been discharged successfully.
                    </p>
                    <Button
                        variant="text"
                        onClick={() => setDischarged(false)}
                    >
                        Close
                    </Button>
                </Box>
            </Modal>
        </Box>
    );
}

export default PatientIntake;