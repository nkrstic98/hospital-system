import {Patient} from "../../types/Patient.ts";
import {GetAuthorizationToken} from "../../utils/utils.ts";
import {RegisterPatientAdmissionRequest, RegisterPatientRequest} from "./types.ts";
import {Admission} from "../../types/Admission.ts";

export class PatientService {
    private readonly baseUrl: string;

    constructor() {
        this.baseUrl = "http://localhost:8080/api/v1/patients";
    }

    async GetPatient(id: string): Promise<Patient|undefined> {
        try {
            const response = await fetch(`${this.baseUrl}/${id}`, {
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

            return JSON.parse(responseBody) as Patient;
        } catch (error) {
            console.error("Failed to get patient:", error);
            return undefined;
        }
    }

    async Register(request: RegisterPatientRequest): Promise<Patient | undefined> {
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

            const responseBody = await response.text();

            return JSON.parse(responseBody) as Patient;
        } catch(error) {
            console.error("Failed to add patient:", error);
            return undefined;
        }
    }

    async RegisterPatientAdmission(request: RegisterPatientAdmissionRequest): Promise<boolean> {
        try {
            const response = await fetch(`${this.baseUrl}/admissions/register`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": GetAuthorizationToken(),
                },
                body: JSON.stringify(request),
            });

            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            return true;
        } catch (error) {
            console.error("Failed to register patient admission:", error);
            return false;
        }
    }

    async GetActiveAdmissions(): Promise<Admission[] | undefined> {
        try {
            const response = await fetch(`${this.baseUrl}/admissions`, {
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
            const data = JSON.parse(responseBody) as Admission[];

            return data.sort((a, b) => {
                return  new Date(b.admissionTime) > new Date(a.admissionTime) ? 1 : -1;
            });
        } catch (error) {
            console.error("Failed to get patient admissions:", error);
            return undefined;
        }
    }

    // async GetActiveAdmissionsByUser(request: GetActiveAdmissionsByUserRequest): Promise<Admission[] | undefined> {
    //     try {
    //         const response = await fetch(`${this.baseUrl}/admissions`, {
    //             method: "POST",
    //             headers: {
    //                 "Content-Type": "application/json",
    //                 "Authorization": GetAuthorizationToken(),
    //             },
    //             body: JSON.stringify(request),
    //         });
    //
    //         if (!response.ok) {
    //             throw new Error(`HTTP error! status: ${response.status}`);
    //         }
    //
    //         const responseBody = await response.text();
    //         const data = JSON.parse(responseBody) as GetAdmissionsResponse;
    //
    //         return data.admissions.sort((a, b) => {
    //             return  new Date(b.admissionTime) > new Date(a.admissionTime) ? 1 : -1;
    //         });
    //     } catch (error) {
    //         console.error("Failed to get patient admissions:", error);
    //         return undefined;
    //     }
    // }

    // async DischargePatient(admissionId: string): Promise<boolean> {
    //     try {
    //         const response = await fetch(`${this.baseUrl}/admissions/${admissionId}/discharge`, {
    //             method: "PATCH",
    //             headers: {
    //                 "Content-Type": "application/json",
    //                 "Authorization": GetAuthorizationToken(),
    //             },
    //         });
    //
    //         if (!response.ok) {
    //             throw new Error(`HTTP error! status: ${response.status}`);
    //         }
    //
    //         return true;
    //     } catch (error) {
    //         console.error("Failed to discharge patient:", error);
    //         return false;
    //     }
    // }
}