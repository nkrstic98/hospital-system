import './App.css'
import { Route, Routes } from 'react-router-dom'
import Login from "./pages/login/Login.tsx";
import { AuthProvider } from "./router/AuthProvider.tsx";
import ProtectedRoute from "./router/ProtectedRoute.tsx";
import Navbar from "./components/Navbar.tsx";
import EmployeeManagement from "./pages/employees/EmployeeManagement.tsx";
import EmployeeRegister from "./pages/employees/EmployeeRegister.tsx";
import PatientIntake from "./pages/intake/PatientIntake.tsx";
import PatientAdmissionStepper from "./pages/intake/PatientAdmissionStepper.tsx";
import Home from "./pages/global/Home.tsx";
import Admissions from "./pages/patient_admissions/Admissions.tsx";

function App() {
    return (
        <AuthProvider>
            <Navbar />

            <Routes>
                <Route index element={<Login />} />
                <Route path="/login" element={<Login />} />

                <Route path={"/home"} element={
                    <ProtectedRoute>
                        <Home />
                    </ProtectedRoute>
                } />

                <Route path={"/patient-intake"} element={
                    <ProtectedRoute section={"INTAKE"}>
                        <PatientIntake section={"INTAKE"} />
                    </ProtectedRoute>
                } />
                <Route path={"/patient-intake/new-admission"} element={
                    <ProtectedRoute section={"INTAKE"} permission={"WRITE"}>
                        <PatientAdmissionStepper />
                    </ProtectedRoute>
                } />

                <Route path={"/employees"} element={
                    <ProtectedRoute section={"EMPLOYEES"}>
                        <EmployeeManagement />
                    </ProtectedRoute>
                } />
                <Route path={"/employees/register"} element={
                    <ProtectedRoute section={"EMPLOYEES"}  permission={"WRITE"}>
                        <EmployeeRegister />
                    </ProtectedRoute>
                } />

                <Route path={"/patients"} element={
                    <ProtectedRoute section={"PATIENTS"}  permission={"READ"}>
                        <Admissions />
                    </ProtectedRoute>
                } />
            </Routes>
        </AuthProvider>
    )
}

export default App;
