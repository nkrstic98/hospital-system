import {GetAuthorizationToken} from "../../utils/utils.ts";
import {GetScheduleResponse} from "./types.ts";

export class ScheduleService {
    private readonly baseUrl: string;

    constructor() {
        this.baseUrl = "http://localhost:8080/api/v1/admin/schedule";
    }

    async GetSchedule(): Promise<GetScheduleResponse | undefined> {
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
            return JSON.parse(responseBody) as GetScheduleResponse;
        } catch (error) {
            console.error("Failed to get schedule:", error);
            return undefined;
        }
    }
}