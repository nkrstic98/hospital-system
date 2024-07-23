import Grid from "@mui/material/Grid";
import {
    Accordion,
    AccordionDetails,
    AccordionSummary,
    Alert, Chip, FormControl, FormHelperText,
    InputLabel,
    OutlinedInput,
    Select,
    TableContainer
} from "@mui/material";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import {AdmissionDetails} from "../../types/Admission.ts";
import Paper from "@mui/material/Paper";
import Table from "@mui/material/Table";
import TableHead from "@mui/material/TableHead";
import TableCell from "@mui/material/TableCell";
import TableRow from "@mui/material/TableRow";
import Typography from "@mui/material/Typography";
import TableBody from "@mui/material/TableBody";
import Button from "@mui/material/Button";
import {useState} from "react";
import MenuItem from "@mui/material/MenuItem";

const labTypes = [
    "Complete Blood Count (CBC)",
    "Basic Metabolic Panel (BMP)",
    "Comprehensive Metabolic Panel (CMP)",
    "Lipid Panel",
    "Liver Function Tests (LFTs)",
    "Thyroid Function Tests",
    "Coagulation Panel",
    "Urinalysis",
    "Hemoglobin A1c (HbA1c)",
    "Blood Culture",
    "C-Reactive Protein (CRP)",
    "Erythrocyte Sedimentation Rate (ESR)",
    "Vitamin and Mineral Tests",
    "Hormone Tests",
    "Infectious Disease Tests",
    "Tumor Markers"
]

export interface LabsPageProps {
    admission: AdmissionDetails;
    orderLabTest: (labTest: string, admissionId: string) => void;
    sections: Map<string, string>;
}

const LabsPage = ({ admission, orderLabTest, sections }: LabsPageProps) => {
    const readLabs = sections.get("PATIENTS:LABS:RESULT") !== undefined && sections.get("PATIENTS:LABS:RESULT") == "READ";
    const orderLabs = sections.get("PATIENTS:LABS:ORDER") !== undefined && sections.get("PATIENTS:LABS:ORDER") == "WRITE";
    const [labTest, setLabTest] = useState<string>("");
    const [submitAttempted, setSubmitAttempted] = useState(false);

    const getTime = (time: Date) => {
        const date = new Date(time);
        return date.toLocaleDateString() + " " + date.toLocaleTimeString();
    }

    const order = () => {
        if (labTest === "") {
            setSubmitAttempted(true);
            return;
        }

        orderLabTest(labTest, admission.id)

        setLabTest("");
        setSubmitAttempted(false);
    }

    return (
        <Grid item xs={12} sx={{mt: 2}}>
            {orderLabs && (
                <Grid container spacing={3} sx={{mb:5}}>
                    <Grid item xs={9}>
                        <FormControl fullWidth required size="small">
                            <InputLabel id="dept-label">Choose Lab Test</InputLabel>
                            <Select
                                labelId="labs"
                                id="labs"
                                value={labTest}
                                onChange={(e) => setLabTest(e.target.value)}
                                input={<OutlinedInput label="Choose Lab Test" />}
                                required
                                fullWidth
                                error={submitAttempted && labTest === ""}
                            >
                                {
                                    labTypes.map((value) => (
                                        <MenuItem key={value} value={value}>{value}</MenuItem>
                                    ))
                                }
                            </Select>
                            {submitAttempted && labTest === "" && <FormHelperText error>This field is required</FormHelperText>}
                        </FormControl>
                    </Grid>
                    <Grid item xs={3}>
                        <Button variant="contained" onClick={() => order()}>Order Test</Button>
                    </Grid>
                </Grid>
            )}

            {readLabs && admission.labs !== null && admission.labs.map((lab) => (
                <Accordion>
                    <AccordionSummary
                        expandIcon={<ExpandMoreIcon/>}
                        aria-controls="panel1-content"
                        id="panel1-header"
                    >
                        <b>&nbsp; {lab.id} - {lab.testType}</b> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
                        {lab.processedAt !== null ? (
                            <Chip label="Processed" color="success" variant="outlined" size="small" />
                        ) : (
                            <Chip label="In Process" color="warning" variant="outlined" size="small" />
                        )}
                    </AccordionSummary>
                    <AccordionDetails>
                        <Grid container sx={{pl: 3}}>
                            <Grid item xs={3}>
                                <p><b>Request Time: </b></p>
                            </Grid>
                            <Grid item xs={8}>
                                <p>{getTime(lab.requestedAt)}</p>
                            </Grid>
                            {lab.processedAt !== null && lab.processedAt !== undefined && (
                                <>
                                    <Grid item xs={3}>
                                        <p><b>Process Time: </b></p>
                                    </Grid>
                                    <Grid item xs={8}>
                                        <p>{getTime(lab.processedAt)}</p>
                                    </Grid>
                                </>
                            )}
                            <Grid item xs={3}>
                                <p><b>Requested By: </b></p>
                            </Grid>
                            <Grid item xs={8}>
                                <p>{lab.requestedBy}</p>
                            </Grid>
                            {lab.processedBy !== null && lab.processedBy !== undefined && (
                                <>
                                    <Grid item xs={3}>
                                        <p><b>Processed By: </b></p>
                                    </Grid>
                                    <Grid item xs={8}>
                                        <p>{lab.processedBy}</p>
                                    </Grid></>
                            )}
                        </Grid>
                        <Typography variant="h6" gutterBottom component="div" sx={{m: 4}}>
                            Test Results
                        </Typography>
                        {lab.testResults !== null ? (
                            <TableContainer component={Paper}>
                                <Table>
                                    <TableHead>
                                        <TableRow>
                                            <TableCell><b>Parameter</b></TableCell>
                                            <TableCell><b>Measured Value</b></TableCell>
                                            <TableCell><b>Reference Range</b></TableCell>
                                        </TableRow>
                                    </TableHead>
                                    <TableBody>
                                        {lab.testResults !== undefined && lab.testResults.map((result, index) => (
                                            <TableRow key={index}>
                                                <TableCell>{result.name}</TableCell>
                                                {result.unit === "positive/negative" && <TableCell>
                                                    {result.result < 1 ? "Negative" : "Positive"}
                                                </TableCell>}
                                                {result.unit !== "positive/negative" && <TableCell>{result.result} {result.unit}</TableCell>}
                                                <TableCell>{result.referenceRange}</TableCell>
                                            </TableRow>
                                        ))}
                                    </TableBody>
                                </Table>
                            </TableContainer>
                        ) : (
                            <Alert severity="info">Waiting for a test result.</Alert>
                        )}
                    </AccordionDetails>
                </Accordion>
            ))}
        </Grid>
    )
}

export default LabsPage;