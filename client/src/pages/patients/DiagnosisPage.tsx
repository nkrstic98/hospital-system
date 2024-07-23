import TextField from "@mui/material/TextField";
import {useState} from "react";
import {AdmissionDetails, Log} from "../../types/Admission.ts";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";

export interface DiagnosisPageProps {
    permission: string;
    admission: AdmissionDetails;
    updateAdmission: (admission: AdmissionDetails, log: Log) => void;
}

const DiagnosisPage = ({permission, admission, updateAdmission}: DiagnosisPageProps) => {
    const [diagnosis, setDiagnosis] = useState<string>(admission?.diagnosis ?? "");
    const [submitAttempted, setSubmitAttempted] = useState<boolean>(false);

    const submit = () => {
        setSubmitAttempted(true);
        if (diagnosis === "") {
            return;
        }

        updateAdmission({
            ...admission,
            diagnosis: diagnosis
        }, {
            timestamp: new Date(),
            action: "Update diagnosis",
            message: "Diagnosis updated: " + diagnosis,
            details: diagnosis?? "",
            performedBy: ""
        });

        setSubmitAttempted(false);
    }

    return (
        <>
            <Typography variant="h6" gutterBottom component="div" sx={{mb: 3}}>
                Patient Diagnosis
            </Typography>
            <TextField
                id="outlined-multiline-static"
                label="Diagnosis"
                value={diagnosis}
                fullWidth
                multiline
                rows={15}
                placeholder="Enter diagnosis..."
                variant="outlined"
                required
                onChange={(e) => setDiagnosis(e.target.value)}
                error={submitAttempted && diagnosis === ""}
                helperText={submitAttempted && diagnosis === "" ? "This field is required" : ""}
                disabled={permission !== "WRITE"}
            />
            {permission == "WRITE" && <Button variant="contained" onClick={() => submit()} sx={{mt: 4}}>Save Diagnosis</Button>}
        </>
    )
}

export default DiagnosisPage;