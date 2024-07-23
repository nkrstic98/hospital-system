import Box from '@mui/material/Box';
import Drawer from '@mui/material/Drawer';
import CssBaseline from '@mui/material/CssBaseline';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import ListItem from '@mui/material/ListItem';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import React, {useEffect, useState} from "react";
import {
    // AdfScanner,
    Bloodtype, CompareArrows, ExitToApp, Groups,
    // History,
    InsertDriveFile, ListAlt,
    Medication,
    MonitorHeart,
    Person,
    // PersonAddAlt1
} from "@mui/icons-material";
import {useNavigate, useParams} from "react-router-dom";
import {PatientService} from "../../services/patient/Patient.ts";
import {AdmissionDetails, Log} from "../../types/Admission.ts";
import Button from "@mui/material/Button";
import {Alert, Chip, CircularProgress, Grid, List, Modal} from "@mui/material";
import {User} from "../../types/User.ts";
import {useAuth} from "../../router/AuthProvider.tsx";
import AnamnesisPage from "./AnamnesisPage.tsx";
import VitalsPage from "./VitalsPage.tsx";
import MedicationsPage from "./MedicationsPage.tsx";
import LabsPage from "./LabsPage.tsx";
import PatientTransferPage from "./PatientTransferPage.tsx";
import DiagnosisPage from "./DiagnosisPage.tsx";
import LogsPage from "./LogsPage.tsx";
import CareTeamPage from "./CareTeamPage.tsx";
import DischargePage from "./DischargePage.tsx";

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

const drawerWidth = 250;

type Section = {
    name: string,
    icon: React.ReactNode,
    hasSubsection: boolean;
}

const infoSections: Map<string, Section> = new Map([
    ["PATIENTS:INFO", {name: "Anamnesis", icon: <Person />, hasSubsection: false}],
    // ["PATIENTS:HISTORY", {name: "Patient History", icon: <History />, hasSubsection: false}],

    ["PATIENTS:VITALS", {name: "Vitals", icon: <MonitorHeart />, hasSubsection: false}],
    ["PATIENTS:MEDICATIONS", {name: "Medications", icon: <Medication />, hasSubsection: true}],
    ["PATIENTS:LABS", {name: "Labs", icon: <Bloodtype />, hasSubsection: true}],
    // ["PATIENTS:IMAGING", {name: "Imaging", icon: <AdfScanner />, hasSubsection: true}],
    // ["PATIENTS:CONSULTS", {name: "Consults", icon: <PersonAddAlt1 />, hasSubsection: false}],
    ["PATIENTS:DIAGNOSIS", {name: "Diagnosis", icon: <InsertDriveFile />, hasSubsection: false}],

    ["PATIENTS:LOGS", {name: "Logs", icon: <ListAlt />, hasSubsection: false}],
    ["PATIENTS:TEAM", {name: "Care Team", icon: <Groups />, hasSubsection: false}],

    ["PATIENTS:TRANSFER", {name: "Transfer Patient", icon: <CompareArrows />, hasSubsection: false}],
    ["PATIENTS:DISCHARGE", {name: "Discharge Patient", icon: <ExitToApp />, hasSubsection: false}],
]);

interface AdmissionDetailsParams {
    id: string;
}

