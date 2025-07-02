package repository

import (
	"database/sql"
	"errors"

	"github.com/Shashanktriathi1703/student-api/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// This is used for a new user
func (r *UserRepository) Create(user *model.CreatedUserRequest) (*model.User, error) {
	query :=
		`INSERT INTO users (name, email) VALUES ($1, $2)
	RETURNING id, name, email, created_at`

	var createdUser model.User
	err := r.db.QueryRow(query, user.Name, user.Email).Scan(
		&createdUser.ID,
		&createdUser.Name,
		&createdUser.Email,
		&createdUser.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &createdUser, nil
}

// This is basically used for "find the user by id"
func (r *UserRepository) FindByID(id int) (*model.User, error) {
	query :=
		`SELECT id, name, email, creates_at FROM users WHERE id = $1`

	var user model.User
	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
	}
	return &user, nil
}

func (r *UserRepository) GetAll() ([]model.User, error) {
	query := `SELECT id, name, email, created_at From users`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1`

	// Exec executes a query without returning any rows.
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	// RowsAffected returns the number of rows affected by an update, insert, or delete. Not every database or database driver may support this.
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no user found with given id")
	}
	return nil
}
