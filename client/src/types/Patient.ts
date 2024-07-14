import {Admission} from "./Admission.ts";

export interface Patient {
    id: string;
    firstname: string;
    lastname: string;
    nationalIdentificationNumber: string;
    medicalRecordNumber: string;
    email: string;
    phoneNumber: string;
    admissions: Array<Admission> | null;
}
