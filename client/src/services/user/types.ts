export interface RegisterUserRequest {
    firstname: string;
    lastname: string;
    national_identification_number: string;
    email: string;
    joining_date: Date;
    role: string;
    team: string | undefined;
}

export interface GetUsersResponse {
    users: UserResponse[];
}

export interface UserResponse {
    firstname:   string;
    lastname:    string;
    national_identification_number: string;
    username: string;
    email: string;
    phone_number: string;
    mailing_address: string;
    city: string;
    state: string;
    zip: string;
    gender: string;
    birthday: Date;
    joining_date: Date;
    verified: boolean;
    archived: boolean;
    role:        string;
    team:        string | null;
    permissions: Array<string> | null;
}
