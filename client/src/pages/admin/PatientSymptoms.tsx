import TextField from "@mui/material/TextField";

export interface PatientSymptomsProps {
    patientSymptoms: string;
    setPatientSymptoms: (value: string) => void;
    patientSymptomsError: boolean;
}

const PatientSymptoms = ({ patientSymptoms, setPatientSymptoms, patientSymptomsError }: PatientSymptomsProps) => {
    return (
        <TextField
            id="outlined-multiline-static"
            label="Patient symptoms"
            fullWidth
            multiline
            rows={10}
            placeholder={"Enter patient symptoms here..."}
            variant="outlined"
            required
            onChange={(e) => setPatientSymptoms(e.target.value)}
            error={patientSymptomsError && patientSymptoms === ""}
            helperText={patientSymptomsError && patientSymptoms === "" ? "This field is required" : ""}
        />
    );
}

export default PatientSymptoms;