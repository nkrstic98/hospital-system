import './App.css'
import { Route, Routes } from 'react-router-dom'
import Login from "./pages/global/Login.tsx";
import { AuthProvider } from "./router/AuthProvider.tsx";
import ProtectedRoute from "./router/ProtectedRoute.tsx";
import {Page} from "./pages/user/Page.tsx";
import Navbar from "./pages/Navbar.tsx";
import EmployeeManagement from "./pages/admin/EmployeeManagement.tsx";
import EmployeeRegister from "./pages/admin/EmployeeRegister.tsx";
import PatientIntake from "./pages/admin/PatientIntake.tsx";
import PatientAdmissionStepper from "./pages/admin/PatientAdmissionStepper.tsx";
import DepartmentManagement from "./pages/admin/DepartmentManagement.tsx";

function App() {
    return (
        <AuthProvider>
            <Navbar />

            <Routes>
                <Route index element={<Login />} />
                <Route path="/login" element={<Login />} />
                {/*<Route path="/account" element={*/}
                {/*    <ProtectedRoute role="">*/}
                {/*        <Account />*/}
                {/*    </ProtectedRoute>*/}
                {/*} />*/}
                <Route path={"/admin"} element={
                    <ProtectedRoute role={"ADMIN"}>
                        <PatientIntake />
                    </ProtectedRoute>
                } />
                <Route path={"/admin/employees"} element={
                    <ProtectedRoute role={"ADMIN"}>
                        <EmployeeManagement />
                    </ProtectedRoute>
                } />
                <Route path={"/admin/employees/register"} element={
                    <ProtectedRoute role={"ADMIN"}>
                        <EmployeeRegister />
                    </ProtectedRoute>
                } />
                <Route path={"/admin/patients/admissions"} element={
                    <ProtectedRoute role={"ADMIN"}>
                        <PatientAdmissionStepper />
                    </ProtectedRoute>
                } />
                <Route path={"/admin/departments"} element={
                    <ProtectedRoute role={"ADMIN"}>
                        <DepartmentManagement />
                    </ProtectedRoute>
                } />
                <Route path={"/user"} element={
                    <ProtectedRoute role={"ATTENDING"}>
                        <Page />
                    </ProtectedRoute>
                } />
            </Routes>
        </AuthProvider>
    )
}

export default App;
