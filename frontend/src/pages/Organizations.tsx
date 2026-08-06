import { useState, useEffect, type FormEvent } from "react";
import {
  listOrganizations,
  createOrganization,
  updateOrganization,
  deleteOrganization,
  type Organization,
} from "../services/organizations";
import "./Organizations.css";

export default function Organizations() {
  const [organizations, setOrganizations] = useState<Organization[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [loadError, setLoadError] = useState<string | null>(null);

  const [name, setName] = useState("");
  const [slug, setSlug] = useState("");
  const [formError, setFormError] = useState<string | null>(null);
  const [isSubmitting, setIsSubmitting] = useState(false);

  const [editingId, setEditingId] = useState<string | null>(null);
  const [editName, setEditName] = useState("");
  const [editError, setEditError] = useState<string | null>(null);
  const [isSaving, setIsSaving] = useState(false);

  useEffect(() => {
    loadOrganizations();
  }, []);

  async function loadOrganizations() {
    setIsLoading(true);
    setLoadError(null);
    try {
      const data = await listOrganizations();
      setOrganizations(data);
    } catch (err) {
      setLoadError("Couldn't load organizations.");
    } finally {
      setIsLoading(false);
    }
  }

  async function handleCreate(e: FormEvent) {
    e.preventDefault();
    setFormError(null);
    setIsSubmitting(true);

    try {
      await createOrganization({ name, slug });
      setName("");
      setSlug("");
      await loadOrganizations();
    } catch (err: any) {
      const message =
        err?.response?.data?.message || "Couldn't create organization.";
      setFormError(message);
    } finally {
      setIsSubmitting(false);
    }
  }

  function startEditing(org: Organization) {
    setEditingId(org.id);
    setEditName(org.name);
    setEditError(null);
  }

  function cancelEditing() {
    setEditingId(null);
    setEditName("");
    setEditError(null);
  }

  async function handleSave(id: string) {
    setEditError(null);
    setIsSaving(true);
    try {
      await updateOrganization(id, { name: editName });
      setEditingId(null);
      await loadOrganizations();
    } catch (err: any) {
      const message =
        err?.response?.data?.message || "Couldn't update organization.";
      setEditError(message);
    } finally {
      setIsSaving(false);
    }
  }

  async function handleDelete(id: string, orgName: string) {
    const confirmed = window.confirm(`Delete "${orgName}"? This can't be undone.`);
    if (!confirmed) return;

    try {
      await deleteOrganization(id);
      await loadOrganizations();
    } catch (err) {
      alert("Couldn't delete organization.");
    }
  }

  return (
    <div className="org-page">
      <h1>Organizations</h1>

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
          placeholder="Slug"
          value={slug}
          onChange={(e) => setSlug(e.target.value)}
          required
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
              <th>Slug</th>
              <th>Created</th>
              <th></th>
            </tr>
          </thead>
          <tbody>
            {organizations.map((org) => (
              <tr key={org.id}>
                {editingId === org.id ? (
                  <>
                    <td colSpan={2}>
                      <input
                        type="text"
                        value={editName}
                        onChange={(e) => setEditName(e.target.value)}
                        autoFocus
                      />
                      {editError && <div className="org-error">{editError}</div>}
                    </td>
                    <td>{new Date(org.created_at).toLocaleDateString()}</td>
                    <td className="org-actions">
                      <button onClick={() => handleSave(org.id)} disabled={isSaving}>
                        {isSaving ? "Saving…" : "Save"}
                      </button>
                      <button onClick={cancelEditing} disabled={isSaving}>
                        Cancel
                      </button>
                    </td>
                  </>
                ) : (
                  <>
                    <td>{org.name}</td>
                    <td>{org.slug}</td>
                    <td>{new Date(org.created_at).toLocaleDateString()}</td>
                    <td className="org-actions">
                      <button onClick={() => startEditing(org)}>Edit</button>
                      <button
                        className="org-delete-btn"
                        onClick={() => handleDelete(org.id, org.name)}
                      >
                        Delete
                      </button>
                    </td>
                  </>
                )}
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
}