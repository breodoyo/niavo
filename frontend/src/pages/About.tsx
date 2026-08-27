import { Link } from "react-router-dom";
import "./Landing.css";

export default function About() {
  return (
    <div className="landing-page">
      <div className="landing-content">
        <span className="landing-eyebrow">NIAVO</span>
        <h1>About Niavo</h1>
        <p>
          Niavo is a work execution platform built to help small teams and
          organizations manage requests from start to finish. Instead of
          juggling spreadsheets, WhatsApp threads, and email chains, teams
          get one shared place to create workflows, assign work, and track
          progress through clear stages — Received, In Progress, and Done.
        </p>
        <div className="landing-actions">
          <Link to="/" className="landing-btn-secondary">
            Back to home
          </Link>
        </div>
      </div>
    </div>
  );
}