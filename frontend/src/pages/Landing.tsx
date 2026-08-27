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

      <footer className="landing-footer">
        <div className="footer-top">
          <div className="footer-brand">
            <span className="footer-logo">NIAVO</span>
            <p className="footer-tagline">Track work from request to completion.</p>

            <div className="footer-socials">
              <a href="https://github.com/breodoyo/niavo" target="_blank" rel="noopener noreferrer" aria-label="GitHub">
                <svg width="20" height="20" viewBox="0 0 24 24" fill="currentColor">
                  <path d="M12 .5C5.65.5.5 5.65.5 12c0 5.08 3.29 9.39 7.86 10.91.57.1.78-.25.78-.55 0-.27-.01-1.17-.02-2.12-3.2.7-3.88-1.36-3.88-1.36-.52-1.33-1.28-1.68-1.28-1.68-1.05-.72.08-.7.08-.7 1.16.08 1.77 1.19 1.77 1.19 1.03 1.77 2.7 1.26 3.36.96.1-.75.4-1.26.73-1.55-2.55-.29-5.23-1.27-5.23-5.68 0-1.25.45-2.28 1.18-3.08-.12-.29-.51-1.46.11-3.04 0 0 .96-.31 3.15 1.18a10.9 10.9 0 0 1 5.74 0c2.19-1.49 3.15-1.18 3.15-1.18.62 1.58.23 2.75.11 3.04.74.8 1.18 1.83 1.18 3.08 0 4.42-2.69 5.39-5.25 5.67.41.36.78 1.06.78 2.14 0 1.55-.01 2.79-.01 3.17 0 .3.21.66.79.55A11.5 11.5 0 0 0 23.5 12C23.5 5.65 18.35.5 12 .5z"/>
                </svg>
              </a>
              <a href="https://twitter.com" target="_blank" rel="noopener noreferrer" aria-label="Twitter">
                <svg width="20" height="20" viewBox="0 0 24 24" fill="currentColor">
                  <path d="M18.9 2H22l-7.6 8.7L23.3 22h-7l-5.5-6.9L4.5 22H1.3l8.1-9.3L1 2h7.2l5 6.3L18.9 2zm-1.2 18.2h1.7L7.4 3.7H5.6l12.1 16.5z"/>
                </svg>
              </a>
              <a href="https://linkedin.com" target="_blank" rel="noopener noreferrer" aria-label="LinkedIn">
                <svg width="20" height="20" viewBox="0 0 24 24" fill="currentColor">
                  <path d="M20.45 20.45h-3.55v-5.57c0-1.33-.02-3.03-1.85-3.03-1.85 0-2.14 1.45-2.14 2.94v5.66H9.36V9h3.41v1.56h.05c.47-.9 1.63-1.85 3.36-1.85 3.6 0 4.27 2.37 4.27 5.45v6.29zM5.34 7.43a2.06 2.06 0 1 1 0-4.12 2.06 2.06 0 0 1 0 4.12zM7.12 20.45H3.56V9h3.56v11.45z"/>
                </svg>
              </a>
            </div>
          </div>

          <div className="footer-col">
            <span className="footer-heading">Product</span>
            <Link to="/signup">Get started</Link>
            <Link to="/login">Sign in</Link>
          </div>

          <div className="footer-col">
            <span className="footer-heading">How it works</span>
            <a href="#">Create a workflow</a>
            <a href="#">Track work items</a>
            <a href="#">Assign your team</a>
          </div>

          <div className="footer-col">
            <span className="footer-heading">Company</span>
            <a href="/about">About</a>
            <a href="#">Contact</a>
          </div>
        </div>

        <p className="footer-copyright">© 2026 Niavo.</p>
      </footer>
    </div>
  );
}