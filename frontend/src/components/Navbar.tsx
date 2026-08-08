import { Link, useNavigate } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";
import "./Navbar.css";

export function Navbar() {
  const { logout } = useAuth();
  const navigate = useNavigate();

  function handleLogout() {
    logout();
    navigate("/login");
  }

  return (
    <nav className="navbar">
      <span className="navbar-brand">NIAVO</span>
      <div className="navbar-links">
        <Link to="/dashboard">Dashboard</Link>
        <Link to="/organizations">Organizations</Link>
        <Link to="/users">Users</Link>
        <Link to="/workflows">Workflows</Link>
      </div>
      <button className="navbar-logout" onClick={handleLogout}>
        Log out
      </button>
    </nav>
  );
}