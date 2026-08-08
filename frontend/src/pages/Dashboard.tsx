import { useNavigate } from "react-router-dom";

export default function Dashboard() {
  const navigate = useNavigate();

  return (
    <div
      style={{
        minHeight: "100vh",
        background: "#f5f7fb",
        padding: "32px",
      }}
    >
      {/* Header */}
      <div
        style={{
          display: "flex",
          justifyContent: "space-between",
          alignItems: "center",
          marginBottom: "32px",
        }}
      >
        <div>
          <h1
            style={{
              margin: 0,
              fontSize: "32px",
              color: "#172b4d",
            }}
          >
            Dashboard
          </h1>

          <p
            style={{
              marginTop: "8px",
              color: "#6b7280",
              fontSize: "16px",
            }}
          >
            Welcome back. Here's what's happening in Niavo.
          </p>
        </div>
      </div>

      {/* Statistics */}
      <div
        style={{
          display: "grid",
          gridTemplateColumns: "repeat(4, 1fr)",
          gap: "20px",
          marginBottom: "32px",
        }}
      >
        <StatCard
          title="Organizations"
          value="2"
          description="Active organizations"
          icon="🏢"
        />

        <StatCard
          title="Users"
          value="12"
          description="Registered users"
          icon="👥"
        />

        <StatCard
          title="Workflows"
          value="5"
          description="Active workflows"
          icon="🔄"
        />

        <StatCard
          title="Work Items"
          value="24"
          description="Total work items"
          icon="📋"
        />
      </div>

      {/* Main content */}
      <div
        style={{
          display: "grid",
          gridTemplateColumns: "2fr 1fr",
          gap: "24px",
        }}
      >
        {/* Work Items */}
        <div
          style={{
            background: "#ffffff",
            borderRadius: "12px",
            padding: "24px",
            boxShadow: "0 2px 8px rgba(0, 0, 0, 0.06)",
          }}
        >
          <div
            style={{
              display: "flex",
              justifyContent: "space-between",
              alignItems: "center",
              marginBottom: "24px",
            }}
          >
            <div>
              <h2
                style={{
                  margin: 0,
                  color: "#172b4d",
                  fontSize: "20px",
                }}
              >
                Work Items
              </h2>

              <p
                style={{
                  margin: "6px 0 0",
                  color: "#6b7280",
                }}
              >
                Overview of your current workload
              </p>
            </div>

            <button
              onClick={() => navigate("/workflows")}
              style={secondaryButton}
            >
              View workflows
            </button>
          </div>

          <StatusRow
            label="Received"
            value={10}
            percentage={42}
          />

          <StatusRow
            label="In Progress"
            value={8}
            percentage={33}
          />

          <StatusRow
            label="Completed"
            value={6}
            percentage={25}
          />
        </div>

        {/* Quick Actions */}
        <div
          style={{
            background: "#ffffff",
            borderRadius: "12px",
            padding: "24px",
            boxShadow: "0 2px 8px rgba(0, 0, 0, 0.06)",
          }}
        >
          <h2
            style={{
              margin: "0 0 8px",
              color: "#172b4d",
              fontSize: "20px",
            }}
          >
            Quick Actions
          </h2>

          <p
            style={{
              margin: "0 0 24px",
              color: "#6b7280",
            }}
          >
            Get things done quickly.
          </p>

          <button
            onClick={() => navigate("/organizations")}
            style={primaryButton}
          >
            + Create Organization
          </button>

          <button
            onClick={() => navigate("/workflows")}
            style={actionButton}
          >
            + Create Workflow
          </button>

          <button
            onClick={() => navigate("/users")}
            style={actionButton}
          >
            + Add User
          </button>
        </div>
      </div>

      {/* Recent Activity */}
      <div
        style={{
          marginTop: "24px",
          background: "#ffffff",
          borderRadius: "12px",
          padding: "24px",
          boxShadow: "0 2px 8px rgba(0, 0, 0, 0.06)",
        }}
      >
        <h2
          style={{
            margin: "0 0 20px",
            color: "#172b4d",
            fontSize: "20px",
          }}
        >
          Recent Activity
        </h2>

        <Activity
          title="Work item created"
          description="Implement user authentication"
          time="Today"
        />

        <Activity
          title="Workflow created"
          description="Development workflow"
          time="Today"
        />

        <Activity
          title="Organization updated"
          description="Zone01 Kisumu"
          time="Yesterday"
        />
      </div>
    </div>
  );
}

