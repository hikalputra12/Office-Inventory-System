package model

/* Category represents a category of inventory items. */
type Category struct {
	Model
	Name        string `json:"name"`
	Description string `json:"description"`
}
