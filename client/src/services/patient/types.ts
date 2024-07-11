import {Patient} from "../../types/Patient.ts";
import {Admission} from "../../types/Admission.ts";

export interface PatientGetResponse {
    patient: Patient | undefined;
}

export interface RegisterPatientRequest {
    firstname: string;
    lastname: string;
    nationalIdentificationNumber: string;
    medicalRecordNumber: string;
    email: string;
    phoneNumber: string;
}

export interface RegisterPatientResponse {
    patient: Patient | undefined;
}

export interface RegisterPatientAdmissionRequest {
    patientId: string;
    department: string; // team in rbac
    physician: string; // actor id
    symptoms: string;
    medications: string[];
    allergies: string[];
}

export interface GetAdmissionsRequest {
    statuses: string[];
}

export interface GetAdmissionsResponse {
    admissions: Admission[];
}
