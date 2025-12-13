package service

import (
	"inventory-system/model"
	"inventory-system/repository"
)

// ServiceCategory provides category-related services.
type ServiceCategory struct {
	repo repository.RepositoryCategoryInterface
}

// ServiceCategoryInterface defines the methods for category services.
type ServiceCategoryInterface interface {
	GetAllCategory() ([]*model.Category, error)
	// CreateCategory(category model.Category) (model.Category, error)
	// GetCategoryByID(id int) (model.Category, error)
	// UpdateCategory(id int, category model.Category) (model.Category, error)
	// DeleteCategory(id int) error
}

// NewServiceCategory creates a new instance of ServiceCategory.
func NewServiceCategory(repo repository.RepositoryCategoryInterface) ServiceCategory {
	return ServiceCategory{repo: repo}
}

// GetAllCategory retrieves all categories using the repository.
func (s *ServiceCategory) GetAllCategory() ([]*model.Category, error) {
	category, err := s.repo.GetAllCategory()
	if err != nil {
		return nil, err
	}
	return category, nil
}