/* -------------------------------- */
/* Components                       */
/* -------------------------------- */

function StatCard({
  title,
  value,
  description,
  icon,
}: {
  title: string;
  value: string;
  description: string;
  icon: string;
}) {
  return (
    <div
      style={{
        background: "#ffffff",
        borderRadius: "12px",
        padding: "22px",
        boxShadow: "0 2px 8px rgba(0, 0, 0, 0.06)",
      }}
    >
      <div
        style={{
          display: "flex",
          justifyContent: "space-between",
          alignItems: "center",
        }}
      >
        <div>
          <p
            style={{
              margin: 0,
              color: "#6b7280",
              fontSize: "14px",
            }}
          >
            {title}
          </p>

          <h2
            style={{
              margin: "8px 0",
              fontSize: "30px",
              color: "#172b4d",
            }}
          >
            {value}
          </h2>

          <p
            style={{
              margin: 0,
              color: "#9ca3af",
              fontSize: "13px",
            }}
          >
            {description}
          </p>
        </div>

        <div
          style={{
            fontSize: "28px",
            background: "#f0f4ff",
            borderRadius: "10px",
            padding: "10px",
          }}
        >
          {icon}
        </div>
      </div>
    </div>
  );
}

function StatusRow({
  label,
  value,
  percentage,
}: {
  label: string;
  value: number;
  percentage: number;
}) {
  return (
    <div style={{ marginBottom: "20px" }}>
      <div
        style={{
          display: "flex",
          justifyContent: "space-between",
          marginBottom: "8px",
        }}
      >
        <span
          style={{
            color: "#374151",
            fontSize: "14px",
          }}
        >
          {label}
        </span>

        <span
          style={{
            color: "#6b7280",
            fontSize: "14px",
          }}
        >
          {value}
        </span>
      </div>

      <div
        style={{
          height: "8px",
          background: "#e5e7eb",
          borderRadius: "10px",
          overflow: "hidden",
        }}
      >
        <div
          style={{
            width: `${percentage}%`,
            height: "100%",
            background: "#172b4d",
            borderRadius: "10px",
          }}
        />
      </div>
    </div>
  );
}

function Activity({
  title,
  description,
  time,
}: {
  title: string;
  description: string;
  time: string;
}) {
  return (
    <div
      style={{
        display: "flex",
        justifyContent: "space-between",
        alignItems: "center",
        padding: "16px 0",
        borderTop: "1px solid #e5e7eb",
      }}
    >
      <div>
        <strong
          style={{
            color: "#172b4d",
          }}
        >
          {title}
        </strong>

        <p
          style={{
            margin: "4px 0 0",
            color: "#6b7280",
            fontSize: "14px",
          }}
        >
          {description}
        </p>
      </div>

      <span
        style={{
          color: "#9ca3af",
          fontSize: "13px",
        }}
      >
        {time}
      </span>
    </div>
  );
}

/* -------------------------------- */
/* Styles                           */
/* -------------------------------- */

const primaryButton: React.CSSProperties = {
  width: "100%",
  padding: "12px 16px",
  marginBottom: "12px",
  border: "none",
  borderRadius: "7px",
  background: "#172b4d",
  color: "#ffffff",
  cursor: "pointer",
  fontWeight: 600,
};

const secondaryButton: React.CSSProperties = {
  padding: "9px 14px",
  border: "1px solid #d1d5db",
  borderRadius: "7px",
  background: "#ffffff",
  color: "#172b4d",
  cursor: "pointer",
};

const actionButton: React.CSSProperties = {
  width: "100%",
  padding: "12px 16px",
  marginBottom: "12px",
  border: "1px solid #d1d5db",
  borderRadius: "7px",
  background: "#ffffff",
  color: "#172b4d",
  cursor: "pointer",
  fontWeight: 500,
};