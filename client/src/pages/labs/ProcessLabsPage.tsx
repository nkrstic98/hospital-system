import Grid from "@mui/material/Grid";
import {
    Alert, Chip, Collapse,
    TableContainer
} from "@mui/material";
import {Lab} from "../../types/Admission.ts";
import Paper from "@mui/material/Paper";
import Table from "@mui/material/Table";
import TableHead from "@mui/material/TableHead";
import TableCell from "@mui/material/TableCell";
import TableRow from "@mui/material/TableRow";
import Typography from "@mui/material/Typography";
import TableBody from "@mui/material/TableBody";
import Button from "@mui/material/Button";
import {useEffect, useState} from "react";
import {LabService} from "../../services/lab/Lab.ts";
import IconButton from "@mui/material/IconButton";
import KeyboardArrowDownIcon from '@mui/icons-material/KeyboardArrowDown';
import KeyboardArrowUpIcon from '@mui/icons-material/KeyboardArrowUp';
import Box from "@mui/material/Box";

const ProcessLabsPage = () => {
    const labService = new LabService();
    const [labs, setLabs] = useState<Lab[]>([]);
    const [fetchError, setFetchError] = useState<boolean>(false);
    const [open, setOpen] = useState<boolean[]>([]);

    const processLab = (labId: string) => {
        labService.ProcessLabTest(labId).then((success) => {
            if (success) {
                labService.GetLabs().then((labs) => {
                    if (labs !== undefined) {
                        setLabs(labs);
                    }
                }).catch(() => {
                    setFetchError(true);
                });
            }
        }).catch(() => {
            setFetchError(true);
        });
    }

    useEffect(() => {
        labService.GetLabs().then((labs) => {
            if (labs !== undefined) {
                setLabs(labs);
                const openArray = new Array<boolean>(labs.length).fill(false);
                setOpen(openArray);
            }
        }).catch(() => {
            setFetchError(true);
        });
    }, []);

    return (
        <Grid item xs={12} sx={{mt: 10}}>
            {fetchError && <Alert severity="error">Failed to fetch labs</Alert>}
            <Typography variant="h4" gutterBottom component="div" sx={{mb: 3}}>
                Ordered Lab Tests
            </Typography>
            <TableContainer component={Paper}>
                <Table>
                    <TableBody>
                        {labs.map((lab, index) => (
                            <>
                                <TableRow sx={{ '& > *': { borderBottom: 'unset' } }}>
                                    <TableCell>
                                        <IconButton
                                            aria-label="expand row"
                                            size="small"
                                            onClick={() => {
                                                const newOpen = [...open];
                                                newOpen[index] = !open[index];
                                                setOpen(newOpen);
                                            }}
                                        >
                                            {open[index] ? <KeyboardArrowUpIcon /> : <KeyboardArrowDownIcon />}
                                        </IconButton>
                                    </TableCell>
                                    <TableCell align="left">{lab.id}</TableCell>
                                    <TableCell align="left">{lab.testType}</TableCell>
                                    <TableCell align="left">
                                        {lab.processedAt !== null ? (
                                            <Chip label="Processed" color="success" variant="outlined" size="small" />
                                        ) : (
                                            <Button onClick={() => processLab(lab.id)} variant="contained" color="primary" size="small">
                                                Process
                                            </Button>
                                        )}
                                    </TableCell>
                                </TableRow>
                                <TableRow>
                                    <TableCell style={{ paddingBottom: 0, paddingTop: 0 }} colSpan={6}>
                                        <Collapse in={open[index]} timeout="auto" unmountOnExit>
                                            <Box sx={{ margin: 3 }}>
                                                <Typography align="center" variant="h6" gutterBottom component="div" sx={{ mb: 2 }}>
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
                                                                            {result.result < 1 ? "negative" : "positive"}
                                                                        </TableCell>}
                                                                        {result.unit !== "positive/negative" && <TableCell>{result.result} {result.unit}</TableCell>}
                                                                        <TableCell>{result.referenceRange}</TableCell>
                                                                    </TableRow>
                                                                ))}
                                                            </TableBody>
                                                        </Table>
                                                    </TableContainer>
                                                ) : (
                                                    <Alert severity="warning">Click PROCESS button to receive test results!</Alert>
                                                )}
                                            </Box>
                                        </Collapse>
                                    </TableCell>
                                </TableRow>
                            </>
                        ))}
                    </TableBody>
                </Table>
            </TableContainer>
        </Grid>
    )
}

export default ProcessLabsPage;
