package workflow

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
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateWorkflow(
	ctx context.Context,
	wf Workflow,
) (Workflow, error) {
	query := `
	INSERT INTO workflows(name, description)
	VALUES($1, $2)
	RETURNING id, name, description, created_at, updated_at
	`
	err := r.db.QueryRow(
		ctx,
		query,
		wf.Name,
		wf.Description,
	).Scan(
		&wf.ID,
		&wf.Name,
		&wf.Description,
		&wf.CreatedAt,
		&wf.UpdatedAt,
	)
	if err != nil {
		return Workflow{}, err
	}
	return wf, nil
}

func (r *Repository) ListWorkflows(
	ctx context.Context,
) ([]Workflow, error) {
	query := `
	SELECT id, name, description, created_at, updated_at
	FROM workflows
	WHERE deleted_at IS NULL
	ORDER BY created_at DESC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	workflows := []Workflow{}

	for rows.Next() {
		var wf Workflow
		err := rows.Scan(
			&wf.ID,
			&wf.Name,
			&wf.Description,
			&wf.CreatedAt,
			&wf.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		workflows = append(workflows, wf)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return workflows, nil
}

func (r *Repository) GetWorkflow(
	ctx context.Context,
	id string,
) (Workflow, error) {
	query := `
	SELECT id, name, description, created_at, updated_at
	FROM workflows
	WHERE id = $1 AND deleted_at IS NULL
	`
	var wf Workflow
	err := r.db.QueryRow(ctx, query, id).Scan(
		&wf.ID,
		&wf.Name,
		&wf.Description,
		&wf.CreatedAt,
		&wf.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Workflow{}, ErrWorkflowNotFound
		}
		return Workflow{}, err
	}
	return wf, nil
}

func (r *Repository) UpdateWorkflow(
	ctx context.Context,
	id string,
	name string,
	description *string,
) (Workflow, error) {
	query := `
	UPDATE workflows
	SET name = $1, description = $2, updated_at = now()
	WHERE id = $3 AND deleted_at IS NULL
	RETURNING id, name, description, created_at, updated_at
	`
	var wf Workflow
	err := r.db.QueryRow(ctx, query, name, description, id).Scan(
		&wf.ID,
		&wf.Name,
		&wf.Description,
		&wf.CreatedAt,
		&wf.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Workflow{}, ErrWorkflowNotFound
		}
		return Workflow{}, err
	}
	return wf, nil
}

func (r *Repository) DeleteWorkflow(
	ctx context.Context,
	id string,
) error {
	query := `
	UPDATE workflows
	SET deleted_at = now()
	WHERE id = $1 AND deleted_at IS NULL
	`
	tag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrWorkflowNotFound
	}
	return nil
}