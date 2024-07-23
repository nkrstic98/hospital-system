export interface RegisterPatientRequest {
    firstname: string;
    lastname: string;
    nationalIdentificationNumber: string;
    medicalRecordNumber: string;
    email: string;
    phoneNumber: string;
}

export interface RegisterPatientAdmissionRequest {
    patientId: string;
    department: string;
    physician: string;
    chiefComplaint: string;
    historyOfPresentIllness: string;
    pastMedicalHistory: string;
    medications: string[];
    allergies: string[];
    familyHistory: string;
    socialHistory: string;
    physicalExamination: string;
    admittedBy: string;
    admittedByTeam: string;
}

export interface OrderLabTestRequest {
    admissionId: string;
    labTest: string;
}