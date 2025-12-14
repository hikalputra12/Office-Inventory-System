package repository

import (
	"context"
	"fmt"
	"inventory-system/model"
	"time"

	"github.com/jackc/pgx/v5"
)

// RepositoryCategory handles database operations related to categories.
type RepositoryCategory struct {
	DB *pgx.Conn
}

// RepositoryCategoryInterface defines the methods for category repository.
type RepositoryCategoryInterface interface {
	GetAllCategory() ([]*model.Category, error)
	CreateCategory(category *model.Category) error
	// GetCategoryByID(id int) (*model.Category, error)
	// UpdateCategory(id int, category *model.Category) (*model.Category, error)
	// DeleteCategory(id int) error
}

// NewRepositoryCategory creates a new instance of RepositoryCategory.
func NewRepositoryCategory(db *pgx.Conn) RepositoryCategory {
	return RepositoryCategory{
		DB: db,
	}
}

// GetAllCategory retrieves all categories from the database.
func (r RepositoryCategory) GetAllCategory() ([]*model.Category, error) {
	query := `SELECT category_id, name, description FROM category
				WHERE deleted_at IS NULL;`
	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var categories []*model.Category

	for rows.Next() {
		var t model.Category
		err := rows.Scan(&t.ID, &t.Name, &t.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &t)
	}
	return categories, nil
}

func (r RepositoryCategory) CreateCategory(category *model.Category) error {
	query := `INSERT INTO category (name,description,created_at, updated_at) VALUES
	($1, $2,$3,$4) RETURNING category_id , created_at, updated_at`

	now := time.Now()
	row := r.DB.QueryRow(context.Background(), query, category.Name, category.Description, now, now)

	category.CreatedAt = now
	category.UpdatedAt = now
	// .Scan adalah fungsi yang menangkap error pelaksanaan query dan error type mismatch.
	err := row.Scan(&category.ID, &category.CreatedAt, &category.UpdatedAt)

	// 4. Periksa Error dari .Scan()
	if err != nil {
		// Jika ada error (koneksi, constraint, duplikat key), kembalikan error tersebut.
		// Gunakan fmt.Errorf untuk memberikan konteks error.
		return fmt.Errorf("repository: gagal membuat kategori: %w", err)
	}
	return nil

}
