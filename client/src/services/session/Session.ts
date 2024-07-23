import {LoginRequest, LoginResponse, LogoutRequest, ValidateSessionRequest, ValidateSessionResponse} from "./types.ts";
import {User} from "../../types/User.ts";
import {RemoveAuthorizationToken} from "../../utils/utils.ts";

export class SessionService {
    private readonly baseUrl: string;

    constructor() {
        this.baseUrl = "http://localhost:8080/api/v1/sessions";
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

            return data.user;
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

            return data.user;
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

            RemoveAuthorizationToken();

            return true;
        } catch (error) {
            console.error("Failed to logout:", error);
        }
    }
}