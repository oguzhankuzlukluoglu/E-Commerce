package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/oguzhan/e-commerce/internal/user/domain"
	"golang.org/x/crypto/bcrypt"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *domain.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO users (email, password, name, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`

	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	return r.db.QueryRow(
		query,
		user.Email,
		hashedPassword,
		user.Name,
		user.Role,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&user.ID)
}

func (r *userRepository) GetByID(id uint) (*domain.User, error) {
	user := &domain.User{}
	query := `SELECT id, email, password, name, role, created_at, updated_at FROM users WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}

	return user, err
}

func (r *userRepository) GetByEmail(email string) (*domain.User, error) {
	user := &domain.User{}
	query := `SELECT id, email, password, name, role, created_at, updated_at FROM users WHERE email = $1`

	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}

	return user, err
}

func (r *userRepository) Update(user *domain.User) error {
	query := `
		UPDATE users 
		SET email = $1, name = $2, role = $3, updated_at = $4
		WHERE id = $5`

	user.UpdatedAt = time.Now()

	_, err := r.db.Exec(query,
		user.Email,
		user.Name,
		user.Role,
		user.UpdatedAt,
		user.ID,
	)

	return err
}

func (r *userRepository) Delete(id uint) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
