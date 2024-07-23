import React, {createContext, useEffect, useState} from 'react';
import {SessionService} from "../services/session/Session.ts";
import {User} from "../types/User.ts";
import {useLocation, useNavigate} from "react-router-dom";
import {ValidateSessionRequest} from "../services/session/types.ts";
import {GetAuthorizationToken, RemoveAuthorizationToken} from "../utils/utils.ts";
export interface AuthProviderProps {
    children: React.ReactNode;
}

export interface AuthContextType {
    isAuthenticated: boolean;
    user: User | undefined;
    onLogin: (user: User) => void;
    onLogout: () => void;
}

export const AuthContext = createContext<AuthContextType>({
    isAuthenticated: false,
    user: undefined,
    onLogin: () => {},
    onLogout: () => {},
});

export const useAuth = () => React.useContext(AuthContext);

export const AuthProvider = ({ children }: AuthProviderProps) => {
    const location = useLocation();
    const navigate = useNavigate();
    const sessionService = new SessionService();
    const [isAuthenticated, setIsAuthenticated] = useState(false);
    const [user, setUser] = useState<User | undefined>();

    const handleLogin = async (user: User) => {
        setIsAuthenticated(true);
        setUser(user);

        navigate("/home");
    };

    const handleLogout = () => {
        setIsAuthenticated(false);
        setUser(undefined);

        const authToken = GetAuthorizationToken();
        if (authToken === "") {
            return;
        }

        sessionService.Logout({
            token: authToken,
        }).then(() => {});
    };

    useEffect(() => {
        const authToken = GetAuthorizationToken();
        if (authToken === "") {
            setIsAuthenticated(false);
            setUser(undefined);
            return;
        }

        const request: ValidateSessionRequest = {
            token: authToken,
        }

        sessionService.Validate(request).then(r => {
            if (r == undefined) {
                setIsAuthenticated(false);
                setUser(undefined);
                RemoveAuthorizationToken();
            }
            else {
                setIsAuthenticated(true);
                setUser(r);

                navigate(location.pathname);
            }
        }).catch((e) => {
            console.error("Failed to validate session:", e);
            RemoveAuthorizationToken();
        });
    }, []);

    const value = {
        isAuthenticated,
        user: user,
        onLogin: handleLogin,
        onLogout: handleLogout,
    };

    return (
        <AuthContext.Provider value={value}>
            {children}
        </AuthContext.Provider>
    );
};
