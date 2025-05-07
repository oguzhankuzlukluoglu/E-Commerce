package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/oguzhan/e-commerce/internal/product/domain"
)

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) domain.ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product *domain.Product) error {
	query := `
		INSERT INTO products (name, description, price, stock, category, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`

	now := time.Now()
	product.CreatedAt = now
	product.UpdatedAt = now

	return r.db.QueryRow(
		query,
		product.Name,
		product.Description,
		product.Price,
		product.Stock,
		product.Category,
		product.CreatedAt,
		product.UpdatedAt,
	).Scan(&product.ID)
}

func (r *productRepository) GetByID(id uint) (*domain.Product, error) {
	product := &domain.Product{}
	query := `
		SELECT id, name, description, price, stock, category, created_at, updated_at 
		FROM products 
		WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Stock,
		&product.Category,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("product not found")
	}

	return product, err
}

func (r *productRepository) GetAll() ([]*domain.Product, error) {
	query := `
		SELECT id, name, description, price, stock, category, created_at, updated_at 
		FROM products`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*domain.Product
	for rows.Next() {
		product := &domain.Product{}
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Stock,
			&product.Category,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *productRepository) Update(product *domain.Product) error {
	query := `
		UPDATE products 
		SET name = $1, description = $2, price = $3, stock = $4, category = $5, updated_at = $6
		WHERE id = $7`

	product.UpdatedAt = time.Now()

	_, err := r.db.Exec(query,
		product.Name,
		product.Description,
		product.Price,
		product.Stock,
		product.Category,
		product.UpdatedAt,
		product.ID,
	)

	return err
}

func (r *productRepository) Delete(id uint) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *productRepository) UpdateStock(id uint, quantity int) error {
	query := `
		UPDATE products 
		SET stock = stock + $1, updated_at = $2
		WHERE id = $3`

	_, err := r.db.Exec(query, quantity, time.Now(), id)
	return err
}
