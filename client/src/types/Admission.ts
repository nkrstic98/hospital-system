export interface Admission {
    id: string;
    patient: string;
    department: string;
    physician: string;
    admissionTime: Date;
    status: string;
}
