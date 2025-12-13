package repository

import (
	"context"
	"inventory-system/model"

	"github.com/jackc/pgx/v5"
)

// RepositoryCategory handles database operations related to categories.
type RepositoryCategory struct {
	DB *pgx.Conn
}

// RepositoryCategoryInterface defines the methods for category repository.
type RepositoryCategoryInterface interface {
	GetAllCategory() ([]*model.Category, error)
	// CreateCategory(category *model.Category) (model.Category, error)
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
			WHERE updated_at IS NULL;`
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
