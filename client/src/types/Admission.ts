export interface Admission {
    id: string;
    startTime: Date;
    endTime?: Date;
    status: string;
    patient: string;
    department: string;
    physician: string;
}
