import React from 'react';
import { Navigate, useLocation } from 'react-router-dom';
import { useAuth } from "./AuthProvider.tsx";
import {GetUserPermission} from "../utils/utils.ts";

export interface ProtectedRouteProps {
    children: React.ReactNode;
    section?: string;
    permission?: string;
}

const ProtectedRoute = ({ children, section, permission }: ProtectedRouteProps) => {
    const location = useLocation();
    const { isAuthenticated, user } = useAuth();

    if (!isAuthenticated || user === undefined) {
        return <Navigate to="/login" replace state={{ from: location }} />;
    }

    if (section == undefined) {
        return children;
    }

    const assignedPermission = GetUserPermission(user, section);

    if(assignedPermission == undefined) {
        return <Navigate to="/login" replace state={{ from: location }} />;
    }

    if (permission !== undefined) {
        switch (permission) {
            case "READ":
                if (assignedPermission == "READ" || assignedPermission == "WRITE") {
                    return children;
                }
                break;
            case "WRITE":
                if (assignedPermission == "WRITE") {
                    return children;
                }
                break;
        }

        return <Navigate to="/login" replace state={{ from: location }} />;
    }

    return children;
};

export default ProtectedRoute;