import {Patient} from "./Patient.ts";
import {User} from "./User.ts";

export interface Anamnesis {
    chiefComplaint: string;
    historyOfPresentIllness: string;
    pastMedicalHistory: string;
    medications: string[];
    allergies: string[];
    familyHistory: string;
    socialHistory: string;
    physicalExamination: string;
}

export interface Vitals {
    bodyTemperature: string;
    heartRate: string;
    bloodPressure: string;
    respiratoryRate: string;
    oxygenSaturation: string;
    painLevel: string;
    bloodGlucoseLevel: string;
}

export interface Medication {
    medication: string;
    dose: string;
}

export interface Lab {
    id: string;
    requestedAt: Date;
    processedAt?: Date;
    testType: string;
    testResults?: LabTest[];
    requestedBy: string;
    processedBy?: string;
}

export interface LabTest {
    name: string;
    unit: string;
    minValue: number;
    maxValue: number;
    referenceRange: string;
    result: number;
}

export interface Log {
    timestamp: Date;
    action: string;
    message: string;
    details: string;
    performedBy: string;
}

export interface JourneyStep {
    transferTime: Date;
    fromTeam: string;
    toTeam: string;
    fromTeamLead: string;
    toTeamLead: string;
}

export interface CareTeam {
    team: string;
    department: string;
    teamLead: string;
    assignments: Map<string, User>
    journey: Array<JourneyStep>;
    pendingTransfer?: JourneyStep;
}

export interface AdmissionDetails {
    id: string;
    startTime: Date;
    endTime?: Date;
    status: string;

    anamnesis: Anamnesis;
    vitals: Vitals;
    diagnosis: string | null;
    medications: Medication[] | null;
    labs: Lab[] | null;

    logs: Log[];

    patient: Patient;
    careTeam: CareTeam;
}
