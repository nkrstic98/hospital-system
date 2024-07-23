import {CareTeam} from "../../types/Admission.ts";
import Box from "@mui/material/Box";
import {User} from "../../types/User.ts";
import {
    Accordion,
    AccordionDetails,
    AccordionSummary,
    Alert,
    Divider, FormControl,
    InputLabel, OutlinedInput, Select,
    TableContainer
} from "@mui/material";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import Typography from "@mui/material/Typography";
import Paper from "@mui/material/Paper";
import Table from "@mui/material/Table";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import TableCell from "@mui/material/TableCell";
import TableBody from "@mui/material/TableBody";
import {useEffect, useState} from "react";
import {UserService} from "../../services/user/User.ts";
import MenuItem from "@mui/material/MenuItem";
import Button from "@mui/material/Button";
import {Remove} from "@mui/icons-material";

export interface CareTeamPageProps {
    careTeam: CareTeam;
    permission: string;
    addTeamMember: (teammateId: string) => void;
    addTeamMemberPermissions: (teammateId: string, section: string, permission: string) => void;
    removeTeamMember: (teammateId: string) => void;
    removeTeamMemberPermissions: (teammateId: string, section: string) => void;
}

const roleToDisplayNameMap = new Map<string, string>([
    ["ATTENDING", "Attending Physician"],
    ["RESIDENT", "Resident Doctor"],
    ["NURSE", "Nurse"]
]);

