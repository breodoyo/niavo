import { useAuth } from "../hooks/useAuth";
import { useNavigate } from "react-router-dom";

export default function Dashboard() {
  const { logout } = useAuth();
  const navigate = useNavigate();

  function handleLogout() {
    logout();
    navigate("/login");
  }

  return (
    <div style={{ padding: "32px" }}>
      <h1>Dashboard</h1>
      <p>You're signed in.</p>
      <button onClick={handleLogout}>Log out</button>
    </div>
  );
}