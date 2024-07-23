import {useEffect, useState} from "react";
import {Department} from "../../services/user/types.ts";
import {UserService} from "../../services/user/User.ts";
import Typography from "@mui/material/Typography";
import {FormControl, FormHelperText, InputLabel, OutlinedInput, Select} from "@mui/material";
import MenuItem from "@mui/material/MenuItem";
import Button from "@mui/material/Button";
import Box from "@mui/material/Box";

export interface PatientTransferPageProps {
    permission: string;
    requestTransfer: (department: string, doctor: string) => void;
}

const PatientTransferPage = ({permission, requestTransfer}: PatientTransferPageProps) => {
    const userService = new UserService();
    const [departmentPhysicians, setDepartmentPhysicians] = useState<Map<string, Department> | undefined>(undefined);
    const [department , setDepartment] = useState<string>("");
    const [doctor, setDoctor] = useState<string>("");
    const [submitAttempted, setSubmitAttempted] = useState<boolean>(false);

    const submit = () => {
        setSubmitAttempted(true);
        if (department === "" || doctor === "") {
            return;
        }
        requestTransfer(department, doctor);

        setDepartment("");
        setDoctor("");
        setSubmitAttempted(false);
    }

    useEffect(() => {
        userService.GetDepartments({
            team: undefined,
            role: "ATTENDING"
        }).then((data) => {
            setDepartmentPhysicians(data);
        });
    }, []);

    if (permission !== "WRITE") {
        return (
            <Typography variant="h4" gutterBottom component="div" sx={{mb: 3}}>
                You do not have permission to view this page
            </Typography>
        );
    }

    return (
        <Box>
            <Typography variant="h6" gutterBottom component="div" sx={{mb: 3}}>
                Choose department and physician to whom you want to transfer the patient
            </Typography>

            <FormControl fullWidth required>
                <InputLabel id="dept-label">Department</InputLabel>
                <Select
                    labelId="dept"
                    id="dept"
                    value={department}
                    onChange={(e) => setDepartment(e.target.value)}
                    input={<OutlinedInput label="Department" />}
                    required
                    fullWidth
                    error={submitAttempted && department === ""}
                >
                    {
                        departmentPhysicians !== undefined &&
                        Array.from((departmentPhysicians as Map<string, Department>).entries()).map(([key, value]) => (
                            <MenuItem key={key} value={key}>{value.displayName}</MenuItem>
                        ))
                    }
                </Select>
                {submitAttempted && department === "" && <FormHelperText error>This field is required</FormHelperText>}
            </FormControl>

            <FormControl fullWidth required sx={{ mt: 3, mb: 4 }}>
                <InputLabel id="phys-label">Physician</InputLabel>
                <Select
                    labelId="phys-label"
                    id="phys"
                    value={doctor}
                    onChange={(e) => {
                        setDoctor(e.target.value);
                    }}
                    input={<OutlinedInput label="Physician" />}
                    required
                    fullWidth
                    error={submitAttempted && doctor === ""}
                >
                    {
                        departmentPhysicians !== undefined &&
                        departmentPhysicians.get(department)?.users.map((p) => (
                            <MenuItem key={p.id} value={p.id}>{p.firstname} {p.lastname}</MenuItem>
                        ))
                    }
                    {
                        department === "" && <MenuItem key={""} value={""}>Choose department first...</MenuItem>
                    }
                </Select>
                {submitAttempted && doctor === "" && <FormHelperText error>This field is required</FormHelperText>}
            </FormControl>

            <Button variant="contained" onClick={() => submit()}>Submit Transfer Request</Button>
        </Box>
    );
}

export default PatientTransferPage;
