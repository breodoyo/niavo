import { Link } from "react-router-dom";
import "./Landing.css";

export default function Landing() {
  return (
    <div className="landing-page">
      <div className="landing-content">
        <span className="landing-eyebrow">NIAVO</span>
        <h1>Track work from request to completion.</h1>
        <p>
          Niavo replaces scattered spreadsheets and chat messages with a
          single place to receive, assign, and track work — from the moment
          it comes in to the moment it's done.
        </p>
        <div className="landing-actions">
          <Link to="/signup" className="landing-btn-primary">
            Get started
          </Link>
          <Link to="/login" className="landing-btn-secondary">
            Sign in
          </Link>
        </div>

        <div className="landing-stamps" aria-hidden="true">
          <span className="stamp stamp-received">RECEIVED</span>
          <span className="stamp stamp-progress">IN PROGRESS</span>
          <span className="stamp stamp-done">DONE</span>
        </div>
      </div>
    </div>
  );
}