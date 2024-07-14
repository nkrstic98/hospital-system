export interface RegisterUserRequest {
    firstname: string;
    lastname: string;
    national_identification_number: string;
    email: string;
    joining_date: Date;
    role: string;
    team: string | undefined;
}

export interface UserResponse {
    firstname:   string;
    lastname:    string;
    national_identification_number: string;
    username: string;
    email: string;
    role:        string;
    team:        string | null;
    permissions: Map<string, string> | null;
}
