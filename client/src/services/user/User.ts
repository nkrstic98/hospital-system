import {GetAuthorizationToken} from "../../utils/utils.ts";
import {User} from "../../types/User.ts";
import {GetUsersResponse, RegisterUserRequest} from "./types.ts";

export class UserService {
    private readonly baseUrl: string;

    constructor() {
        this.baseUrl = "http://localhost:8080/api/v1/admin/users";
    }

    async Register(request: RegisterUserRequest): Promise<boolean> {
        try {
            const response = await fetch(`${this.baseUrl}`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": GetAuthorizationToken(),
                },
                body: JSON.stringify(request),
            })

            if(!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            return true;
        } catch (error) {
            console.error("Failed to register user:", error);
            return false;
        }
    }

    async GetUsers(): Promise<User[] | undefined> {
        try {
            const response = await fetch(`${this.baseUrl}`, {
                method: "GET",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": GetAuthorizationToken(),
                },
            })

            if(!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            const responseBody = await response.text();
            const data = JSON.parse(responseBody) as GetUsersResponse;
            return data.users.map(user => {
                return {
                    firstname: user.firstname,
                    lastname: user.lastname,
                    nationalIdentificationNumber: user.national_identification_number,
                    username: user.username,
                    email: user.email,
                    role: user.role,
                    team: user.team,
                    permissions: user.permissions,
                }
            });
        } catch (error) {
            console.error("Failed to get users:", error);
        }
    }
}
