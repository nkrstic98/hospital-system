import {User} from "../../types/User.ts";

export interface LoginRequest {
    username: string;
    password: string;
}

export interface LoginResponse {
    token: string;
    user: User | undefined;
}

export interface ValidateSessionRequest {
    token: string;
}

export interface ValidateSessionResponse {
    user: User | undefined;
}

export interface LogoutRequest {
    token: string;
}
