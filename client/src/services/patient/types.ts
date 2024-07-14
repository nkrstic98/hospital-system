export interface RegisterPatientRequest {
    firstname: string;
    lastname: string;
    national_identification_number: string;
    medical_record_number: string;
    email: string;
    phone_number: string;
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
