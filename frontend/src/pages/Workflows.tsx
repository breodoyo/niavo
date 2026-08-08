import { useState, useEffect, type FormEvent } from "react";
import { Link } from "react-router-dom";
import {
  listWorkflows,
  createWorkflow,
  deleteWorkflow,
  type Workflow,
} from "../services/workflows";
import "./Organizations.css";

export default function Workflows() {
  const [workflows, setWorkflows] = useState<Workflow[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [loadError, setLoadError] = useState<string | null>(null);

  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [formError, setFormError] = useState<string | null>(null);
  const [isSubmitting, setIsSubmitting] = useState(false);

  useEffect(() => {
    load();
  }, []);

  async function load() {
    setIsLoading(true);
    setLoadError(null);
    try {
      const data = await listWorkflows();
      setWorkflows(data);
    } catch (err) {
      setLoadError("Couldn't load workflows.");
    } finally {
      setIsLoading(false);
    }
  }

  async function handleCreate(e: FormEvent) {
    e.preventDefault();
    setFormError(null);
    setIsSubmitting(true);
    try {
      await createWorkflow({ name, description: description || undefined });
      setName("");
      setDescription("");
      await load();
    } catch (err: any) {
      setFormError(err?.response?.data?.message || "Couldn't create workflow.");
    } finally {
      setIsSubmitting(false);
    }
  }

  async function handleDelete(id: string, wfName: string) {
    if (!window.confirm(`Delete "${wfName}"? This can't be undone.`)) return;
    try {
      await deleteWorkflow(id);
      await load();
    } catch (err) {
      alert("Couldn't delete workflow.");
    }
  }

  return (
    <div className="org-page">
      <h1>Workflows</h1>

      <form className="org-create-form" onSubmit={handleCreate}>
        {formError && <div className="org-error">{formError}</div>}
        <input
          type="text"
          placeholder="Name"
          value={name}
          onChange={(e) => setName(e.target.value)}
          required
        />
        <input
          type="text"
          placeholder="Description (optional)"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
        />
        <button type="submit" disabled={isSubmitting}>
          {isSubmitting ? "Creating…" : "Create"}
        </button>
      </form>

      {isLoading && <p>Loading…</p>}
      {loadError && <div className="org-error">{loadError}</div>}

      {!isLoading && !loadError && (
        <table className="org-table">
          <thead>
            <tr>
              <th>Name</th>
              <th>Description</th>
              <th>Created</th>
              <th></th>
            </tr>
          </thead>
          <tbody>
            {workflows.map((wf) => (
              <tr key={wf.id}>
                <td>
                  <Link to={`/workflows/${wf.id}/items`}>{wf.name}</Link>
                </td>
                <td>{wf.description || "—"}</td>
                <td>{new Date(wf.created_at).toLocaleDateString()}</td>
                <td className="org-actions">
                  <button className="org-delete-btn" onClick={() => handleDelete(wf.id, wf.name)}>
                    Delete
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
}