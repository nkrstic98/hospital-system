import { Autocomplete } from "@mui/material";
import TextField from "@mui/material/TextField";

const commonDrugs: string[] = [
    "Paracetamol (Acetaminophen)",
    "Ibuprofen",
    "Aspirin (Acetylsalicylic Acid)",
    "Loratadine",
    "Cetirizine",
    "Omeprazole",
    "Simvastatin",
    "Metformin",
    "Salbutamol (Albuterol)",
    "Amoxicillin",
    "Amlodipine",
    "Lisinopril",
    "Atorvastatin",
    "Warfarin",
    "Metoprolol",
    "Clopidogrel",
    "Levothyroxine",
    "Pantoprazole",
    "Diazepam",
    "Tramadol",
    "Citalopram",
    "Fluoxetine",
    "Gabapentin",
    "Sertraline",
    "Hydrochlorothiazide",
    "Montelukast",
    "Prednisone",
    "Alprazolam",
    "Methotrexate",
    "Aspirin-Dipyridamole"
];

export interface PatientMedicationsProps {
    handleMedicationsChange: (value: string[]) => void;
}

const PatientMedications = ({ handleMedicationsChange }: PatientMedicationsProps) => {
    return (
        <Autocomplete
            fullWidth
            multiple
            id="tags-outlined"
            options={commonDrugs}
            getOptionLabel={(option) => option}
            filterSelectedOptions
            onChange={(_event, value) => handleMedicationsChange(value)}
            renderInput={(params) => (
                <TextField
                    {...params}
                    label="Patient Therapy Medications"
                    placeholder="Select drugs used for therapy"
                />
            )}
        />
    )
}

export default PatientMedications;