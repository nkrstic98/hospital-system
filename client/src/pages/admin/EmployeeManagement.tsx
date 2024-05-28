import * as React from 'react';
import Box from '@mui/material/Box';
import Collapse from '@mui/material/Collapse';
import IconButton from '@mui/material/IconButton';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Typography from '@mui/material/Typography';
import Paper from '@mui/material/Paper';
import KeyboardArrowDownIcon from '@mui/icons-material/KeyboardArrowDown';
import KeyboardArrowUpIcon from '@mui/icons-material/KeyboardArrowUp';
import CssBaseline from "@mui/material/CssBaseline";
import {UserService} from "../../services/user/User.ts";
import {User} from "../../types/User.ts";
import Button from "@mui/material/Button";
import AddIcon from '@mui/icons-material/Add';
import {Alert, Grid} from "@mui/material";
import {useNavigate} from "react-router-dom";

const UserRow = (props: { user: User })=> {
    const [open, setOpen] = React.useState(false);

    return (
        <React.Fragment>
            <TableRow sx={{ '& > *': { borderBottom: 'unset' } }}>
                <TableCell>
                    <IconButton
                        aria-label="expand row"
                        size="small"
                        onClick={() => setOpen(!open)}
                    >
                        {open ? <KeyboardArrowUpIcon /> : <KeyboardArrowDownIcon />}
                    </IconButton>
                </TableCell>
                <TableCell align="center">{props.user.username}</TableCell>
                <TableCell component="th" scope="row" align="center">
                    {props.user.firstname + " " + props.user.lastname}
                </TableCell>
                <TableCell align="center">{props.user.team}</TableCell>
                <TableCell align="center">{props.user.role}</TableCell>
                <TableCell align="center">{props.user.phoneNumber}</TableCell>
                <TableCell align="center">{props.user.email}</TableCell>
            </TableRow>
            <TableRow>
                <TableCell style={{ paddingBottom: 0, paddingTop: 0 }} colSpan={7}>
                    <Collapse in={open} timeout="auto" unmountOnExit>
                        <Box sx={{ margin: 1, p: 2 }}>
                            <Typography variant="h6" gutterBottom component="div">
                                Employee Details
                            </Typography>
                            {props.user.verified ? (
                                <Grid container sx={{ pl: 3 }}>
                                    <Grid item xs={3}>
                                        <p><b>National ID: </b></p>
                                    </Grid>
                                    <Grid item xs={8}>
                                        <p>{props.user.nationalIdentificationNumber}</p>
                                    </Grid>
                                    <Grid item xs={3}>
                                        <p><b>Address: </b></p>
                                    </Grid>
                                    <Grid item xs={8}>
                                        <p>{props.user.mailingAddress}, {props.user.city}, {props.user.state}, {props.user.zip}</p>
                                    </Grid>
                                </Grid>
                            ) : (
                                <Alert severity="warning">Details unavailable! Waiting for employee to finish registration.</Alert>
                            )}
                        </Box>
                    </Collapse>
                </TableCell>
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
                                <TableCell />
                                <TableCell align="center">Username</TableCell>
                                <TableCell align="center">Full Name</TableCell>
                                <TableCell align="center">Department</TableCell>
                                <TableCell align="center">Role</TableCell>
                                <TableCell align="center">Contact Phone</TableCell>
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