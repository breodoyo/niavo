import { useState, type FormEvent } from "react";
import { useNavigate, Link, useLocation } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";
import "./Login.css";

export default function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState<string | null>(null);
  const [isSubmitting, setIsSubmitting] = useState(false);

  const { login } = useAuth();
  const location = useLocation();
  const from = (location.state as { from?: Location})?.from?.pathname || "/dashboard";
  const navigate = useNavigate();

  async function handleSubmit(e: FormEvent) {
    e.preventDefault();
    setError(null);
    setIsSubmitting(true);

    try {
      await login({ email, password });
      navigate(from, { replace: true });
    } catch (err) {
      setError("Incorrect email or password.");
    } finally {
      setIsSubmitting(false);
    }
  }

  return (
    <div className="login-page">
      <div className="login-panel">
        <div className="ticket-stack" aria-hidden="true">
          <div className="ticket ticket-1">
            <span className="stamp stamp-received">RECEIVED</span>
          </div>
          <div className="ticket ticket-2">
            <span className="stamp stamp-progress">IN PROGRESS</span>
          </div>
          <div className="ticket ticket-3">
            <span className="stamp stamp-done">DONE</span>
          </div>
        </div>
      </div>

      <div className="login-form-panel">
        <form className="login-form" onSubmit={handleSubmit}>
          <div className="login-header">
            <span className="login-eyebrow">NIAVO</span>
            <h1>Sign in</h1>
            <p>Track work from request to completion.</p>
          </div>

          {error && <div className="login-error">{error}</div>}

          <label className="login-field">
            <span>Email</span>
            <input
              type="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
              autoComplete="email"
            />
          </label>

          <label className="login-field">
            <span>Password</span>
            <input
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
              autoComplete="current-password"
            />
          </label>

          <button type="submit" disabled={isSubmitting}>
            {isSubmitting ? "Signing in…" : "Sign in"}
          </button>

          <p className="login-footer">
            New here? <Link to="/signup">Create an account</Link>
          </p>
        </form>
      </div>
    </div>
  );
}