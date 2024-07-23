import TableContainer from "@mui/material/TableContainer";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableRow from "@mui/material/TableRow";
import TableCell from "@mui/material/TableCell";
import Paper from "@mui/material/Paper";
import {Anamnesis} from "../../types/Admission.ts";

export interface AnamnesisPageProps {
    anamnesis: Anamnesis;
}

const AnamnesisPage = ({ anamnesis }: AnamnesisPageProps) => {
    return (
        <TableContainer component={Paper}>
            <Table>
                <TableBody>
                    <TableRow key="cc">
                        <TableCell><b>Chief Complaint</b></TableCell>
                        <TableCell>{anamnesis.chiefComplaint}</TableCell>
                    </TableRow>
                    <TableRow key="hpi">
                        <TableCell><b>History of Present Illness</b></TableCell>
                        <TableCell>{anamnesis.historyOfPresentIllness}</TableCell>
                    </TableRow>
                    <TableRow key="pmh">
                        <TableCell><b>Past Medical History</b></TableCell>
                        <TableCell>{anamnesis.pastMedicalHistory}</TableCell>
                    </TableRow>
                    <TableRow key="meds">
                        <TableCell><b>Medications</b></TableCell>
                        <TableCell>{anamnesis.medications}</TableCell>
                    </TableRow>
                    <TableRow key="al">
                        <TableCell><b>Allergies</b></TableCell>
                        <TableCell>{anamnesis.allergies}</TableCell>
                    </TableRow>
                    <TableRow key="fh">
                        <TableCell><b>Family History</b></TableCell>
                        <TableCell>{anamnesis.familyHistory}</TableCell>
                    </TableRow>
                    <TableRow key="sh">
                        <TableCell><b>Social History</b></TableCell>
                        <TableCell>{anamnesis.socialHistory}</TableCell>
                    </TableRow>
                    <TableRow key="pe">
                        <TableCell><b>Physical Examination</b></TableCell>
                        <TableCell>{anamnesis.physicalExamination}</TableCell>
                    </TableRow>
                </TableBody>
            </Table>
        </TableContainer>
    );
}

export default AnamnesisPage;
