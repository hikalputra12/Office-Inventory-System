package cmd

import (
	"fmt"
	"inventory-system/handler"
)

// GetAllCategory retrieves and displays all categories using the provided handler.
func GetAllCategory(handler handler.HandlerCategory) {
	categories, err := handler.GetAllCategory()
	if err != nil {
		fmt.Println("Error fetching categories:", err)
		return
	}
	fmt.Println("\nList of Categories:")
	for _, category := range categories {
		fmt.Printf("ID: %d, Name: %s, Description: %s\n", category.ID, category.Name, category.Description)
	}
}
