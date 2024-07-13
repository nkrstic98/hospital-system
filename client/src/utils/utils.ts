import {User} from "../types/User.ts";

export function GetAuthorizationToken() {
    const cookieName = "authToken=";
    const cookies = document.cookie.split(';');
    for (let i = 0; i < cookies.length; i++) {
        const cookie = cookies[i].trim();
        if (cookie.startsWith(cookieName)) {
            return cookie.substring(cookieName.length, cookie.length);
        }
    }

    return "";
}

export function GetUserPermission(user: User | undefined, section: string): string | undefined {
    if (user == undefined || user.permissions == undefined) {
        return undefined;
    }

    const userPermissions = new Map<string, string>(Object.entries(user.permissions));

    return userPermissions.get(section);
}
