import {Admission, Patient} from "../../types/Patient.ts";
import {GetAuthorizationToken} from "../../utils/utils.ts";
import {OrderLabTestRequest, RegisterPatientAdmissionRequest, RegisterPatientRequest} from "./types.ts";
import {AdmissionDetails} from "../../types/Admission.ts";

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

    async RegisterPatient(request: RegisterPatientRequest): Promise<Patient | undefined> {
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
            const response = await fetch(`${this.baseUrl}/admissions`, {
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

    async GetActiveAdmissions(userId: string): Promise<Admission[] | undefined> {
        try {
            const response = await fetch(`${this.baseUrl}${userId != "" ? `/users/${userId}` : ""}/admissions`, {
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
            return JSON.parse(responseBody) as Admission[];
        } catch (error) {
            console.error("Failed to get patient admissions:", error);
            return undefined;
        }
    }

    async GetAdmissionDetails(id: string): Promise<AdmissionDetails | undefined> {
        try {
            const response = await fetch(`${this.baseUrl}/admissions/${id}`, {
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
            return JSON.parse(responseBody) as AdmissionDetails;
        } catch (error) {
            console.error("Failed to get patient admissions:", error);
            return undefined;
        }
    }

    async AcknowledgeTransfer(id: string, accept: boolean): Promise<boolean> {
        try {
            const response = await fetch(`${this.baseUrl}/admissions/${id}/transfer?accept=${accept}`, {
                method: "GET",
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
            console.error("Failed to acknowledge admission:", error);
            return false;
        }
    }

    async AdmitPatient(id: string): Promise<boolean> {
        try {
            const response = await fetch(`${this.baseUrl}/admissions/${id}/admit`, {
                method: "GET",
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
            console.error("Failed to admit patient:", error);
            return false;
        }
    }

    async UpdateAdmissionDetails(admission: AdmissionDetails) : Promise<AdmissionDetails|undefined> {
        try {
            const response = await fetch(`${this.baseUrl}/admissions`, {
                method: "PUT",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": GetAuthorizationToken(),
                },
                body: JSON.stringify(admission),
            });

            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            const responseBody = await response.text();
            return JSON.parse(responseBody) as AdmissionDetails;
        } catch (error) {
            console.error("Failed to update admission details:", error);
            return undefined;
        }
    }

    async OrderLabTest(req: OrderLabTestRequest): Promise<boolean> {
        try {
            const response = await fetch(`${this.baseUrl}/labs`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": GetAuthorizationToken(),
                },
                body: JSON.stringify(req),
            });

            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            return true;
        } catch (error) {
            console.error("Failed to order lab test:", error);
            return false;
        }
    }

    async AddTeamMember(admissionId: string, userId: string): Promise<boolean> {
        try {
            const response = await fetch(`${this.baseUrl}/admissions/add-team-member`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": GetAuthorizationToken(),
                },
                body: JSON.stringify({
                    admissionId: admissionId,
                    userId: userId,
                }),
            });

            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            return true;
        } catch (error) {
            console.error("Failed to add team member:", error);
            return false;
        }
    }

    async AddTeamMemberPermissions(admissionId: string, userId: string, section: string, permission: string): Promise<boolean> {
        try {
            const response = await fetch(`${this.baseUrl}/admissions/add-team-member-permissions`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": GetAuthorizationToken(),
                },
                body: JSON.stringify({
                    admissionId: admissionId,
                    userId: userId,
                    section: section,
                    permission: permission,
                }),
            });

            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            return true;
        } catch (error) {
            console.error("Failed to add team member permissions:", error);
            return false;
        }
    }

    async RemoveTeamMember(admissionId: string, userId: string): Promise<boolean> {
        try {
            const response = await fetch(`${this.baseUrl}/admissions/remove-team-member`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": GetAuthorizationToken(),
                },
                body: JSON.stringify({
                    admissionId: admissionId,
                    userId: userId,
                }),
            });

            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            return true;
        } catch (error) {
            console.error("Failed to remove team member:", error);
            return false;
        }
    }

    async RemoveTeamMemberPermissions(admissionId: string, userId: string, section: string): Promise<boolean> {
        try {
            const response = await fetch(`${this.baseUrl}/admissions/remove-team-member-permissions`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": GetAuthorizationToken(),
                },
                body: JSON.stringify({
                    admissionId: admissionId,
                    userId: userId,
                    section: section,
                }),
            });

            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            return true;
        } catch (error) {
            console.error("Failed to remove team member permissions:", error);
            return false;
        }
    }

    async RequestPatientTransfer(admissionId: string, department: string, doctor: string): Promise<boolean> {
        try {
            const response = await fetch(`${this.baseUrl}/admissions/request-transfer`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": GetAuthorizationToken(),
                },
                body: JSON.stringify({
                    admissionId: admissionId,
                    toTeam: department,
                    toTeamLead: doctor,
                }),
            });

            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            return true;
        } catch (error) {
            console.error("Failed to request patient transfer:", error);
            return false;
        }
    }

    async DischargePatient(admissionId: string): Promise<boolean> {
        try {
            const response = await fetch(`${this.baseUrl}/admissions/${admissionId}/discharge`, {
                method: "PATCH",
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
            console.error("Failed to discharge patient:", error);
            return false;
        }
    }
}