const CareTeamPage = ({
    careTeam,
    permission,
    addTeamMember,
    addTeamMemberPermissions,
    removeTeamMember,
    removeTeamMemberPermissions,
}: CareTeamPageProps) => {
    const userService = new UserService();

    const team: Map<string, User> = new Map<string, User>(Object.entries(careTeam.assignments));
    const teamLead = team.get(careTeam.teamLead);

    const [attendings, setAttendings] = useState<User[]>([]);
    const [residents, setResidents] = useState<User[]>([]);
    const [nurses, setNurses] = useState<User[]>([]);

    const [chosenRole, setChosenRole] = useState<string>("");
    const [chosenUser, setChosenUser] = useState<string>("");

    const [chosenTeamMember, setChosenTeamMember] = useState<string>("");
    const [chosenSection, setChosenSection] = useState<string>("");
    const [chosenPermission, setChosenPermission] = useState<string>("");

    const chooseRole = (role: string) => {
        setChosenRole(role);
        setChosenUser("");
    }

    const addMember = () => {
        addTeamMember(chosenUser);
        setChosenRole("");
    }

    const addPermissions = () => {
        addTeamMemberPermissions(chosenTeamMember, chosenSection, chosenPermission);

        setChosenTeamMember("");
        setChosenSection("");
        setChosenPermission("");
    }

    const removeMember = (teammateId: string) => {
        removeTeamMember(teammateId);
        getDepartmentUsers();
    }

    const removePermissions = (teammateId: string, section: string) => {
        removeTeamMemberPermissions(teammateId, section);
        getDepartmentUsers();
    }

    const getDepartmentUsers = () => {
        userService.GetDepartments({
            team: careTeam.team,
            role: undefined,
        }).then((departments) => {
            if (departments !== undefined) {
                const a: User[] = [];
                const r: User[] = [];
                const n: User[] = [];

                departments.get(careTeam.team)?.users.forEach((user) => {
                    if (user.role == "ATTENDING"  && team.get(user.id) == undefined) {
                        a.push(user);
                    } else if (user.role == "RESIDENT"  && team.get(user.id) == undefined) {
                        r.push(user);
                    } else if (user.role == "NURSE" && team.get(user.id) == undefined) {
                        n.push(user);
                    }

                    setAttendings(a);
                    setResidents(r);
                    setNurses(n);
                });
            }
        });
    }

    useEffect(() => {
        getDepartmentUsers();
    }, []);

    return (
        <Box>
            <Typography variant="h5" gutterBottom component="div" sx={{mb: 4}}>
                <b>Care Department:</b> {careTeam.department}
            </Typography>
            <Typography variant="h5" gutterBottom component="div" sx={{m: 4}}>
                <b>Team Lead:</b> {teamLead?.firstname} {teamLead?.lastname}, {roleToDisplayNameMap.get(teamLead?.role ?? "")}
            </Typography>
            <Divider />
            {permission === "WRITE" && (
                <>
                    <Typography variant="h5" gutterBottom component="div" sx={{m: 4}}>
                        Add Team Member
                    </Typography>
                    <TableContainer component={Paper}>
                        <Table>
                            <TableBody>
                                <TableRow>
                                    <TableCell>
                                        <FormControl fullWidth required size="small" sx={{minWidth: "200px"}}>
                                            <InputLabel id="dept-label">Choose team member role</InputLabel>
                                            <Select
                                                labelId="role"
                                                id="role"
                                                value={chosenRole}
                                                onChange={(e) => chooseRole(e.target.value)}
                                                input={<OutlinedInput label="Choose team member role" />}
                                                required
                                                fullWidth
                                            >
                                                <MenuItem key={"ATTENDING"} value={"ATTENDING"}>Attending Physician</MenuItem>
                                                <MenuItem key={"RESIDENT"} value={"RESIDENT"}>Resident Doctor</MenuItem>
                                                <MenuItem key={"NURSE"} value={"NURSE"}>Nurse</MenuItem>
                                            </Select>
                                        </FormControl>
                                    </TableCell>
                                    <TableCell>
                                        <FormControl fullWidth required size="small" sx={{minWidth: "200px"}}>
                                            <InputLabel id="dept-label">Choose team member</InputLabel>
                                            <Select
                                                labelId="role"
                                                id="role"
                                                value={chosenUser}
                                                onChange={(e) => setChosenUser(e.target.value)}
                                                input={<OutlinedInput label="Choose Team Member" />}
                                                required
                                                fullWidth
                                            >
                                                {chosenRole === "ATTENDING" && attendings.map((a) => (
                                                    <MenuItem key={a.id} value={a.id}>{a.firstname} {a.lastname}</MenuItem>
                                                ))}
                                                {chosenRole === "RESIDENT" && residents.map((r) => (
                                                    <MenuItem key={r.id} value={r.id}>{r.firstname} {r.lastname}</MenuItem>
                                                ))}
                                                {chosenRole === "NURSE" && nurses.map((n) => (
                                                    <MenuItem key={n.id} value={n.id}>{n.firstname} {n.lastname}</MenuItem>
                                                ))}
                                                {chosenRole === "" && <MenuItem value={""}>Select a role first</MenuItem>}
                                            </Select>
                                        </FormControl>
                                    </TableCell>
                                    <TableCell>
                                        <Button variant="contained" color="success" size="small" disabled={chosenRole == "" || chosenUser == ""} onClick={() => addMember()}>ADD</Button>
                                    </TableCell>
                                </TableRow>
                            </TableBody>
                        </Table>
                    </TableContainer>
                    <Divider />
                    <Typography variant="h5" gutterBottom component="div" sx={{m: 4}}>
                        Give permissions to the team member
                    </Typography>
                    <TableContainer component={Paper}>
                        <Table>
                            <TableBody>
                                <TableRow>
                                    <TableCell>
                                        <FormControl fullWidth required size="small" sx={{minWidth: "200px"}}>
                                            <InputLabel id="dept-label">Choose team member</InputLabel>
                                            <Select
                                                labelId="role"
                                                id="role"
                                                value={chosenTeamMember}
                                                onChange={(e) => setChosenTeamMember(e.target.value)}
                                                input={<OutlinedInput label="Choose team member" />}
                                                required
                                                fullWidth
                                            >
                                                {Array.from(team.entries()).map(([key, user]) => (
                                                    key !== teamLead?.id && <MenuItem key={key} value={key}>{user.firstname} {user.lastname}</MenuItem>
                                                ))}
                                            </Select>
                                        </FormControl>
                                    </TableCell>
                                    <TableCell>
                                        <FormControl fullWidth required size="small" sx={{minWidth: "200px"}}>
                                            <InputLabel id="dept-label">Choose section</InputLabel>
                                            <Select
                                                labelId="role"
                                                id="role"
                                                value={chosenSection}
                                                onChange={(e) => setChosenSection(e.target.value)}
                                                input={<OutlinedInput label="Choose section" />}
                                                required
                                                fullWidth
                                            >
                                                <MenuItem key="PATIENTS:DIAGNOSIS" value="PATIENTS:DIAGNOSIS">PATIENTS:DIAGNOSIS</MenuItem>
                                                <MenuItem key="PATIENTS:CONSULTS" value="PATIENTS:CONSULTS">PATIENTS:CONSULTS</MenuItem>
                                                <MenuItem key="PATIENTS:TRANSFER" value="PATIENTS:TRANSFER">PATIENTS:TRANSFER</MenuItem>
                                                <MenuItem key="PATIENTS:DISCHARGE" value="PATIENTS:DISCHARGE">PATIENTS:DIAGNOSIS</MenuItem>
                                                <MenuItem key="PATIENTS:MEDICATIONS:PRESCRIBE" value="PATIENTS:MEDICATIONS:GIVE">PATIENTS:MEDICATIONS:GIVE</MenuItem>
                                                <MenuItem key="PATIENTS:MEDICATIONS:PRESCRIBE" value="PATIENTS:MEDICATIONS:PRESCRIBE">PATIENTS:MEDICATIONS:PRESCRIBE</MenuItem>
                                                <MenuItem key="PATIENTS:LABS:ORDER" value="PATIENTS:LABS:ORDER">PATIENTS:LABS:ORDER</MenuItem>
                                                <MenuItem key="PATIENTS:LABS:ORDER" value="PATIENTS:LABS:RESULT">PATIENTS:LABS:RESULT</MenuItem>
                                                <MenuItem key="PATIENTS:IMAGING:ORDER" value="PATIENTS:IMAGING:ORDER">PATIENTS:DIAGNOSIS</MenuItem>
                                                <MenuItem key="PATIENTS:LOGS" value="PATIENTS:LOGS">PATIENTS:LOGS</MenuItem>
                                                <MenuItem key="PATIENTS:TEAM" value="PATIENTS:TEAM">PATIENTS:TEAM</MenuItem>
                                            </Select>
                                        </FormControl>
                                    </TableCell>
                                    <TableCell>
                                        <FormControl fullWidth required size="small" sx={{minWidth: "200px"}}>
                                            <InputLabel id="dept-label">Choose permission</InputLabel>
                                            <Select
                                                labelId="role"
                                                id="role"
                                                value={chosenPermission}
                                                onChange={(e) => setChosenPermission(e.target.value)}
                                                input={<OutlinedInput label="Choose permission" />}
                                                required
                                                fullWidth
                                            >
                                                <MenuItem key="READ" value="READ">READ</MenuItem>
                                                <MenuItem key="WRITE" value="WRITE">WRITE</MenuItem>
                                            </Select>
                                        </FormControl>
                                    </TableCell>
                                    <TableCell>
                                        <Button variant="contained" color="success" size="small" disabled={chosenTeamMember == "" || chosenSection == "" || chosenPermission == ""} onClick={() => addPermissions()}>ADD</Button>
                                    </TableCell>
                                </TableRow>
                            </TableBody>
                        </Table>
                    </TableContainer>
                    <Divider />
                </>
            )}
            <Typography variant="h5" gutterBottom component="div" sx={{m: 4}}>
                Care Team
            </Typography>
            {Array.from(new Map<string, User>(Object.entries(careTeam.assignments)).entries()).map(([key, user]) => (
                <Accordion key={key}>
                    <AccordionSummary
                        expandIcon={<ExpandMoreIcon/>}
                        aria-controls="panel1-content"
                        id="panel1-header"
                    >
                        <b>&nbsp; {user.firstname} {user.lastname} - {roleToDisplayNameMap.get(user.role)}</b>
                    </AccordionSummary>
                    <AccordionDetails>
                        {permission == "WRITE" && key !== teamLead?.id && <Button variant="contained" color="error" size="small" onClick={() => {
                            removeMember(key)
                        }}>REMOVE FROM THE TEAM</Button>}
                        <Typography variant="h6" gutterBottom component="div" sx={{m: 4}}>
                            Permissions
                        </Typography>
                        {user.permissions !== null ? (
                            <TableContainer component={Paper}>
                                <Table>
                                    <TableHead>
                                        <TableRow>
                                            <TableCell><b>Section</b></TableCell>
                                            <TableCell><b>Permission</b></TableCell>
                                            <TableCell></TableCell>
                                        </TableRow>
                                    </TableHead>
                                    <TableBody>
                                        {user.permissions !== undefined && Array.from(new Map<string, string>(Object.entries(user.permissions)).entries()).map(([section, p]) => (
                                            <TableRow key={`${key}-${section}`}>
                                                <TableCell>{section}</TableCell>
                                                <TableCell>{p}</TableCell>
                                                <TableCell align="right">
                                                    {permission === "WRITE" && key !== teamLead?.id && <Button variant="outlined" color="error" size="small" onClick={() => {
                                                        removePermissions(key, section)
                                                    }}><Remove /></Button>}
                                                </TableCell>
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
        </Box>
    );
}

export default CareTeamPage;
