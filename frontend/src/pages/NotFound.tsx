import { Link } from "react-router-dom";

export default function NotFound() {
  return (
    <div style={{ padding: "48px", textAlign: "center" }}>
      <h1>Page not found</h1>
      <p>The page you're looking for doesn't exist.</p>
      <Link to="/login">Go to sign in</Link>
    </div>
  );
}