import * as React from 'react';
import Box from '@mui/material/Box';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import CssBaseline from "@mui/material/CssBaseline";
import {UserService} from "../../services/user/User.ts";
import {User} from "../../types/User.ts";
import Button from "@mui/material/Button";
import AddIcon from '@mui/icons-material/Add';
import {useNavigate} from "react-router-dom";

const UserRow = (props: { user: User })=> {
    return (
        <React.Fragment>
            <TableRow sx={{ '& > *': { borderBottom: 'unset' } }}>
                <TableCell align="center">{props.user.username}</TableCell>
                <TableCell component="th" scope="row" align="center">
                    {props.user.firstname + " " + props.user.lastname}
                </TableCell>
                <TableCell align="center">{props.user.team}</TableCell>
                <TableCell align="center">{props.user.role}</TableCell>
                <TableCell align="center">{props.user.email}</TableCell>
            </TableRow>
        </React.Fragment>
    );
}

const EmployeeManagement = ()=> {
    const navigate = useNavigate();
    const userService: UserService = new UserService();
    const [users, setUsers] = React.useState<User[]>([]);

    React.useEffect(() => {
        userService.GetUsers().then((response) => {
            if (response === undefined) {
                return;
            }
            setUsers(response);
        });
    }, []);

    return (
        <Box sx={{ display: 'flex', mt: 4 }}>
            <CssBaseline />
            <Box component="main" sx={{ flexGrow: 1, p: 3 }}>
                <Button
                    variant="contained"
                    sx={{ mt: 3, mb: 5 }}
                    startIcon={<AddIcon />}
                    onClick={() => navigate("/admin/employees/register")}
                >
                    Add New Employee
                </Button>
                <TableContainer component={Paper}>
                    <Table aria-label="collapsible table">
                        <TableHead>
                            <TableRow>
                                <TableCell align="center">Username</TableCell>
                                <TableCell align="center">Full Name</TableCell>
                                <TableCell align="center">Department</TableCell>
                                <TableCell align="center">Role</TableCell>
                                <TableCell align="center">Contact Email</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {users.map((user) => (
                                <UserRow key={user.username} user={user} />
                            ))}
                        </TableBody>
                    </Table>
                </TableContainer>
            </Box>
        </Box>
    );
}

export default EmployeeManagement;