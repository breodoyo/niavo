package organization

import (
	"context"

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
func (r *Repository) CreateOrganization(
	ctx context.Context,
	org Organization,
) (Organization, error) {

	query := `
	INSERT INTO organizations(name, slug)
	VALUES($1, $2)
	RETURNING id, name, slug, created_at
	`
	err := r.db.QueryRow(
		ctx,
		query,
		org.Name,
		org.Slug,
	).Scan(
		&org.ID,
		&org.Name,
		&org.Slug,
		&org.CreatedAt,
	)
	if err != nil {
		return Organization{}, err
	}
	return org, nil
}
func (r *Repository) ListOrganizations(
	ctx context.Context,
) ([]Organization, error) {
	query := `
	SELECT id, name, slug, created_at
	FROM organizations
	ORDER BY created_at DESC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	organizations := []Organization{}

	for rows.Next() {
		var org Organization

		err := rows.Scan(
			&org.ID,
			&org.Name,
			&org.Slug,
			&org.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		organizations = append(organizations, org)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return organizations, nil
}
func (r *Repository) GetOrganization(
	ctx context.Context,
	id string,
) (Organization, error) {
	query := `
	SELECT id, name, slug, created_at
    FROM organizations
	WHERE id = $1;
`
	var org Organization
	err := r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&org.ID,
		&org.Name,
		&org.Slug,
		&org.CreatedAt,
	)
	if err != nil {
		return Organization{}, err
	}
	return org, nil
}
func (r *Repository) UpdateOrganization(
	ctx context.Context,
	id string,
	name string,
) (Organization, error) {
	query := `
	UPDATE organizations
	SET name = $1, updated_at = now()
	WHERE id = $2
	RETURNING id, name, slug, created_at, updated_at
	`
	var org Organization
	err := r.db.QueryRow(ctx, query, name, id).Scan(
		&org.ID,
		&org.Name,
		&org.Slug,
		&org.CreatedAt,
		&org.UpdatedAt,
	)
	if err != nil {
		return Organization{}, err
	}
	return org, nil
}
func (r *Repository) DeleteOrganization(
	ctx context.Context,
	id string,
) error {
	query := `
	UPDATE organization
	SET deleted_at = now()
	WHERE id = $1 AND deleted_at IS NULL
	`
	tag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrOrganizationNotFound
	}
	return nil
}
