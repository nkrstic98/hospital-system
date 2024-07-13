import TextField from "@mui/material/TextField";

export interface TextInputProps {
    currentStep: number;
    value: string;
    setValue: (value: string) => void;
    error?: boolean;
    isRequired: boolean;
}

type stepInfo = {
    label: string;
    placeholder: string;
}

const stepsAndLabels: Map<number, stepInfo> = new Map([
    [2, {label: "Chief Complaint (CC)", placeholder: "The primary reason the patient is seeking medical care, described in their own words."}],
    [3, {label: "History of Present Illness (HPI)", placeholder: "Detailed description of the symptoms, including onset, duration, severity, and factors that alleviate or exacerbate the condition."}],
    [4, {label: "Past Medical History (PMH)", placeholder: "Information about previous illnesses, hospitalizations, surgeries, and chronic conditions."}],
    [7, {label: "Family History (FH)", placeholder: "Health information about immediate family members, including hereditary conditions."}],
    [8, {label: "Social History", placeholder: "Information about lifestyle factors such as smoking, alcohol use, drug use, occupation, and living conditions."}],
    [9, {label: "Physical Examination", placeholder: "Result of the physical examination of the patient."}],
]);

const TextInput = ({ currentStep, value, setValue, error, isRequired }: TextInputProps) => {
    return (
        <>
            <TextField
                id="outlined-multiline-static"
                label={stepsAndLabels.get(currentStep)?.label}
                value={value}
                fullWidth
                multiline
                rows={10}
                placeholder={stepsAndLabels.get(currentStep)?.placeholder}
                variant="outlined"
                required={isRequired}
                onChange={(e) => setValue(e.target.value)}
                error={error && value === ""}
                helperText={error && value === "" ? "This field is required" : ""}
            />
        </>
    );
}

export default TextInput;