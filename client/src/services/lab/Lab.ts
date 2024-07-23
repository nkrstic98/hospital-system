import {Lab} from "../../types/Admission.ts";
import {GetAuthorizationToken} from "../../utils/utils.ts";

export class LabService {
    private readonly baseUrl: string;

    constructor() {
        this.baseUrl = 'http://localhost:8080/api/v1/labs';
    }

    async GetLabs(): Promise<Lab[]|undefined> {
        try {
            const response = await fetch(`${this.baseUrl}`, {
                method: "GET",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": GetAuthorizationToken(),
                },
            });

            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            const responseBody = await response.text();

            return JSON.parse(responseBody) as Lab[];
        } catch (error) {
            console.error("Failed to get patient:", error);
        }
    }

    async ProcessLabTest(labId: string): Promise<boolean> {
        try {
            const response = await fetch(`${this.baseUrl}/${labId}/process`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": GetAuthorizationToken(),
                },
            });

            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            return true;
        } catch (error) {
            console.error("Failed to get patient:", error);
            return false;
        }
    }
}
