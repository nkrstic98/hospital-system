export interface User {
    id:          string;
    firstname:   string;
    lastname:    string;
    nationalIdentificationNumber: string;
    username: string;
    email: string;
    role:        string;
    team:        string | null;
    permissions: Map<string, string> | null;
}
