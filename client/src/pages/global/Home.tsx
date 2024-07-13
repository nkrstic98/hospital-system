import {useAuth} from "../../router/AuthProvider.tsx";
import CssBaseline from "@mui/material/CssBaseline";
import Box from "@mui/material/Box";
import Card from "@mui/material/Card";
import CardMedia from "@mui/material/CardMedia";
import Typography from "@mui/material/Typography";
import Container from "@mui/material/Container";

const Home = () => {
    const { user } = useAuth();

    return (
        <Container component="main" maxWidth="xl">
            <CssBaseline />
            <Box sx={{mt: 12}}></Box>
            <Box
                sx={{
                    marginTop: 5,
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center',
                }}
                maxWidth="xl"
            >
                <Card>
                    <CardMedia
                        component="img"
                        height="250"
                        image="logo-new.png"
                        alt="Genesis Medical Center"
                    />
                </Card>
            </Box>
            <Typography variant="h3" gutterBottom component="div" sx={{mt: 5}}>
                Welcome {user?.firstname} {user?.lastname}!
            </Typography>
        </Container>
    );
}

export default Home;
