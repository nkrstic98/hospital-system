import {User} from "../../types/User.ts";

export interface RegisterUserRequest {
    firstname: string;
    lastname: string;
    nationalIdentificationNumber: string;
    email: string;
    role: string;
    team: string | undefined;
}

export interface GetDepartments {
    team: string | undefined;
    role: string | undefined;
}

export interface GetDepartmentsResponse {
    departments: Map<string, Department>;
}

export interface Department {
    displayName: string;
    users: User[];
}
