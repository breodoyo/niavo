import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { listOrganizations } from "../services/organizations";
import { listUsers } from "../services/users";
import { listWorkflows } from "../services/workflows";
import { listWorkItems, type WorkItem } from "../services/workitems";

export default function Dashboard() {
  const navigate = useNavigate();

  const [orgCount, setOrgCount] = useState<number | null>(null);
  const [userCount, setUserCount] = useState<number | null>(null);
  const [workflowCount, setWorkflowCount] = useState<number | null>(null);
  const [allItems, setAllItems] = useState<WorkItem[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [loadError, setLoadError] = useState<string | null>(null);

  useEffect(() => {
    load();
  }, []);

  async function load() {
    setIsLoading(true);
    setLoadError(null);
    try {
      const [orgs, users, workflows] = await Promise.all([
        listOrganizations(),
        listUsers(),
        listWorkflows(),
      ]);

      setOrgCount(orgs.length);
      setUserCount(users.length);
      setWorkflowCount(workflows.length);

      const itemLists = await Promise.all(
        workflows.map((wf) => listWorkItems(wf.id).catch(() => []))
      );
      setAllItems(itemLists.flat());
    } catch (err) {
      setLoadError("Couldn't load dashboard data.");
    } finally {
      setIsLoading(false);
    }
  }

  const received = allItems.filter((i) => i.status === "received").length;
  const inProgress = allItems.filter((i) => i.status === "in_progress").length;
  const done = allItems.filter((i) => i.status === "done").length;
  const total = allItems.length;

  function pct(n: number) {
    return total === 0 ? 0 : Math.round((n / total) * 100);
  }

  return (
    <div style={{ minHeight: "100vh", background: "var(--color-paper)", padding: "32px" }}>
      <div style={{ marginBottom: "32px" }}>
        <h1 style={{ margin: 0, fontSize: "32px", color: "var(--color-text)" }}>Dashboard</h1>
        <p style={{ marginTop: "8px", color: "var(--color-text-secondary)", fontSize: "16px" }}>
          Welcome back. Here's what's happening in Niavo.
        </p>
      </div>

      {loadError && (
        <div style={{ marginBottom: "24px", color: "var(--color-error)" }}>{loadError}</div>
      )}

      <div style={{ display: "grid", gridTemplateColumns: "repeat(4, 1fr)", gap: "20px", marginBottom: "32px" }}>
        <StatCard title="Organizations" value={isLoading ? "…" : String(orgCount)} description="Active organizations" icon="🏢" />
        <StatCard title="Users" value={isLoading ? "…" : String(userCount)} description="Registered users" icon="👥" />
        <StatCard title="Workflows" value={isLoading ? "…" : String(workflowCount)} description="Active workflows" icon="🔄" />
        <StatCard title="Work Items" value={isLoading ? "…" : String(total)} description="Total work items" icon="📋" />
      </div>

      <div style={{ display: "grid", gridTemplateColumns: "2fr 1fr", gap: "24px" }}>
        <div style={{ background: "var(--color-surface)", borderRadius: "var(--radius-lg)", padding: "24px", boxShadow: "var(--shadow-sm)" }}>
          <div style={{ display: "flex", justifyContent: "space-between", alignItems: "center", marginBottom: "24px" }}>
            <div>
              <h2 style={{ margin: 0, color: "var(--color-text)", fontSize: "20px" }}>Work Items</h2>
              <p style={{ margin: "6px 0 0", color: "var(--color-text-secondary)" }}>Overview of your current workload</p>
            </div>
            <button onClick={() => navigate("/workflows")} style={secondaryButton}>View workflows</button>
          </div>

          {isLoading ? (
            <p style={{ color: "var(--color-text-secondary)" }}>Loading…</p>
          ) : total === 0 ? (
            <p style={{ color: "var(--color-text-secondary)" }}>No work items yet.</p>
          ) : (
            <>
              <StatusRow label="Received" value={received} percentage={pct(received)} />
              <StatusRow label="In Progress" value={inProgress} percentage={pct(inProgress)} />
              <StatusRow label="Completed" value={done} percentage={pct(done)} />
            </>
          )}
        </div>

        <div style={{ background: "var(--color-surface)", borderRadius: "var(--radius-lg)", padding: "24px", boxShadow: "var(--shadow-sm)" }}>
          <h2 style={{ margin: "0 0 8px", color: "var(--color-text)", fontSize: "20px" }}>Quick Actions</h2>
          <p style={{ margin: "0 0 24px", color: "var(--color-text-secondary)" }}>Get things done quickly.</p>
          <button onClick={() => navigate("/organizations")} style={primaryButton}>+ Create Organization</button>
          <button onClick={() => navigate("/workflows")} style={actionButton}>+ Create Workflow</button>
          <button onClick={() => navigate("/users")} style={actionButton}>+ Add User</button>
        </div>
      </div>
    </div>
  );
}

function StatCard({ title, value, description, icon }: { title: string; value: string; description: string; icon: string }) {
  return (
    <div style={{ background: "var(--color-surface)", borderRadius: "var(--radius-lg)", padding: "22px", boxShadow: "var(--shadow-sm)" }}>
      <div style={{ display: "flex", justifyContent: "space-between", alignItems: "center" }}>
        <div>
          <p style={{ margin: 0, color: "var(--color-text-secondary)", fontSize: "14px" }}>{title}</p>
          <h2 style={{ margin: "8px 0", fontSize: "30px", color: "var(--color-text)" }}>{value}</h2>
          <p style={{ margin: 0, color: "var(--color-text-muted)", fontSize: "13px" }}>{description}</p>
        </div>
        <div style={{ fontSize: "28px", background: "var(--color-surface-muted)", borderRadius: "var(--radius-md)", padding: "10px" }}>{icon}</div>
      </div>
    </div>
  );
}

function StatusRow({ label, value, percentage }: { label: string; value: number; percentage: number }) {
  return (
    <div style={{ marginBottom: "20px" }}>
      <div style={{ display: "flex", justifyContent: "space-between", marginBottom: "8px" }}>
        <span style={{ color: "var(--color-text)", fontSize: "14px" }}>{label}</span>
        <span style={{ color: "var(--color-text-secondary)", fontSize: "14px" }}>{value}</span>
      </div>
      <div style={{ height: "8px", background: "var(--color-surface-muted)", borderRadius: "10px", overflow: "hidden" }}>
        <div style={{ width: `${percentage}%`, height: "100%", background: "var(--color-ink)", borderRadius: "10px" }} />
      </div>
    </div>
  );
}

const primaryButton: React.CSSProperties = { width: "100%", padding: "12px 16px", marginBottom: "12px", border: "none", borderRadius: "var(--radius-md)", background: "var(--color-ink)", color: "var(--color-paper)", cursor: "pointer", fontWeight: 600 };
const secondaryButton: React.CSSProperties = { padding: "9px 14px", border: "1px solid var(--color-border)", borderRadius: "var(--radius-md)", background: "var(--color-surface)", color: "var(--color-text)", cursor: "pointer" };
const actionButton: React.CSSProperties = { width: "100%", padding: "12px 16px", marginBottom: "12px", border: "1px solid var(--color-border)", borderRadius: "var(--radius-md)", background: "var(--color-surface)", color: "var(--color-text)", cursor: "pointer", fontWeight: 500 };