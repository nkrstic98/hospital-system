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

export interface Admission {
    id: string;
    startTime: Date;
    endTime?: Date;
    status: string;
    patient: string;
    department: string;
    physician: string;
}
