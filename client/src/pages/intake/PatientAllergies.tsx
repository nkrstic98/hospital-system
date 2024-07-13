import {Checkbox, FormControl, FormControlLabel, FormGroup} from "@mui/material";

const commonDrugAllergies = [
    'Penicillin and other antibiotics',
    'Nonsteroidal anti-inflammatory drugs (NSAIDs)',
    'Chemotherapy drugs',
    'Monoclonal antibody therapy',
    'Anticonvulsants',
    'Insulin',
    'Radiocontrast media',
    'Anesthetics',
    'Vaccines',
    'Biological drugs',
    'Antiretroviral drugs',
    'Antipsychotic medications',
    'ACE inhibitors',
    'Statins',
    'Sulfonamides'
];

export interface PatientAllergiesProps {
    checkedAllergies: string[];
    handleCheckboxChange: (value: string, checked: boolean) => void;
}

const PatientAllergies = ({ checkedAllergies, handleCheckboxChange }: PatientAllergiesProps) => {
    return (
        <FormControl component="fieldset">
            <FormGroup>
                {commonDrugAllergies.map((allergy, index) => (
                    <FormControlLabel
                        key={index}
                        control={<Checkbox
                            checked={checkedAllergies.includes(allergy)}
                            onChange={(e) => handleCheckboxChange(e.target.value, e.target.checked)}
                            value={allergy}
                        />}
                        label={allergy}
                    />
                ))}
            </FormGroup>
        </FormControl>
    )
}

export default PatientAllergies;