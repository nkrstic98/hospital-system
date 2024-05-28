export interface GetScheduleResponse {
    schedule: Schedule[];
}

export interface Schedule {
    startDate: string;
    endDate: string;
    // department -> day -> shift -> employee_type -> ids
    schedule: Map<
        string,
        Map<
            string,
            Map<
                string,
                Map<
                    string,
                    Map<
                        string,
                        string[]
                    >
                >
            >
        >
    >
}
