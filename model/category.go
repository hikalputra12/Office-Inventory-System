package model

type Category struct {
	Model
	Name        string `json:"name"`
	Description string `json:"description"`
}
