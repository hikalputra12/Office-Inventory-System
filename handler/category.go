package handler

import (
	"inventory-system/model"
	"inventory-system/service"
)

// HandlerCategory provides category-related handlers.
type HandlerCategory struct {
	Service service.ServiceCategoryInterface
}

// HandlerCategoryInterface defines the methods for category handlers.
type HandlerCategoryInterface interface {
	GetAllCategory() ([]*model.Category, error)
	// CreateCategory(category model.Category) (model.Category, error)
	// GetCategoryByID(id int) (model.Category, error)
	// UpdateCategory(id int, category model.Category) (model.Category, error)
	// DeleteCategory(id int) error
}

// NewHandlerCategory creates a new instance of HandlerCategory.
func NewHandlerCategory(service service.ServiceCategoryInterface) HandlerCategory {
	return HandlerCategory{Service: service}
}

// GetAllCategory retrieves all categories using the service.
func (h *HandlerCategory) GetAllCategory() ([]*model.Category, error) {
	category, err := h.Service.GetAllCategory()
	if err != nil {
		return nil, err
	}
	return category, nil
}
