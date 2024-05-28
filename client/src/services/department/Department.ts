import {Department, GetDepartmentsResponse} from "./types.ts";
import {GetAuthorizationToken} from "../../utils/utils.ts";

export class DepartmentService {
    private readonly baseUrl: string;

    constructor() {
        this.baseUrl = "http://localhost:8080/api/v1/admin/departments";
    }

    async GetDepartments(): Promise<Map<string, Department> | undefined> {
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
            const data = JSON.parse(responseBody) as GetDepartmentsResponse;
            return new Map<string, Department>(Object.entries(data.departments));
        } catch (error) {
            console.error("Failed to get departments:", error);
            return undefined;
        }
    }
}