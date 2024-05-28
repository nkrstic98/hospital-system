import React from 'react';
import { Navigate, useLocation } from 'react-router-dom';
import { useAuth } from "./AuthProvider.tsx";

export interface ProtectedRouteProps {
    children: React.ReactNode;
    role: string;
}

const ProtectedRoute = ({ children, role }: ProtectedRouteProps) => {
    const location = useLocation();
    const { isAuthenticated, user } = useAuth();

    if (!isAuthenticated || user === undefined || (role !== "" && user.role !== role)) {
        return <Navigate to="/login" replace state={{ from: location }} />;
    }

    return children;
};

export default ProtectedRoute;