import { Outlet } from "react-router-dom";
import { Navbar } from "../components/Navbar";

export function DashboardLayout() {
  return (
    <div>
      <Navbar />
      <main style={{ padding: "0" }}>
        <Outlet />
      </main>
    </div>
  );
}