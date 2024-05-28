export interface GetDepartmentsResponse {
    departments: Map<string, Department>;
}

export interface Department {
    displayName: string;
    physicians: Employee[];
    residents: Employee[];
    nurses: Employee[];
}

export interface Employee {
    id: string;
    fullName: string;
}
