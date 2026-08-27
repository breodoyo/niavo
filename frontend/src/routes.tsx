import { Routes, Route } from "react-router-dom";
import Login from "./pages/Login";
import Signup from "./pages/Signup";
import Dashboard from "./pages/Dashboard";
import Organizations from "./pages/Organizations";
import Users from "./pages/Users";
import Workflows from "./pages/Workflows";
import WorkItems from "./pages/WorkItems";
import NotFound from "./pages/NotFound";
import Landing from "./pages/Landing";
import About from "./pages/About";
import { ProtectedRoute } from "./components/ProtectedRoute";
import { DashboardLayout } from "./layouts/DashboardLayout";

export function AppRoutes() {
  return (
    <Routes>
      <Route path="/" element={<Landing />} />
      <Route path="/login" element={<Login />} />
      <Route path="/signup" element={<Signup />} />

      <Route
        element={
          <ProtectedRoute>
            <DashboardLayout />
          </ProtectedRoute>
        }
      >
        <Route path="/dashboard" element={<Dashboard />} />
        <Route path="/organizations" element={<Organizations />} />
        <Route path="/users" element={<Users />} />
        <Route path="/workflows" element={<Workflows />} />
        <Route path="/workflows/:workflowId/items" element={<WorkItems />} />
        <Route path="/about" element={<About />} />
      </Route>

      <Route path="*" element={<NotFound />} />
    </Routes>
  );
}