const AdmissionDetailsPage = () => {
    const { user } = useAuth();
    const navigate = useNavigate();
    const patientService = new PatientService();
    const { id } = useParams<keyof AdmissionDetailsParams>();
    const [filteredInfoSection, setFilteredInfoSection] = useState(infoSections);
    const [admission, setAdmission] = useState<AdmissionDetails|undefined>(undefined);
    const [fetchError, setFetchError] = useState(false);

    const [permissions, setPermissions] = useState<Map<string, string>>(new Map());
    const [chosenSection, setChosenSection] = useState<string>("PATIENTS:INFO");
    const [chosenSectionPermission, setChosenSectionPermission] = useState<string>("");
    const [chosenSubsections, setChosenSubsections] = useState<Map<string, string>>(new Map());

    const [updateSuccess, setUpdateSuccess] = useState(false);
    const [updateFail, setUpdateFail] = useState(false);

    const filterSections = (section: Map<string, Section>, permissionKeys: string[]) => {
        const mySections = new Map(section);
        for (const key of mySections.keys()) {
            let filterKey = true;
            for (const pk of permissionKeys) {
                if (pk.includes(key)) {
                    filterKey = false;
                    break;
                }
            }
            if (filterKey) {
                mySections.delete(key);
            }
        }

        setFilteredInfoSection(mySections);
    }

    const chooseSection = (section: string, hasSubsection: boolean) => {
        setChosenSection(section);

        if (hasSubsection) {
            const subsections = new Map<string, string>();
            for (const key of permissions.keys()) {
                if (key.includes(section)) {
                    subsections.set(key, permissions.get(key)!);
                }
            }
            setChosenSubsections(subsections);
        } else {
            setChosenSectionPermission(permissions.get(section)!);
        }
    }

    const getPage = () => {
        if (admission == undefined) {
            return (
                <Grid container xs={12} sx={{mt: 2}}>
                    <Grid item xs={3}>
                        <p><b>Patient: </b></p>
                    </Grid>
                    <Grid item xs={8}>
                        <p>Not Found</p>
                    </Grid>
                </Grid>
            );
        }

        switch (chosenSection) {
            case "PATIENTS:INFO":
                return <AnamnesisPage anamnesis={admission.anamnesis} />;
            case "PATIENTS:VITALS":
                return <VitalsPage permission={chosenSectionPermission} admission={admission} updateAdmission={updateAdmission} />
            case "PATIENTS:MEDICATIONS":
                return <MedicationsPage admission={admission} updateAdmission={updateAdmission} sections={chosenSubsections} />
            case "PATIENTS:LABS":
                return <LabsPage admission={admission} orderLabTest={orderLabTest} sections={chosenSubsections} />
            case "PATIENTS:DIAGNOSIS":
                return <DiagnosisPage permission={chosenSectionPermission} admission={admission} updateAdmission={updateAdmission} />
            case "PATIENTS:LOGS":
                return <LogsPage permission={chosenSectionPermission} logs={admission.logs} />
            case "PATIENTS:TEAM":
                return <CareTeamPage
                    careTeam={admission.careTeam}
                    permission={chosenSectionPermission}
                    addTeamMember={addTeamMember}
                    addTeamMemberPermissions={addTeamMemberPermissions}
                    removeTeamMember={removeTeamMember}
                    removeTeamMemberPermissions={removeTeamMemberPermissions}
                />
            case "PATIENTS:TRANSFER":
                return <PatientTransferPage permission={chosenSectionPermission} requestTransfer={requestTransfer}/>
            case "PATIENTS:DISCHARGE":
                return <DischargePage permission={chosenSectionPermission} canDischarge={admission.diagnosis !== null} dischargePatient={dischargePatient} />
        }
    }

    const dischargePatient = () => {
        patientService.DischargePatient(admission!.id).then(() => {
            navigate("/patients/admissions");
        }).catch(() => {
            setUpdateFail(true);
        });
    }

    const updateAdmission = (a: AdmissionDetails, log: Log) => {
        log.performedBy = (user?.firstname ?? "") + " " + (user?.lastname ?? "");
        a = {...a, logs: [log, ...a.logs]};
        setAdmission(a);

        patientService.UpdateAdmissionDetails(a).then((res) => {
            if (res == undefined) {
                setUpdateFail(true);
                setTimeout(() => {
                    setUpdateFail(false);
                }, 2000);
                return;
            }

            setAdmission(res);
            setUpdateSuccess(true);
            setTimeout(() => {
                setUpdateSuccess(false);
            }, 2000);
        }).catch(() => {
            setUpdateFail(true);
        });
    }

    const requestTransfer = (department: string, doctor: string) => {
        if (admission == undefined) {
            setUpdateFail(true);
            setTimeout(() => {
                setUpdateFail(false);
            }, 2000);
            return;
        }

        patientService.RequestPatientTransfer(admission!.id, department, doctor).then((success) => {
            if (!success) {
                setUpdateFail(true);
                setTimeout(() => {
                    setUpdateFail(false);
                }, 2000);
                return;
            }

            setUpdateSuccess(true);
            setTimeout(() => {
                setUpdateSuccess(false);
            }, 2000);

            updateAdmission({
                ...admission,
            }, {
                timestamp: new Date(),
                action: "Request transfer",
                message: "Transfer requested to department: " + department + ", doctor: " + doctor,
                details: "",
                performedBy: "",
            });
        }).catch(() => {
            setUpdateFail(true);
        });
    }

    const orderLabTest = (labTest: string, admissionId: string) => {
        patientService.OrderLabTest({labTest: labTest, admissionId: admissionId}).then((success) => {
            if (!success) {
                setUpdateFail(true);
                setTimeout(() => {
                    setUpdateFail(false);
                }, 2000);
                return;
            }

            setUpdateSuccess(true);
            setTimeout(() => {
                setUpdateSuccess(false);
            }, 2000);

            if (admission == undefined) {
                return;
            }

            updateAdmission({
                ...admission,
            }, {
                timestamp: new Date(),
                action: "Order lab test",
                message: "Lab test ordered: " + labTest,
                details: "",
                performedBy: "",
            });
        }).catch(() => {
            setUpdateFail(true);
        });
    }

    const addTeamMember = (teammateId: string) => {
        if (admission == undefined) {
            setUpdateFail(true);
            setTimeout(() => {
                setUpdateFail(false);
            }, 2000);
            return;
        }

        patientService.AddTeamMember(admission.id, teammateId).then((success) => {
            if (!success) {
                setUpdateFail(true);
                setTimeout(() => {
                    setUpdateFail(false);
                }, 2000);
                return;
            }

            setUpdateSuccess(true);
            setTimeout(() => {
                setUpdateSuccess(false);
            }, 2000);

            updateAdmission({
                ...admission,
            }, {
                timestamp: new Date(),
                action: "Add team member",
                message: "Team member added: " + teammateId,
                details: "",
                performedBy: "",
            });
        }).catch(() => {
            setUpdateFail(true);
        });
    }

    const addTeamMemberPermissions = (teammateId: string, section: string, permission: string) => {
        if (admission == undefined) {
            setUpdateFail(true);
            setTimeout(() => {
                setUpdateFail(false);
            }, 2000);
            return;
        }

        patientService.AddTeamMemberPermissions(admission.id, teammateId, section, permission).then((success) => {
            if (!success) {
                setUpdateFail(true);
                setTimeout(() => {
                    setUpdateFail(false);
                }, 2000);
                return;
            }

            setUpdateSuccess(true);
            setTimeout(() => {
                setUpdateSuccess(false);
            }, 2000);

            updateAdmission({
                ...admission,
            }, {
                timestamp: new Date(),
                action: "Add team member permissions",
                message: "Team member " + teammateId + " permission added: " + section + " - " + permission,
                details: "",
                performedBy: "",
            });
        }).catch(() => {
            setUpdateFail(true);
        });
    }

    const removeTeamMember = (teammateId: string) => {
        if (admission == undefined) {
            setUpdateFail(true);
            setTimeout(() => {
                setUpdateFail(false);
            }, 2000);
            return;
        }

        patientService.RemoveTeamMember(admission.id, teammateId).then((success) => {
            if (!success) {
                setUpdateFail(true);
                setTimeout(() => {
                    setUpdateFail(false);
                }, 2000);
                return;
            }

            setUpdateSuccess(true);
            setTimeout(() => {
                setUpdateSuccess(false);
            }, 2000);

            updateAdmission({
                ...admission,
            }, {
                timestamp: new Date(),
                action: "Remove team member",
                message: "Team member removed: " + teammateId,
                details: "",
                performedBy: "",
            });
        }).catch(() => {
            setUpdateFail(true);
        });
    }

    const removeTeamMemberPermissions = (teammateId: string, section: string) => {
        if (admission == undefined) {
            setUpdateFail(true);
            setTimeout(() => {
                setUpdateFail(false);
            }, 2000);
            return;
        }

        patientService.RemoveTeamMemberPermissions(admission.id, teammateId, section).then((success) => {
            if (!success) {
                setUpdateFail(true);
                setTimeout(() => {
                    setUpdateFail(false);
                }, 2000);
                return;
            }

            setUpdateSuccess(true);
            setTimeout(() => {
                setUpdateSuccess(false);
            }, 2000);

            updateAdmission({
                ...admission,
            }, {
                timestamp: new Date(),
                action: "Remove team member permissions",
                message: "Team member " + teammateId + " permission removed: " + section,
                details: "",
                performedBy: "",
            });
        }).catch(() => {
            setUpdateFail(true);
        })
    };

    useEffect(() => {
        if (id == undefined) {
            setFetchError(true);
            return;
        }

        patientService.GetAdmissionDetails(id).then((admission) => {
            if (admission !== undefined) {
                if (user !== undefined) {
                    const assignments = new Map<string, User>(Object.entries(admission.careTeam.assignments));
                    const teammate = assignments.get(user.id);
                    if (teammate == undefined) {
                        setFetchError(true);
                        return;
                    }

                    if (teammate.permissions == null) {
                        setFetchError(true);
                        return;
                    }

                    const perm =  new Map<string, string>(Object.entries(teammate.permissions));
                    setPermissions(perm);

                    const permissionKeys = Array.from(perm.keys());
                    filterSections(infoSections, permissionKeys);
                }
                setAdmission(admission);

                return;
            }

            setFetchError(true);
        }).catch((e) => {
            console.error(e);
            setFetchError(true);
        });
    }, [id, user]);

    if (admission == undefined && !fetchError) {
        return (
            <Box
                sx={{
                    position: 'fixed',
                    top: 0,
                    left: 0,
                    width: '100%',
                    height: '100%',
                    display: 'flex',
                    justifyContent: 'center',
                    alignItems: 'center',
                    backgroundColor: 'rgba(249,251,252,0.73)', // Optional: to add a semi-transparent background
                    zIndex: 9999, // Optional: to make sure it's above other elements
                }}
            >
                <CircularProgress color="inherit" />
            </Box>
        )
    }

    return (
        <>
            {admission != undefined ? (
                <Box sx={{ display: 'flex' }}>
                    <AppBar
                        position="fixed"
                        color="default"
                        sx={{ width: `calc(100% - ${drawerWidth}px)`, ml: `${drawerWidth}px`, zIndex: (theme) => theme.zIndex.drawer + 1, mt: 8.5 }}
                    >
                        <Toolbar>
                            {
                                admission.status === "pending" ?
                                    <Chip label="Pending Admission" color="warning" />
                                    :
                                    <Chip label="Admitted" color="success" />
                            }
                            <Typography
                                variant="h6"
                                noWrap
                                component="div"
                                sx={{
                                    fontWeight: 400,
                                    letterSpacing: '.1rem',
                                    ml: 3
                                }}
                            >
                                {`${admission.patient.firstname} ${admission.patient.lastname} 
                                (NID: ${admission.patient.nationalIdentificationNumber},
                                Phone: ${admission.patient.phoneNumber}, 
                                Email: ${admission.patient.email || 'N/A'})`}
                            </Typography>
                        </Toolbar>
                    </AppBar>
                    <CssBaseline />
                    <Drawer
                        variant="permanent"
                        sx={{
                            width: drawerWidth,
                            flexShrink: 0,
                            [`& .MuiDrawer-paper`]: { width: drawerWidth, boxSizing: 'border-box' },
                        }}
                    >
                        <Box sx={{ overflow: 'auto', mt: 3, pl: 1 }}>
                            <Toolbar />
                            <List>
                                {Array.from(filteredInfoSection.entries()).map(([key, value]) => (
                                    <ListItem key={value.name} disablePadding onClick={() => chooseSection(key, value.hasSubsection)}>
                                        <ListItemButton>
                                            <ListItemIcon>
                                                {value.icon}
                                            </ListItemIcon>
                                            <ListItemText primary={value.name} />
                                        </ListItemButton>
                                    </ListItem>
                                ))}
                            </List>
                        </Box>
                    </Drawer>
                    <Box component="main" sx={{ flexGrow: 1, p: 3, mt: 16 }}>
                        {updateFail && <Alert variant="filled" severity="error" sx={{ mb: 2 }}>Error updating admission details!</Alert>}
                        {updateSuccess && <Alert variant="filled" severity="success" sx={{ mb: 2 }}>Success!</Alert>}
                        {getPage()}
                    </Box>
                </Box>
            ) :(
                <Box/>
            )}

            <Modal
                sx={{
                    backgroundColor: 'rgba(249,251,252,0.73)', // Optional: to add a semi-transparent background
                    zIndex: 9999, // Optional: to make sure it's above other elements
                }}
                open={fetchError}
                onClose={() => navigate("/patients/admissions")}
                aria-labelledby="parent-modal-title"
                aria-describedby="parent-modal-description"
            >
                <Box sx={{ ...style, width: 400 }}>
                    <h2 id="parent-modal-title">Error!</h2>
                    <p id="parent-modal-description">
                        Fetching admission details failed. Please try again later, or check if you have correct access rights.
                    </p>
                    <Button
                        variant="text"
                        onClick={() => navigate("/patients/admissions")}
                    >
                        Go Back to Admissions page
                    </Button>
                </Box>
            </Modal>
        </>
    );
};

export default AdmissionDetailsPage;
