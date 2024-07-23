import Box from "@mui/material/Box";
import Paper from "@mui/material/Paper";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableRow from "@mui/material/TableRow";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import {AdmissionDetails, Log, Medication} from "../../types/Admission.ts";
import TableHead from "@mui/material/TableHead";
import Button from "@mui/material/Button";
import {useState} from "react";
import {FormControl, InputLabel, OutlinedInput, Select} from "@mui/material";
import MenuItem from "@mui/material/MenuItem";
import TextField from "@mui/material/TextField";

const commonDrugs: Array<string> = [
    "Paracetamol (Acetaminophen)",
    "Ibuprofen",
    "Aspirin (Acetylsalicylic Acid)",
    "Loratadine",
    "Cetirizine",
    "Omeprazole",
    "Simvastatin",
    "Metformin",
    "Salbutamol (Albuterol)",
    "Amoxicillin",
    "Amlodipine",
    "Lisinopril",
    "Atorvastatin",
    "Warfarin",
    "Metoprolol",
    "Clopidogrel",
    "Levothyroxine",
    "Pantoprazole",
    "Diazepam",
    "Tramadol",
    "Citalopram",
    "Fluoxetine",
    "Gabapentin",
    "Sertraline",
    "Hydrochlorothiazide",
    "Montelukast",
    "Prednisone",
    "Alprazolam",
    "Methotrexate",
    "Aspirin-Dipyridamole"
];

export interface MedicationsPageProps {
    admission: AdmissionDetails;
    updateAdmission: (admission: AdmissionDetails, log: Log) => void;
    sections: Map<string, string>;
}

const MedicationsPage = ({ admission, updateAdmission, sections }: MedicationsPageProps) => {
    const giveMedication = sections.get("PATIENTS:MEDICATIONS:GIVE") !== undefined && sections.get("PATIENTS:MEDICATIONS:GIVE") == "WRITE";
    const prescribeMedication = sections.get("PATIENTS:MEDICATIONS:PRESCRIBE") !== undefined && sections.get("PATIENTS:MEDICATIONS:PRESCRIBE") == "WRITE";
    const [medications, setMedication] = useState<Medication[]>([]);
    const [submitAttempted, setSubmitAttempted] = useState(false);

    const addNewMedication = () => {
        setMedication([...medications, {medication: "", dose: ""}]);
    }

    const setMedicationName = (value: string, index: number) => {
        const newMeds = [...medications];
        newMeds[index].medication = value;
        setMedication(newMeds);
    }

    const setMedicationDose = (value: string, index: number) => {
        const newMeds = [...medications];
        newMeds[index].dose = value;
        setMedication(newMeds);
    }

    const removeMedication = (index: number) => {
        const newMeds = [...medications];
        newMeds.splice(index, 1);
        setMedication(newMeds);
    }

    const giveDose = (index: number) => {
        if (admission.medications === null) {
            return;
        }

        const medication = admission.medications[index];

        updateAdmission({
            ...admission,
        }, {
            timestamp: new Date(),
            action: "Medication given",
            message: "Medication given: " + medication.medication + ", " + medication.dose,
            details: JSON.stringify(medication),
            performedBy: ""
        });
    }

    const removePrescribedMedication = (medication: Medication, index: number) => {
        const meds = [...medications];
        meds.splice(index, 1);
        updateAdmission({
            ...admission,
            medications: meds,
        }, {
            timestamp: new Date(),
            action: "Medication removed",
            message: "Medication removed: " + medication.medication,
            details: JSON.stringify(medications),
            performedBy: ""
        });
    }

    const submit = () => {
        setSubmitAttempted(true);

        if (medications.some((medication) => medication.medication === "" || medication.dose === "")) {
            return;
        }

        const meds = admission.medications !== null ? admission.medications : [];
        meds.push(...medications);

        updateAdmission({
            ...admission,
            medications: meds,
        }, {
            timestamp: new Date(),
            action: "Medications prescribed",
            message: "Medications prescribed: " + medications[0].medication,
            details: JSON.stringify(medications),
            performedBy: ""
        });

        setMedication([]);
        setSubmitAttempted(false);
    }

    return (
        <Box>
            <TableContainer component={Paper}>
                <Table>
                    <TableHead>
                        <TableRow>
                            <TableCell colSpan={2 + (giveMedication || prescribeMedication ? 1 : 0)} align="center" sx={{ fontSize: "20px"}}>
                                Prescribed Medications
                            </TableCell>
                        </TableRow>
                        {admission.medications !== null && admission.medications.length > 0 && <TableRow>
                            <TableCell sx={{fontSize: "18px"}}>Medication</TableCell>
                            <TableCell sx={{fontSize: "18px"}}>Dose</TableCell>
                            <TableCell></TableCell>
                        </TableRow>}
                    </TableHead>
                    <TableBody>
                        {admission.medications !== null && admission.medications.map((medication, index) => (
                            <TableRow key="cc">
                                <TableCell><b>{medication.medication}</b></TableCell>
                                <TableCell>{medication.dose}</TableCell>
                                {<TableCell align="right">
                                    {giveMedication &&
                                        <Button variant="outlined" color="success" size="small" sx={{mr: 2}} onClick={() => giveDose(index)}>Give Dose</Button>}
                                    {prescribeMedication &&
                                        <Button variant="outlined" color="error" size="small" onClick={() => {removePrescribedMedication(medication, index)}}>Remove</Button>}
                                </TableCell>}
                            </TableRow>
                        ))}
                        {prescribeMedication  && medications.map((medication: Medication, index: number) => (
                            <TableRow key={`prescribeMeds${index}`}>
                                <TableCell>
                                    <FormControl fullWidth required size="small" sx={{minWidth: "200px"}}>
                                        <InputLabel id="dept-label">Choose Medication</InputLabel>
                                        <Select
                                            labelId={`prescribeMeds${index}`}
                                            id={`prescribeMeds${index}`}
                                            value={medication.medication}
                                            onChange={(e) => setMedicationName(e.target.value, index)}
                                            input={<OutlinedInput label="Choose Medication" />}
                                            required
                                            fullWidth
                                            error={submitAttempted && medication.medication === ""}
                                        >
                                            {
                                                commonDrugs.map((value) => (
                                                    <MenuItem key={value} value={value}>{value}</MenuItem>
                                                ))
                                            }
                                        </Select>
                                    </FormControl>
                                </TableCell>
                                <TableCell>
                                    <TextField
                                        id={`prescribeMedsDose${index}`}
                                        name={`prescribeMedsDose${index}`}
                                        type="text"
                                        label="Dose"
                                        value={medication.dose}
                                        onChange={(e) => setMedicationDose(e.target.value, index)}
                                        required
                                        fullWidth
                                        size="small"
                                        error={submitAttempted && medication.dose === ""}
                                    />
                                </TableCell>
                                <TableCell align="right">
                                    <Button variant="outlined" color="error" size="small" onClick={() => removeMedication(index)}>Remove</Button>
                                </TableCell>
                            </TableRow>
                        ))}
                        {prescribeMedication && <TableRow key="prescribeMeds">
                            <TableCell colSpan={2 + (giveMedication ? 1 : 0)} align="center" sx={{fontSize: "20px"}}>
                                {
                                    medications.length === 0 ?
                                        <Button variant="contained" onClick={() => addNewMedication()}>Prescribe New Medication</Button>
                                        :
                                        <Button variant="contained" color="warning" onClick={() => submit()}>Save Prescription</Button>
                                }
                            </TableCell>
                        </TableRow>}
                    </TableBody>
                </Table>
            </TableContainer>
        </Box>
    );
}

export default MedicationsPage;
