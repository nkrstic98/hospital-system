export interface LoginRequest {
    username: string;
    password: string;
}

export interface LoginResponse {
    token: string;
    user: UserResponse | undefined;
}

export interface ValidateSessionRequest {
    token: string;
}

export interface ValidateSessionResponse {
    user: UserResponse | undefined;
}

export interface LogoutRequest {
    token: string;
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