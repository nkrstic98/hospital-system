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
    department: string;
    physician: string;
    chief_complaint: string;
    history_of_present_illness: string;
    past_medical_history: string;
    medications: string[];
    allergies: string[];
    family_history: string;
    social_history: string;
    physical_examination: string;
}

export interface GetAdmissionsRequest {
    statuses: string[];
}

export interface GetActiveAdmissionsByUserRequest {
    userId: string;
    statuses: string[];
}

export interface GetAdmissionsResponse {
    admissions: Admission[];
}
