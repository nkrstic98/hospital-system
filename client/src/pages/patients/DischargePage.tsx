import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import {ExitToApp} from "@mui/icons-material";
import Box from "@mui/material/Box";
import {Modal} from "@mui/material";
import {useState} from "react";

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

export interface DischargePageProps {
    permission: string;
    canDischarge: boolean;
    dischargePatient(): void;
}

const DischargePage = ({permission, canDischarge, dischargePatient}: DischargePageProps) => {
    const [open, setOpen] = useState(false);

    return (
        <>
            <div>
                <Typography variant="h4" gutterBottom component="div" sx={{mb: 3}}>
                    Discharge Patient
                </Typography>
                {canDischarge && permission == "WRITE" && <Button startIcon={<ExitToApp />} variant="contained" color="secondary" onClick={() => setOpen(true)}>Discharge Patient</Button>}
                {!canDischarge && <Typography variant="h5" gutterBottom component="div" sx={{mb: 3}}>Finish diagnosis first, and then discharge patient!</Typography>}
            </div>

            <Modal
                sx={{
                    backgroundColor: 'rgba(249,251,252,0.73)', // Optional: to add a semi-transparent background
                    zIndex: 9999, // Optional: to make sure it's above other elements
                }}
                open={open}
                onClose={() => setOpen(false)}
                aria-labelledby="parent-modal-title"
                aria-describedby="parent-modal-description"
            >
                <Box sx={{ ...style, width: 400 }}>
                    <h2 id="parent-modal-title">Please Confirm!</h2>
                    <p id="parent-modal-description">
                        Are you sure you want to discharge this patient? This action cannot be undone.
                    </p>
                    <Button
                        variant="text"
                        onClick={() => dischargePatient()}
                    >
                        Confirm Patient Discharge
                    </Button>
                </Box>
            </Modal>
        </>
    );
}

export default DischargePage;
