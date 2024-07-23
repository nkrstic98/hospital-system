import Box from "@mui/material/Box";
import {useEffect, useState} from "react";
import {PatientService} from "../../services/patient/Patient.ts";
import {useAuth} from "../../router/AuthProvider.tsx";
import {Admission} from "../../types/Patient.ts";
import CssBaseline from "@mui/material/CssBaseline";
import TableContainer from "@mui/material/TableContainer";
import Paper from "@mui/material/Paper";
import Table from "@mui/material/Table";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import TableCell from "@mui/material/TableCell";
import TableBody from "@mui/material/TableBody";
import {Alert, Chip} from "@mui/material";
import {getTimeString} from "../../utils/utils.ts";
import Button from "@mui/material/Button";
import {useNavigate} from "react-router-dom";

const Admissions = () => {
    const { user } = useAuth();
    const navigate = useNavigate();
    const patientService = new PatientService();
    const [admissions, setAdmissions] = useState<Admission[]>([]);
    const [failure, setFailure] = useState<boolean>(false);
    const [success, setSuccess] = useState<boolean>(false);

    const getActiveAdmissions = () => {
        patientService.GetActiveAdmissions(user?.id ?? "").then((admissions) => {
            if(admissions !== undefined) {
                setAdmissions(admissions);
            }
        }).catch(() => {
            console.log("Failed to load admissions");
        });
    }

    const executeTransfer = (admissionId: string, accept: boolean) => {
        patientService.AcknowledgeTransfer(admissionId, accept).then((success) => {
            if(success) {
                setSuccess(true);
                getActiveAdmissions();
                setTimeout(() => {
                    setSuccess(false);
                }, 2000);
            } else {
                setFailure(true);
            }
        }).catch(() => {
            setFailure(true);
        });
    }

    const admitPatient = (admissionId: string) => {
        patientService.AdmitPatient(admissionId).then((success) => {
            if(success) {
                setSuccess(true);
                getActiveAdmissions();
                setTimeout(() => {
                    setSuccess(false);
                }, 2000);
            } else {
                setFailure(true);
            }
        }).catch(() => {
            setFailure(true);
        });
    }

    useEffect(() => {
        getActiveAdmissions();
    }, []);

    return (
        <Box sx={{ display: 'flex', mt: 7 }}>
            <CssBaseline />
            <Box component="main" sx={{ flexGrow: 1, p: 3 }}>
                {failure && <Alert severity="error">Operation failed!</Alert>}
                {success && <Alert severity="success">Success!</Alert>}
                <TableContainer component={Paper} sx={{mt: 4}}>
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
                                <TableRow key={a.id}>
                                    <TableCell />
                                    <TableCell align="center">{a.patient}</TableCell>
                                    <TableCell align="center">{a.department}</TableCell>
                                    <TableCell align="center">{a.physician}</TableCell>
                                    <TableCell align="center">{getTimeString(a.startTime)}</TableCell>
                                    <TableCell align="center">
                                        {
                                            a.status === "pending" ?
                                                <Chip label="Pending" color="warning" />
                                                :
                                                a.status === "awaiting-transfer" ?
                                                    <Chip label="Transfer Requested" color="info" />
                                                    :
                                                    <Chip label="Admitted" color="success" />
                                        }
                                    </TableCell>
                                    {a.status === "pending" && <TableCell align="center">
                                        <Button
                                            variant="contained"
                                            onClick={() => admitPatient(a.id)}
                                            size="small"
                                            color="secondary"
                                        >
                                            Admit
                                        </Button>
                                    </TableCell>}
                                    {a.status === "admitted" && <TableCell align="center">
                                        <Button
                                            variant="text"
                                            onClick={() => navigate(`/patients/admissions/${a.id}`)}
                                            size="small"
                                            color="primary"
                                        >
                                            DETAILS
                                        </Button>
                                    </TableCell>}
                                    {a.status == "awaiting-transfer" && <TableCell align="center">
                                        <Button
                                            variant="contained"
                                            onClick={() => executeTransfer(a.id, true)}
                                            size="small"
                                            color="success"
                                            sx={{ mr: 1 }}
                                        >
                                            ACCEPT
                                        </Button>
                                        <Button
                                            variant="contained"
                                            onClick={() => executeTransfer(a.id, false)}
                                            size="small"
                                            color="error"
                                        >
                                            DECLINE
                                        </Button>
                                    </TableCell>}
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </TableContainer>
            </Box>
        </Box>
    );
}

export default Admissions;