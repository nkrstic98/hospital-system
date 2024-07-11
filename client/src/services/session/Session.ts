import {LoginRequest, LoginResponse, LogoutRequest, ValidateSessionRequest, ValidateSessionResponse} from "./types.ts";
import {User} from "../../types/User.ts";

export class SessionService {
    private readonly baseUrl: string;

    constructor() {
        this.baseUrl = "http://localhost:8080/api/v1/session";
    }

    async Login(loginRequest: LoginRequest): Promise<User | undefined> {
        try {
            const response = await fetch(`${this.baseUrl}/login`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(loginRequest),
            });

            if(!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            const responseBody = await response.text();
            const data = JSON.parse(responseBody) as LoginResponse;
            document.cookie = `authToken=${data.token}`

            if (data.user == undefined) {
                return undefined;
            }

            return {
                firstname: data.user.firstname,
                lastname: data.user.lastname,
                nationalIdentificationNumber: data.user.national_identification_number,
                username: data.user.username,
                email: data.user.email,
                role: data.user.role,
                team: data.user.team,
                permissions: data.user.permissions,
            };
        } catch (error) {
            console.error("Failed to login:", error);
        }
    }

    async Validate(request: ValidateSessionRequest): Promise<User | undefined> {
        try {
            const response = await fetch(`${this.baseUrl}/validate`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(request),
            });

            if(!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            const responseBody = await response.text();
            const data = JSON.parse(responseBody) as ValidateSessionResponse;

            if (data.user == undefined) {
                return undefined;
            }

            return {
                firstname: data.user.firstname,
                lastname: data.user.lastname,
                nationalIdentificationNumber: data.user.national_identification_number,
                username: data.user.username,
                email: data.user.email,
                role: data.user.role,
                team: data.user.team,
                permissions: data.user.permissions,
            };
        } catch (error) {
            console.error("Failed to validate:", error);
        }
    }

    async Logout(request: LogoutRequest): Promise<boolean|undefined> {
        try {
            const response = await fetch(`${this.baseUrl}/logout`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(request),
            });

            if(!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            document.cookie = 'authToken=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;';

            return true;
        } catch (error) {
            console.error("Failed to logout:", error);
        }
    }
}