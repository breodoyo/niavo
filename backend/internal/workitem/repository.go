package workitem

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateWorkItem(ctx context.Context, item WorkItem) (WorkItem, error) {
	query := `
	INSERT INTO workitems(workflow_id, title, description, priority, assigned_user_id, due_date)
	VALUES($1, $2, $3, $4, $5, $6)
	RETURNING id, workflow_id, title, description, status, priority, assigned_user_id, due_date, created_at, updated_at
	`
	err := r.db.QueryRow(ctx, query,
		item.WorkflowID, item.Title, item.Description, item.Priority, item.AssignedUserID, item.DueDate,
	).Scan(
		&item.ID, &item.WorkflowID, &item.Title, &item.Description,
		&item.Status, &item.Priority, &item.AssignedUserID, &item.DueDate,
		&item.CreatedAt, &item.UpdatedAt,
	)
	if err != nil {
		return WorkItem{}, err
	}
	return item, nil
}

func (r *Repository) ListWorkItems(ctx context.Context, workflowID string) ([]WorkItem, error) {
	query := `
	SELECT id, workflow_id, title, description, status, priority, assigned_user_id, due_date, created_at, updated_at
	FROM workitems
	WHERE workflow_id = $1 AND deleted_at IS NULL
	ORDER BY created_at DESC
	`
	rows, err := r.db.Query(ctx, query, workflowID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []WorkItem{}
	for rows.Next() {
		var item WorkItem
		err := rows.Scan(
			&item.ID, &item.WorkflowID, &item.Title, &item.Description,
			&item.Status, &item.Priority, &item.AssignedUserID, &item.DueDate,
			&item.CreatedAt, &item.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *Repository) GetWorkItem(ctx context.Context, id string) (WorkItem, error) {
	query := `
	SELECT id, workflow_id, title, description, status, priority, assigned_user_id, due_date, created_at, updated_at
	FROM workitems
	WHERE id = $1 AND deleted_at IS NULL
	`
	var item WorkItem
	err := r.db.QueryRow(ctx, query, id).Scan(
		&item.ID, &item.WorkflowID, &item.Title, &item.Description,
		&item.Status, &item.Priority, &item.AssignedUserID, &item.DueDate,
		&item.CreatedAt, &item.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return WorkItem{}, ErrWorkItemNotFound
		}
		return WorkItem{}, err
	}
	return item, nil
}

func (r *Repository) UpdateWorkItem(ctx context.Context, id string, item WorkItem) (WorkItem, error) {
	query := `
	UPDATE workitems
	SET title = $1, description = $2, status = $3, priority = $4,
	    assigned_user_id = $5, due_date = $6, updated_at = now()
	WHERE id = $7 AND deleted_at IS NULL
	RETURNING id, workflow_id, title, description, status, priority, assigned_user_id, due_date, created_at, updated_at
	`
	err := r.db.QueryRow(ctx, query,
		item.Title, item.Description, item.Status, item.Priority,
		item.AssignedUserID, item.DueDate, id,
	).Scan(
		&item.ID, &item.WorkflowID, &item.Title, &item.Description,
		&item.Status, &item.Priority, &item.AssignedUserID, &item.DueDate,
		&item.CreatedAt, &item.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return WorkItem{}, ErrWorkItemNotFound
		}
		return WorkItem{}, err
	}
	return item, nil
}

func (r *Repository) DeleteWorkItem(ctx context.Context, id string) error {
	query := `UPDATE workitems SET deleted_at = now() WHERE id = $1 AND deleted_at IS NULL`
	tag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrWorkItemNotFound
	}
	return nil
}