import * as React from 'react';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import IconButton from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';
import Menu from '@mui/material/Menu';
import Avatar from '@mui/material/Avatar';
import Button from '@mui/material/Button';
import Tooltip from '@mui/material/Tooltip';
import MenuItem from '@mui/material/MenuItem';
import CssBaseline from "@mui/material/CssBaseline";
import {useAuth} from "../router/AuthProvider.tsx";
import {Link, useNavigate} from "react-router-dom";
import {GetUserPermission} from "../utils/utils.ts";

const adminPages: Map<string, string[]> = new Map([
    ["INTAKE", ['Patient Intake', '/patient-intake']],
    ["EMPLOYEES", ['Employee Management', '/employees']],
    ["PATIENTS", ['Assigned Patients', '/patients']],
]);

function ResponsiveAppBar() {
    const { isAuthenticated, user, onLogout } = useAuth();
    const navigate = useNavigate();
    const [anchorElUser, setAnchorElUser] = React.useState<null | HTMLElement>(null);

    const handleOpenUserMenu = (event: React.MouseEvent<HTMLElement>) => {
        setAnchorElUser(event.currentTarget);
    };

    const handleCloseUserMenu = () => {
        setAnchorElUser(null);
    };

    if (!isAuthenticated) {
        return;
    }

    return (
        <Box sx={{ display: 'flex' }}>
            <CssBaseline />
            <AppBar position="fixed" sx={{ zIndex: (theme) => theme.zIndex.drawer + 1 }}>
                    <Toolbar>
                        <Avatar
                            src="/logo21.png"
                            alt="Logo"
                            sx={{ display: { xs: 'none', md: 'flex' }, mr: 2 }}
                            style={{width: '50px', height: '50px'}}
                        />
                        <Typography
                            variant="h6"
                            noWrap
                            component="div"
                            sx={{
                                mr: 3,
                                display: { xs: 'none', md: 'flex' },
                                fontWeight: 1000,
                                letterSpacing: '.3rem',
                                color: 'inherit',
                                textDecoration: 'none',
                            }}
                        >
                            <Link to="/home" style={{ color: 'inherit', textDecoration: 'none' }}>
                                Genesis Medical
                            </Link>
                        </Typography>

                        <Avatar
                            src="/logo21.png"
                            alt="Logo"
                            sx={{ display: { xs: 'flex', md: 'none' }, mr: 1 }}
                            style={{width: '50px', height: '50px'}}
                        />
                        <Typography
                            variant="h5"
                            noWrap
                            component="div"
                            sx={{
                                mr: 2,
                                display: { xs: 'flex', md: 'none' },
                                flexGrow: 1,
                                fontWeight: 700,
                                letterSpacing: '.3rem',
                                color: 'inherit',
                                textDecoration: 'none',
                            }}
                        >
                            <Link to="/home" style={{ color: 'inherit', textDecoration: 'none' }}>
                                Genesis Medical
                            </Link>
                        </Typography>

                        <Box sx={{flexGrow: 1, display: {xs: 'none', md: 'flex'}}}>
                            {Array.from(adminPages.entries()).map(([key, value]) => (
                                GetUserPermission(user, key) && <Button
                                    variant="outlined"
                                    key={key}
                                    sx={{my: 2, color: 'white', display: 'block'}}
                                    onClick={() => navigate(value[1])}
                                >
                                    {value[0]}
                                </Button>
                            ))}
                        </Box>

                        <Box sx={{flexGrow: 0}}>
                            <Tooltip title="Open settings">
                                <IconButton onClick={handleOpenUserMenu} sx={{p: 0}}>
                                    <Avatar alt={user?.firstname + " " + user?.lastname} />
                                </IconButton>
                            </Tooltip>
                            <Menu
                                sx={{mt: '45px'}}
                                id="menu-appbar"
                                anchorEl={anchorElUser}
                                anchorOrigin={{
                                    vertical: 'top',
                                    horizontal: 'right',
                                }}
                                keepMounted
                                transformOrigin={{
                                    vertical: 'top',
                                    horizontal: 'right',
                                }}
                                open={Boolean(anchorElUser)}
                                onClose={handleCloseUserMenu}
                            >
                                {/*<MenuItem key="Profile" onClick={() => {*/}
                                {/*    handleCloseUserMenu();*/}
                                {/*    navigate("/profile")*/}
                                {/*}}>*/}
                                {/*    <Typography textAlign="center">Profile</Typography>*/}
                                {/*</MenuItem>*/}
                                {/*<MenuItem key="Account" onClick={() => {*/}
                                {/*    handleCloseUserMenu();*/}
                                {/*    navigate("/account")*/}
                                {/*}}>*/}
                                {/*    <Typography textAlign="center">Account</Typography>*/}
                                {/*</MenuItem>*/}
                                <MenuItem key="Logout" onClick={() => {
                                    handleCloseUserMenu();
                                    onLogout();
                                }}>
                                    <Typography textAlign="center">Log Out</Typography>
                                </MenuItem>
                            </Menu>
                        </Box>
                    </Toolbar>
            </AppBar>
        </Box>
    );
}

export default ResponsiveAppBar;