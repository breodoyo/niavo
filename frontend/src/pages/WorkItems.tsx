import { useState, useEffect, type FormEvent } from "react";
import { useParams, Link } from "react-router-dom";
import {
  listWorkItems,
  createWorkItem,
  updateWorkItem,
  deleteWorkItem,
  type WorkItem,
} from "../services/workitems";
import "./Organizations.css";

const STATUSES = ["received", "in_progress", "done"];
const PRIORITIES = ["low", "medium", "high", "urgent"];

export default function WorkItems() {
  const { workflowId } = useParams<{ workflowId: string }>();
  const [items, setItems] = useState<WorkItem[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [loadError, setLoadError] = useState<string | null>(null);

  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const [priority, setPriority] = useState("medium");
  const [dueDate, setDueDate] = useState("");
  const [formError, setFormError] = useState<string | null>(null);
  const [isSubmitting, setIsSubmitting] = useState(false);

  useEffect(() => {
    if (workflowId) load();
  }, [workflowId]);

  async function load() {
    if (!workflowId) return;
    setIsLoading(true);
    setLoadError(null);
    try {
      const data = await listWorkItems(workflowId);
      setItems(data);
    } catch (err) {
      setLoadError("Couldn't load work items.");
    } finally {
      setIsLoading(false);
    }
  }

  async function handleCreate(e: FormEvent) {
    e.preventDefault();
    if (!workflowId) return;
    setFormError(null);
    setIsSubmitting(true);
    try {
      await createWorkItem({
        workflow_id: workflowId,
        title,
        description: description || undefined,
        priority,
        due_date: dueDate || undefined,
      });
      setTitle("");
      setDescription("");
      setPriority("medium");
      setDueDate("");
      await load();
    } catch (err: any) {
      setFormError(err?.response?.data?.message || "Couldn't create work item.");
    } finally {
      setIsSubmitting(false);
    }
  }

  async function handleStatusChange(item: WorkItem, newStatus: string) {
    try {
      await updateWorkItem(item.id, {
        title: item.title,
        description: item.description,
        status: newStatus,
        priority: item.priority,
        due_date: item.due_date,
      });
      await load();
    } catch (err) {
      alert("Couldn't update status.");
    }
  }

  async function handleDelete(id: string, itemTitle: string) {
    if (!window.confirm(`Delete "${itemTitle}"? This can't be undone.`)) return;
    try {
      await deleteWorkItem(id);
      await load();
    } catch (err) {
      alert("Couldn't delete work item.");
    }
  }

  return (
    <div className="org-page">
      <p><Link to="/workflows">← Back to Workflows</Link></p>
      <h1>Work Items</h1>

      <form className="org-create-form" onSubmit={handleCreate}>
        {formError && <div className="org-error">{formError}</div>}
        <input
          type="text"
          placeholder="Title"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          required
        />
        <input
          type="text"
          placeholder="Description (optional)"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
        />
        <select value={priority} onChange={(e) => setPriority(e.target.value)}>
          {PRIORITIES.map((p) => (
            <option key={p} value={p}>{p}</option>
          ))}
        </select>
        <input
          type="date"
          value={dueDate}
          onChange={(e) => setDueDate(e.target.value)}
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
              <th>Title</th>
              <th>Status</th>
              <th>Priority</th>
              <th>Due</th>
              <th></th>
            </tr>
          </thead>
          <tbody>
            {items.map((item) => (
              <tr key={item.id}>
                <td>{item.title}</td>
                <td>
                  <select
                    value={item.status}
                    onChange={(e) => handleStatusChange(item, e.target.value)}
                  >
                    {STATUSES.map((s) => (
                      <option key={s} value={s}>{s}</option>
                    ))}
                  </select>
                </td>
                <td>{item.priority}</td>
                <td>{item.due_date || "—"}</td>
                <td className="org-actions">
                  <button className="org-delete-btn" onClick={() => handleDelete(item.id, item.title)}>
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