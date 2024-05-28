import {Admission} from "./Admission.ts";

export interface Patient {
    id: string;
    firstname: string;
    lastname: string;
    nationalIdentificationNumber: string;
    medicalRecordNumber: string;
    birthday: Date;
    gender: string;
    email: string;
    phoneNumber: string;
    admissions: Array<Admission>;
}