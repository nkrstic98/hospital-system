import Typography from "@mui/material/Typography";
import {FormControl, FormHelperText, InputLabel, OutlinedInput, Select} from "@mui/material";
import MenuItem from "@mui/material/MenuItem";
import {useState} from "react";
import {User} from "../../types/User.ts";
import {Department} from "../../services/user/types.ts";

export interface DepartmentAdPhysiciansProps {
    user?: User,
    departmentPhysicians: Map<string, Department> | undefined;
    department: string;
    physician: string;
    setDepartment: (department: string) => void;
    setPhysicianName: (physicianId: string) => void;
    setPhysicianId: (physician: string) => void;
    departmentError: boolean;
    physicianError: boolean;
}

const DepartmentsAndPhysicians = ({
    user,
    departmentPhysicians,
    department,
    physician,
    setDepartment,
    setPhysicianId,
    setPhysicianName,
    departmentError,
    physicianError
}: DepartmentAdPhysiciansProps) => {
    const [doctor, setDoctor] = useState<string>(JSON.stringify({
        id: user?.id ?? "",
        firstname: user?.firstname ?? "",
        lastname: user?.lastname ?? "",
        nationalIdentificationNumber: user?.nationalIdentificationNumber ?? "",
        username: user?.username ?? "",
        email: user?.email ?? "",
        role: "",
        team: null,
        permissions: null,
    }));

    return (
        <>
            <Typography variant="h6" gutterBottom component="div" sx={{mb: 3}}>
                Choose department and physician
            </Typography>
            <FormControl fullWidth required>
                <InputLabel id="dept-label">Department</InputLabel>
                <Select
                    labelId="dept-label"
                    id="dept"
                    value={department}
                    onChange={(e) => setDepartment(e.target.value)}
                    input={<OutlinedInput label="Department" />}
                    required
                    fullWidth
                    error={departmentError && department === ""}
                >
                    {
                        departmentPhysicians !== undefined &&
                        Array.from((departmentPhysicians as Map<string, Department>).entries()).map(([key, value]) => (
                            <MenuItem key={key} value={key}>{value.displayName}</MenuItem>
                        ))
                    }
                </Select>
                {departmentError && department === "" && <FormHelperText error>This field is required</FormHelperText>}
            </FormControl>
            <FormControl fullWidth required sx={{ mt: 3 }}>
                <InputLabel id="phys-label">Physician</InputLabel>
                <Select
                    labelId="phys-label"
                    id="phys"
                    value={doctor}
                    onChange={(e) => {
                        setDoctor(e.target.value as string);
                        const physician: User = JSON.parse(e.target.value);
                        setPhysicianId(physician.id);
                        setPhysicianName(physician.firstname + " " + physician.lastname);
                    }}
                    input={<OutlinedInput label="Physician" />}
                    required
                    fullWidth
                    error={physicianError && physician === ""}
                >
                    {
                        departmentPhysicians !== undefined &&
                        departmentPhysicians.get(department)?.users.map((p) => (
                            <MenuItem key={p.nationalIdentificationNumber} value={JSON.stringify(p)}>{p.firstname} {p.lastname}</MenuItem>
                        ))
                    }
                    {
                        department === "" && <MenuItem key={""} value={""}>Choose department first...</MenuItem>
                    }
                </Select>
                {physicianError && physician === "" && <FormHelperText error>This field is required</FormHelperText>}
            </FormControl>
        </>
    )
}

export default DepartmentsAndPhysicians;