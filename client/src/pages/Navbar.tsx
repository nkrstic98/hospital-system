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
import {useNavigate} from "react-router-dom";

const adminPages: Map<string, string> = new Map([
    ['Patient Intake', '/admin'],
    ['Employee Management', '/admin/employees'],
    // ['Department Management', '/admin/departments'],
]);

const attendingPages: Map<string, string> = new Map([
    ['Patient Intake', '/admin'],
    ['Employee Management', '/admin/employees'],
    // ['Department Management', '/admin/departments'],
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
                                mr: 2,
                                display: { xs: 'none', md: 'flex' },
                                fontWeight: 700,
                                letterSpacing: '.3rem',
                                color: 'inherit',
                                textDecoration: 'none',
                            }}
                        >
                            Zmaj Medical Center
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
                            Zmaj Medical
                        </Typography>

                        <Box sx={{flexGrow: 1, display: {xs: 'none', md: 'flex'}}}>
                            {user !== undefined && user.role === "ADMIN" && Array.from(adminPages.entries()).map(([key, value]) => (
                                <Button
                                    variant="outlined"
                                    key={key}
                                    sx={{my: 2, color: 'white', display: 'block'}}
                                    onClick={() => navigate(value)}
                                >
                                    {key}
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
                                    <Typography textAlign="center">Logout</Typography>
                                </MenuItem>
                            </Menu>
                        </Box>
                    </Toolbar>
            </AppBar>
        </Box>
    );
}

export default ResponsiveAppBar;