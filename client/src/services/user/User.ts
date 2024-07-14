import {GetAuthorizationToken} from "../../utils/utils.ts";
import {User} from "../../types/User.ts";
import {Department, GetDepartmentsResponse, GetDepartments, RegisterUserRequest} from "./types.ts";

export class UserService {
    private readonly baseUrl: string;

    constructor() {
        this.baseUrl = "http://localhost:8080/api/v1/users";
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
            return JSON.parse(responseBody) as User[];
        } catch (error) {
            console.error("Failed to get users:", error);
        }
    }

    async GetDepartments(request: GetDepartments): Promise<Map<string, Department> | undefined> {
        try {
            const response = await fetch(`${this.baseUrl}/departments`, {
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

            const responseBody = await response.text();
            const data = JSON.parse(responseBody) as GetDepartmentsResponse;
            return new Map<string, Department>(Object.entries(data.departments));
        } catch (error) {
            console.error("Failed to get departments:", error);
            return undefined;
        }
    }
}
