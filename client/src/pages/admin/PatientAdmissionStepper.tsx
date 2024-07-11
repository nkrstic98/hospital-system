import * as React from 'react';
import Box from '@mui/material/Box';
import Stepper from '@mui/material/Stepper';
import Step from '@mui/material/Step';
import StepLabel from '@mui/material/StepLabel';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import CssBaseline from "@mui/material/CssBaseline";
import TextField from "@mui/material/TextField";
import Container from "@mui/material/Container";
import {useCallback, useEffect, useState} from "react";
import {PatientService} from "../../services/patient/Patient.ts";
import {Patient} from "../../types/Patient.ts";
import PatientInformation from "./PatientInformation.tsx";
import PatientRegister from "./PatientRegister.tsx";
import {Alert, Modal} from "@mui/material";
import PatientAllergies from "./PatientAllergies.tsx";
import PatientMedications from "./PatientMedications.tsx";
import PatientSymptoms from "./PatientSymptoms.tsx";
import DepartmentsAndPhysicians from "./DepartmentsAndPhysicians.tsx";
import {Department} from "../../services/department/types.ts";
import {useNavigate} from "react-router-dom";
import {DepartmentService} from "../../services/department/Department.ts";

const style = {
    position: 'absolute' as const,
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

const steps = ["Check Patient Existence", "Patient Information", "Enter Symptoms", "Enter Medications", "Enter Allergies", "Choose Department and Physician", "Summary"]

export type RegisterPatientFormFields = {
    firstname: string;
    lastname: string;
    nationalIdentificationNumber: string;
    medicalRecordNumber: string;
    email: string;
    phoneNumber: string;
}

const useRegisterPatientForm = () => {
    const [form, setForm] = useState<RegisterPatientFormFields>({
        firstname: "",
        lastname: "",
        nationalIdentificationNumber: "",
        medicalRecordNumber: "",
        email: "",
        phoneNumber: "",
    });

    const clearField = (fieldName: string) => {
        setForm({
            ...form,
            [fieldName]: ''
        });
    };

    const updateFormField = (fieldName: string, value: string) => {
        setForm({
            ...form,
            [fieldName]: value
        });
    };

    const resetForm = () => {
        setForm({
            firstname: "",
            lastname: "",
            nationalIdentificationNumber: "",
            medicalRecordNumber: "",
            email: "",
            phoneNumber: "",
        });
    }

    return { form, updateFormField, clearField, resetForm };
}

export default function PatientAdmissionStepper() {
    const navigate = useNavigate();

    const patientService = new PatientService();
    const departmentService = new DepartmentService();

    const [patientId, setPatientId] = useState("");
    const [findPatientAttempted, setFindPatientAttempted] = useState(false);
    const [patient, setPatient] = useState<Patient|undefined>(undefined);

    const [patientAdmitted, setPatientAdmitted] = useState(false);

    const { form, updateFormField, resetForm } = useRegisterPatientForm();
    const [userRegisterAttempted, setUserRegisterAttempted] = useState(false);
    const [patientRegisterError, setPatientRegisterError] = useState(false);

    const [patientSymptomsError, setPatientSymptomsError] = useState(false);
    const [patientSymptoms, setPatientSymptoms] = useState<string>("");

    const [patientMedications, setPatientMedications] = useState<string[]>([]);
    const [checkedAllergies, setCheckedAllergies] = useState<string[]>([]);

    const [departmentPhysicians, setDepartmentPhysicians] = useState<Map<string, Department> | undefined>(undefined);

    const [department , setDepartment] = useState<string>("");
    const [departmentError, setDepartmentError] = useState<boolean>(false);

    const [physician, setPhysician] = useState<string>("");
    const [physicianName, setPhysicianName] = useState<string>("");
    const [physicianError, setPhysicianError] = useState<boolean>(false);

    const [activeStep, setActiveStep] = React.useState(0);
    const [skipped, setSkipped] = React.useState(new Set<number>());

    const [successfulAdmission, setSuccessfulAdmission] = useState(false);
    const [submissionFailure, setSubmissionFailure] = useState(false);

    const isStepOptional = (step: number) => {
        return step === 0 || step === 3 || step == 4;
    };

    const isStepSkipped = (step: number) => {
        return skipped.has(step);
    };

    const handleNext = () => {
        const as = activeStep;

        let newSkipped = skipped;
        if (isStepSkipped(activeStep)) {
            newSkipped = new Set(newSkipped.values());
            newSkipped.delete(activeStep);
        }

        setActiveStep((prevActiveStep) => prevActiveStep + 1);
        setSkipped(newSkipped);

        if (as === steps.length - 1) {
            handleFinish();
        }
    };

    const handleBack = () => {
        setActiveStep((prevActiveStep) => prevActiveStep - 1);
    };

    const handleSkip = () => {
        if (!isStepOptional(activeStep)) {
            throw new Error("You can't skip a step that isn't optional.");
        }

        if (activeStep == 3) {
            setPatientMedications([]);
        }
        if (activeStep == 4) {
            setCheckedAllergies([]);

        }

        setActiveStep((prevActiveStep) => prevActiveStep + 1);
        setSkipped((prevSkipped) => {
            const newSkipped = new Set(prevSkipped.values());
            newSkipped.add(activeStep);
            return newSkipped;
        });
    };

    const handleFinish = () => {
        patientService.RegisterPatientAdmission({
            patientId: patient?.id as string,
            department: department,
            physician: physician,
            symptoms: patientSymptoms,
            medications: patientMedications,
            allergies: checkedAllergies,
        }).then((r) => {
            if (!r) {
                setSubmissionFailure(true);
                setActiveStep((prevActiveStep) => prevActiveStep - 1);
                setTimeout(() => {
                    setSubmissionFailure(false);
                }, 2000);
                return;
            }
            setSuccessfulAdmission(true);
        })
    };

    const getStepContent = (step: number) => {
        switch (step) {
            case 0:
                return <TextField
                    sx={{ml: 5, mr: 5}}
                    id="patientId"
                    name="patientId"
                    type="text"
                    label="Enter Patient National or Medical ID"
                    value={patientId}
                    onChange={(e) => setPatientId(e.target.value)}
                    margin="normal"
                    required
                    fullWidth
                    autoFocus
                    error={findPatientAttempted && patientId === ""}
                    helperText={findPatientAttempted && patientId === "" ? "This field is required" : ""}
                />;
            case 1:
                if (patient !== undefined) {
                    return <PatientInformation patient={patient}/>
                }

                return <>
                    {!patientRegisterError &&
                        <Alert severity="warning">Patient is not in the system! Register it in order to
                            continue.</Alert>}
                    {patientRegisterError &&
                        <Alert variant="filled" severity="error" sx={{m: 3}}>Register failed. Check your data and try
                            again!</Alert>}
                    <PatientRegister
                        form={form}
                        updateFormField={updateFormField}
                        userRegisterAttempted={userRegisterAttempted}/>
                </>;
            case 2:
                return <PatientSymptoms patientSymptoms={patientSymptoms} setPatientSymptoms={setPatientSymptoms}
                                        patientSymptomsError={patientSymptomsError}/>
            case 3:
                return <PatientMedications handleMedicationsChange={setPatientMedications}/>;
            case 4:
                return <PatientAllergies checkedAllergies={checkedAllergies}
                                         handleCheckboxChange={handleSelectAllergiesChange}/>;
            case 5:
                return <DepartmentsAndPhysicians
                    departmentPhysicians={departmentPhysicians}
                    department={department}
                    physician={physician}
                    setDepartment={setDepartment}
                    setPhysicianName={setPhysicianName}
                    setPhysicianId={setPhysician}
                    departmentError={departmentError}
                    physicianError={physicianError}
                />;
            case 6:
                return <div>
                    <h3>Please check data before proceeding!</h3>
                    <p><b>Chosen Department:</b> {departmentPhysicians?.get(department)?.displayName}</p>
                    <p><b>Chosen Physician: </b>{physicianName}</p>
                    <p><b>Admission Date: </b>{new Date().toLocaleDateString()}</p>
                    <p><b>Symptoms: </b> {patientSymptoms}</p>
                    <p><b>Medications: </b> {patientMedications.length > 0 ? patientMedications.join(", ") : "/"}</p>
                    <p><b>Allergies: </b>{checkedAllergies.length > 0 ? checkedAllergies.join(", ") : "/"}</p>
                </div>
            default:
                return "Unknown step";
        }
    }

    const getHandleNextFunction = (step: number) => {
        switch (step) {
            case 0:
                return handleFindPatient;
            case 1:
                if (patient === undefined) {
                    return handlePatientRegister;
                } else {
                    return handleNext;
                }
            case 2:
                return handlePatientSymptoms;
            case 5:
                return handleDepartmentsAndPhysicians;
            default:
                return handleNext;
        }
    }

    const handleFindPatient = () => {
        if (patientId === "") {
            setFindPatientAttempted(true);
            return;
        }

        patientService.GetPatient(patientId).then(r => {
            if (r !== undefined) {
                const activeAdmissions  = r.admissions.filter(
                    (admission) => admission.status === "admitted" || admission.status === "pending"
                );
                if (activeAdmissions.length > 0) {
                    setPatientAdmitted(true);
                    return;
                }
            }

            setPatient(r);
            setFindPatientAttempted(false);
            setPatientId("");
        }).finally(() => {
            if (!patientAdmitted) {
                handleNext();
            }
        });
    }

    const handlePatientRegister = useCallback(() => {
        if (form.firstname == "" ||
            form.lastname == "" ||
            form.nationalIdentificationNumber == "" ||
            form.medicalRecordNumber == "" ||
            form.email == "" ||
            form.phoneNumber == "") {
            setUserRegisterAttempted(true);
            return;
        }

        setUserRegisterAttempted(false);
        patientService.Register({
            firstname: form.firstname,
            lastname: form.lastname,
            nationalIdentificationNumber: form.nationalIdentificationNumber,
            medicalRecordNumber: form.medicalRecordNumber,
            email: form.email,
            phoneNumber: form.phoneNumber,
        }).then(r => {
            if (!r) {
                setPatientRegisterError(true);
            } else {
                setPatientRegisterError(false);
                setPatient(r);
                handleNext();
                resetForm();
            }
        });
    }, [form]);

    const handlePatientSymptoms = () => {
        if (patientSymptoms === "") {
            setPatientSymptomsError(true);
        } else {
            setPatientSymptomsError(false);
            handleNext();
        }
    }

    const handleSelectAllergiesChange = (value: string, checked: boolean) => {
        if (checked) {
            setCheckedAllergies([...checkedAllergies, value]);
        } else {
            setCheckedAllergies(checkedAllergies.filter((item) => item !== value));
        }
    };

    const handleDepartmentsAndPhysicians = () => {
        let fieldsValid = true;

        if (department === "") {
            setDepartmentError(true);
            fieldsValid = false;
        }

        if (physician === "") {
            setPhysicianError(true);
            fieldsValid = false;
        }

        if (!fieldsValid) {
            return;
        }

        setDepartmentError(false);
        setPhysicianError(false);
        handleNext();
    }

    useEffect(() => {
        departmentService.GetDepartments().then((data) => {
            setDepartmentPhysicians(data);
        });
    }, []);

    return (
        <Box sx={{ width: '100%', mt: 10 }}>
            {submissionFailure && <Alert variant="filled" severity="error" sx={{ m: 3 }}>Patient Intake failed! Please check data and try again.</Alert>}
            <Stepper activeStep={activeStep} alternativeLabel>
                {steps.map((label, index) => {
                    const stepProps: { completed?: boolean } = {};
                    const labelProps: {
                        optional?: React.ReactNode;
                    } = {};
                    if (isStepOptional(index)) {
                        labelProps.optional = (
                            <Typography variant="caption">Optional</Typography>
                        );
                    }
                    if (isStepSkipped(index)) {
                        stepProps.completed = false;
                    }
                    return (
                        <Step key={label} {...stepProps}>
                            <StepLabel {...labelProps}>{label}</StepLabel>
                        </Step>
                    );
                })}
            </Stepper>
            {activeStep === steps.length ? (
                <React.Fragment>
                    <Typography sx={{mt: 2, mb: 1}}>
                        You have successfully registered patient visit!
                    </Typography>
                    <div>
                        <h3>Patient Intake Summary</h3>
                        <p><b>Chosen Department:</b> {departmentPhysicians?.get(department)?.displayName}</p>
                        <p><b>Chosen Physician: </b>{physicianName}</p>
                        <p><b>Admission Date: </b>{new Date().toLocaleDateString()}</p>
                        <p><b>Symptoms: </b> {patientSymptoms}</p>
                        <p><b>Medications: </b> {patientMedications.length > 0 ? patientMedications.join(", ") : "/"}</p>
                        <p><b>Allergies: </b>{checkedAllergies.length > 0 ? checkedAllergies.join(", ") : "/"}</p>
                    </div>
                    <Box sx={{display: 'flex', flexDirection: 'row', pt: 2}}>
                        <Box sx={{flex: '1 1 auto'}}/>
                        <Button onClick={handleFinish}>Back to Patient Intake Dashboard</Button>
                    </Box>
                </React.Fragment>
            ) : (
                <React.Fragment>
                    <Container component="main" maxWidth="sm">
                        <CssBaseline/>
                        <Box
                            sx={{
                                marginTop: 8,
                                display: 'flex',
                                flexDirection: 'column',
                                alignItems: 'center',
                                marginBottom: 10
                            }}
                        >
                        {getStepContent(activeStep)}
                        </Box>
                    </Container>
                    <Box sx={{ display: 'flex', flexDirection: 'row', pt: 2 }}>
                        <Button
                            color="inherit"
                            disabled={activeStep === 0}
                            onClick={handleBack}
                            sx={{ mr: 1 }}
                        >
                            Back
                        </Button>
                        <Box sx={{ flex: '1 1 auto' }} />
                        {isStepOptional(activeStep) && (
                            <Button color="inherit" onClick={handleSkip} sx={{ mr: 1 }}>
                                Skip
                            </Button>
                        )}
                        <Button onClick={getHandleNextFunction(activeStep)}>
                            {activeStep === steps.length - 1 ? 'Finish' : 'Next'}
                        </Button>
                    </Box>
                </React.Fragment>
            )}

            <Modal
                open={successfulAdmission}
                onClose={() => navigate("/admin")}
                aria-labelledby="parent-modal-title"
                aria-describedby="parent-modal-description"
            >
                <Box sx={{...style, width: 400}}>
                    <h2 id="parent-modal-title">Success!</h2>
                    <p id="parent-modal-description">
                        Patient admitted successfully!
                    </p>
                    <p id="parent-modal-description">
                        You can now go back to the Patient Intake page.
                    </p>
                    <Button
                        variant="text"
                        onClick={() => navigate("/admin")}
                    >
                        Go Back to Dashboard
                    </Button>
                </Box>
            </Modal>

            <Modal
                open={patientAdmitted}
                onClose={() => navigate("/admin")}
                aria-labelledby="parent-modal-title"
                aria-describedby="parent-modal-description"
            >
                <Box sx={{...style, width: 400}}>
                    <h2 id="parent-modal-title">Failure!</h2>
                    <p id="parent-modal-description">
                        Patient with this ID is already admitted, or waiting to be admitted.
                    </p>
                    <p id="parent-modal-description">
                        If you want to admit the patient again, please discharge the patient first.
                    </p>
                    <Button
                        variant="text"
                        onClick={() => navigate("/admin")}
                    >
                        Return to Patient Intake
                    </Button>
                </Box>
            </Modal>
        </Box>
    );
}