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
import {GetUserPermission} from "../../utils/utils.ts";
import {useAuth} from "../../router/AuthProvider.tsx";

const style = {
    position: 'absolute',
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

export interface PatientIntakeProps {
    section: string;
}

const PatientIntake = ({ section }: PatientIntakeProps) => {
    const { isAuthenticated, user } = useAuth();
    const navigate = useNavigate();
    const [permission, setPermission] = useState<string>();
    const patientService = new PatientService();
    const [admissions, setAdmissions] = useState<Admission[]>([]);
    const [discharged, setDischarged] = useState<boolean>(false);

    const getAdmissionTime = (admissionTime: Date) => {
        const date = new Date(admissionTime);
        return date.toLocaleDateString() + " " + date.toLocaleTimeString();
    }

    useEffect(() => {
        const permission = GetUserPermission(user, section);
        if (permission === undefined) {
            navigate("/login");
        }

        setPermission(permission);
    }, [isAuthenticated, section, user]);

    useEffect(() => {
        patientService.GetActiveAdmissions({
            statuses: ["pending", "admitted"]
        }).then((admissions) => {
            if(admissions) {
                setAdmissions(admissions);
            }
        });
    }, []);

    return (
        <Box sx={{ display: 'flex', mt: 7 }}>
            <CssBaseline />
            <Box component="main" sx={{ flexGrow: 1, p: 3 }}>
                {permission == "WRITE" && <Button
                    variant="contained"
                    sx={{ mb: 5 }}
                    startIcon={<AddIcon />}
                    onClick={() => navigate("/patient-intake/new-admission")}
                >
                    Admit Patient
                </Button>}
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