package user

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

func (r *Repository) CreateUser(
	ctx context.Context,
	u User,
) (User, error) {
	query := `
	INSERT INTO users(email, first_name, last_name, password_hash)
	VALUES($1, $2, $3, $4)
	RETURNING id, email, first_name, last_name, password_hash, created_at, updated_at
	`
	err := r.db.QueryRow(
		ctx,
		query,
		u.Email,
		u.FirstName,
		u.LastName,
		u.PasswordHash,
	).Scan(
		&u.ID,
		&u.Email,
		&u.FirstName,
		&u.LastName,
		&u.PasswordHash,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return User{}, err
	}
	return u, nil
}

func (r *Repository) ListUsers(
	ctx context.Context,
) ([]User, error) {
	query := `
	SELECT id, email, first_name, last_name, password_hash, created_at, updated_at
	FROM users
	ORDER BY created_at DESC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var u User
		err := rows.Scan(
			&u.ID,
			&u.Email,
			&u.FirstName,
			&u.LastName,
			&u.PasswordHash,
			&u.CreatedAt,
			&u.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Repository) GetUser(
	ctx context.Context,
	id string,
) (User, error) {
	query := `
	SELECT id, email, first_name, last_name, password_hash, created_at, updated_at
	FROM users
	WHERE id = $1
	`
	var u User
	err := r.db.QueryRow(ctx, query, id).Scan(
		&u.ID,
		&u.Email,
		&u.FirstName,
		&u.LastName,
		&u.PasswordHash,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return User{}, err
	}
	return u, nil
}
func (r *Repository) GetUserByEmail(ctx context.Context, email string) (User, error) {
	query := `
	SELECT id, email, first_name, last_name, password_hash, created_at, updated_at
	FROM users
	WHERE email = $1
	`
	var u User
	err := r.db.QueryRow(ctx, query, email).Scan(
		&u.ID,
		&u.Email,
		&u.FirstName,
		&u.LastName,
		&u.PasswordHash,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return User{}, err
	}
	return u, nil
}
func (r *Repository) UpdateUser(
	ctx context.Context,
	id string,
	firstName string,
	lastName string,
) (User, error) {
	query := `
	UPDATE users
	SET first_name = $1, last_name = $2, updated_at = now()
	WHERE id = $3
	RETURNING id, email, first_name, last_name, password_hash, created_at, updated_at
	`
	var u User
	err := r.db.QueryRow(ctx, query, firstName, lastName, id).Scan(
		&u.ID,
		&u.Email,
		&u.FirstName,
		&u.LastName,
		&u.PasswordHash,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return User{}, err
	}
	return u, nil
}
func (r *Repository) DeleteUser(
	ctx context.Context,
	id string,
) error {
	query := `
	DELETE FROM users WHERE id = $1
	`
	tag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrUserNotFound
	}
	return nil
}
