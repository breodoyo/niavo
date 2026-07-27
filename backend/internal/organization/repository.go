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
