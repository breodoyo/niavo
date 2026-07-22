# Database Design

## Overview

Niavo uses PostgreSQL as the primary database.

The MVP focuses on the core workflow of:

- Organizations
- Users
- Workflows
- Work Items

The database is designed to be simple, flexible, and expandable.

---

# Entity Relationship

```
Organization
      |
      |
      +------ Users
      |
      +------ Workflows
                    |
                    |
                    +------ Work Items
```

---

# Tables

## Organizations

Represents a workspace or team using Niavo.

### Schema

| Column | Type | Description |
|---|---|---|
| id | UUID | Primary key |
| name | VARCHAR | Organization name |
| slug | VARCHAR | Unique identifier |
| created_at | TIMESTAMP | Creation time |
| updated_at | TIMESTAMP | Last update time |

---

## Users

Represents people who access the system.

### Schema

| Column | Type | Description |
|---|---|---|
| id | UUID | Primary key |
| organization_id | UUID | User organization |
| name | VARCHAR | Full name |
| email | VARCHAR | Login email |
| password_hash | TEXT | Encrypted password |
| created_at | TIMESTAMP | Creation time |
| updated_at | TIMESTAMP | Last update time |

---

## Workflows

Defines how work moves through different stages.

Example:

```
New → In Progress → Completed
```

### Schema

| Column | Type | Description |
|---|---|---|
| id | UUID | Primary key |
| organization_id | UUID | Owner organization |
| name | VARCHAR | Workflow name |
| description | TEXT | Workflow details |
| created_at | TIMESTAMP | Creation time |
| updated_at | TIMESTAMP | Last update time |

---

## Work Items

The main unit of work in Niavo.

Examples:

- Customer requests
- Internal tasks
- Support tickets
- Projects

### Schema

| Column | Type | Description |
|---|---|---|
| id | UUID | Primary key |
| workflow_id | UUID | Related workflow |
| title | VARCHAR | Work title |
| description | TEXT | Work details |
| status | VARCHAR | Current state |
| priority | VARCHAR | Importance level |
| creator_id | UUID | User who created it |
| assignee_id | UUID | Assigned user |
| due_date | DATE | Deadline |
| created_at | TIMESTAMP | Creation time |
| updated_at | TIMESTAMP | Last update time |

---

# Relationships

## Organization → Users

One organization can have many users.

```
Organization 1 ---- * Users
```

---

## Organization → Workflows

One organization can have many workflows.

```
Organization 1 ---- * Workflows
```

---

## Workflow → Work Items

One workflow can contain many work items.

```
Workflow 1 ---- * Work Items
```

---

## User → Work Items

A user can create and be assigned many work items.

```
User 1 ---- * Work Items
```

---

# MVP Exclusions

The following are intentionally not included:

- Membership table
- Comments
- File attachments
- Notifications
- Billing
- Automation
- AI features

These can be introduced in future versions.