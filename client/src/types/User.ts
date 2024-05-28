export interface User {
    firstname:   string;
    lastname:    string;
    nationalIdentificationNumber: string;
    username: string;
    email: string;
    phoneNumber: string;
    mailingAddress: string;
    city: string;
    state: string;
    zip: string;
    gender: string;
    birthday: Date;
    joiningDate: Date;
    verified: boolean;
    archived: boolean;
    role:        string;
    team:        string | null;
    permissions: Array<string> | null;
}
