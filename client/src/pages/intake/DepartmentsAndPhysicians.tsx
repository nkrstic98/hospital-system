import {Department, Employee} from "../../services/department/types.ts";
import Typography from "@mui/material/Typography";
import {FormControl, FormHelperText, InputLabel, OutlinedInput, Select} from "@mui/material";
import MenuItem from "@mui/material/MenuItem";
import {useState} from "react";

export interface DepartmentAdPhysiciansProps {
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
    departmentPhysicians,
    department,
    physician,
    setDepartment,
    setPhysicianId,
    setPhysicianName,
    departmentError,
    physicianError
}: DepartmentAdPhysiciansProps) => {
    const [doctor, setDoctor] = useState<string>("");

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
                        const physician: Employee = JSON.parse(e.target.value);
                        setPhysicianId(physician.id);
                        setPhysicianName(physician.fullName);
                    }}
                    input={<OutlinedInput label="Physician" />}
                    required
                    fullWidth
                    error={physicianError && physician === ""}
                >
                    {
                        departmentPhysicians !== undefined &&
                        departmentPhysicians.get(department)?.physicians.map((p) => (
                            <MenuItem key={p.id} value={JSON.stringify(p)}>{p.fullName}</MenuItem>
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