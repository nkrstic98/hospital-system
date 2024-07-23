import {AdmissionDetails, Log, Vitals} from "../../types/Admission.ts";
import React, {useState} from "react";
import Box from "@mui/material/Box";
import Paper from "@mui/material/Paper";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableRow from "@mui/material/TableRow";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TextField from "@mui/material/TextField";
import {InputAdornment} from "@mui/material";
import Button from "@mui/material/Button";

export interface VitalsProps {
    permission: string;
    admission: AdmissionDetails;
    updateAdmission: (admission: AdmissionDetails, log: Log) => void;
}

const useVitalsForm = (v: Vitals) => {
    const [form, setForm] = useState<Vitals>(v);

    const updateFormField = (fieldName: string, value: string) => {
        setForm({
            ...form,
            [fieldName]: value
        });
    };

    return {form, updateFormField};
}

const VitalsPage = ({ permission, admission, updateAdmission }: VitalsProps) => {
    const { form, updateFormField } = useVitalsForm(admission.vitals);
    const [submitAttempted, setSubmitAttempted] = useState(false);

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();

        setSubmitAttempted(true);

        if (
            form.bodyTemperature === "" ||
            form.heartRate === "" ||
            form.bloodPressure === "" ||
            form.respiratoryRate === "" ||
            form.oxygenSaturation === "" ||
            form.painLevel === "" ||
            form.bloodGlucoseLevel === ""
        ) {
            return;
        }

        updateAdmission({
            ...admission,
            vitals: form
        }, {
            timestamp: new Date(),
            action: "Vitals updated",
            message: "Vitals updated",
            details: JSON.stringify(form),
            performedBy: ""
        });
    }

    return (
        <Box component="form" noValidate onSubmit={handleSubmit}>
            <TableContainer component={Paper}>
                <Table>
                    <TableBody>
                        <TableRow key="bt">
                            <TableCell><b>Body Temperature</b></TableCell>
                            <TableCell>
                                <TextField
                                    id="bt"
                                    name="bt"
                                    type="text"
                                    value={form.bodyTemperature}
                                    onChange={(e) => updateFormField("bodyTemperature", e.target.value)}
                                    required
                                    size="small"
                                    InputProps={{
                                        endAdornment: <InputAdornment position="end">°C</InputAdornment>,
                                    }}
                                    disabled={permission !== "WRITE"}
                                    error={submitAttempted && form.bodyTemperature === ""}
                                />
                            </TableCell>
                        </TableRow>
                        <TableRow key="hr">
                            <TableCell><b>Hearth Rate</b></TableCell>
                            <TableCell>
                                <TextField
                                    id="hr"
                                    name="hr"
                                    type="text"
                                    value={form.heartRate}
                                    onChange={(e) => updateFormField("heartRate", e.target.value)}
                                    required
                                    size="small"
                                    InputProps={{
                                        endAdornment: <InputAdornment position="end">bpm</InputAdornment>,
                                    }}
                                    disabled={permission !== "WRITE"}
                                    error={submitAttempted && form.heartRate === ""}
                                />
                            </TableCell>
                        </TableRow>
                        <TableRow key="bp">
                            <TableCell><b>Blood Pressure</b></TableCell>
                            <TableCell>
                                <TextField
                                    id="bloodPressure"
                                    name="bloodPressure"
                                    type="text"
                                    value={form.bloodPressure}
                                    onChange={(e) => updateFormField("bloodPressure", e.target.value)}
                                    required
                                    size="small"
                                    InputProps={{
                                        endAdornment: <InputAdornment position="end">mmHg</InputAdornment>,
                                    }}
                                    disabled={permission !== "WRITE"}
                                    error={submitAttempted && form.bloodPressure === ""}
                                />
                            </TableCell>
                        </TableRow>
                        <TableRow key="rr">
                            <TableCell><b>Respiratory Rate</b></TableCell>
                            <TableCell>
                                <TextField
                                    id="respiratoryRate"
                                    name="respiratoryRate"
                                    type="text"
                                    value={form.respiratoryRate}
                                    onChange={(e) => updateFormField("respiratoryRate", e.target.value)}
                                    required
                                    size="small"
                                    InputProps={{
                                        endAdornment: <InputAdornment position="end">bpm</InputAdornment>,
                                    }}
                                    disabled={permission !== "WRITE"}
                                    error={submitAttempted && form.respiratoryRate === ""}
                                />
                            </TableCell>
                        </TableRow>
                        <TableRow key="os">
                            <TableCell><b>Oxygen Saturation</b></TableCell>
                            <TableCell>
                                <TextField
                                    id="oxygenSaturation"
                                    name="oxygenSaturation"
                                    type="text"
                                    value={form.oxygenSaturation}
                                    onChange={(e) => updateFormField("oxygenSaturation", e.target.value)}
                                    required
                                    size="small"
                                    InputProps={{
                                        endAdornment: <InputAdornment position="end">%</InputAdornment>,
                                    }}
                                    disabled={permission !== "WRITE"}
                                    error={submitAttempted && form.oxygenSaturation === ""}
                                />
                            </TableCell>
                        </TableRow>
                        <TableRow key="bgl">
                            <TableCell><b>Blood Glucose Level</b></TableCell>
                            <TableCell>
                                <TextField
                                    id="bloodGlucoseLevel"
                                    name="bloodGlucoseLevel"
                                    type="text"
                                    value={form.bloodGlucoseLevel}
                                    onChange={(e) => updateFormField("bloodGlucoseLevel", e.target.value)}
                                    required
                                    size="small"
                                    InputProps={{
                                        endAdornment: <InputAdornment position="end">mg/dL</InputAdornment>,
                                    }}
                                    disabled={permission !== "WRITE"}
                                    error={submitAttempted && form.bloodGlucoseLevel === ""}
                                />
                            </TableCell>
                        </TableRow>
                        <TableRow key="pl">
                            <TableCell><b>Pain Level</b></TableCell>
                            <TableCell>
                                <TextField
                                    id="painLevel"
                                    name="painLevel"
                                    type="text"
                                    value={form.painLevel}
                                    onChange={(e) => updateFormField("painLevel", e.target.value)}
                                    required
                                    size="small"
                                    disabled={permission !== "WRITE"}
                                    error={submitAttempted && form.painLevel === ""}
                                />
                            </TableCell>
                        </TableRow>
                    </TableBody>
                </Table>
            </TableContainer>
            <Button
                type="submit"
                variant="contained"
                sx={{ mt: 3, mb: 2 }}
            >
                SAVE
            </Button>
        </Box>
    )
};

export default VitalsPage;
