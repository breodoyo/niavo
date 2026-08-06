import { useState, useEffect } from "react";
import { listUsers } from "../services/users";
import type { User } from "../services/auth";
import "./Organizations.css";

export default function Users() {
  const [users, setUsers] = useState<User[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [loadError, setLoadError] = useState<string | null>(null);

  useEffect(() => {
    loadUsers();
  }, []);

  async function loadUsers() {
    setIsLoading(true);
    setLoadError(null);
    try {
      const data = await listUsers();
      setUsers(data);
    } catch (err) {
      setLoadError("Couldn't load users.");
    } finally {
      setIsLoading(false);
    }
  }

  return (
    <div className="org-page">
      <h1>Users</h1>

      {isLoading && <p>Loading…</p>}
      {loadError && <div className="org-error">{loadError}</div>}

      {!isLoading && !loadError && (
        <table className="org-table">
          <thead>
            <tr>
              <th>Name</th>
              <th>Email</th>
              <th>Joined</th>
            </tr>
          </thead>
          <tbody>
            {users.map((u) => (
              <tr key={u.id}>
                <td>{u.first_name} {u.last_name}</td>
                <td>{u.email}</td>
                <td>{new Date(u.created_at).toLocaleDateString()}</td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
}