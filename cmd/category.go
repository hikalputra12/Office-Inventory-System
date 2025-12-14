package cmd

import (
	"context"
	"fmt"
	"inventory-system/handler"
	"inventory-system/model"
	"inventory-system/utils"
	"time"
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

// CreateCategory prompts the user for category details and creates a new category using the provided handler.
func CreateCategory(handler handler.HandlerCategory) {
	fmt.Println("--- Input Kategori Baru ---")
	nameInput := utils.ReadLine("Masukkan Nama Kategori: ")

	descInput := utils.ReadLine("Masukkan Deskripsi Kategori: ")

	// Buat struct baru dan ambil alamatnya (pointer)
	newCategory := &model.Category{
		Name:        nameInput,
		Description: descInput,
	}

	// Panggil Handler (memberikan pointer ke struct)
	// Ingat, handler.CreateCategory MENGUBAH struct ini (menambah ID/CreatedAt)
	err := handler.CreateCategory(newCategory)

	if err != nil {
		fmt.Println("❌ Error saat membuat kategori:", err)
		return
	}

	fmt.Println("\n✅ Kategori berhasil dibuat!")
	fmt.Printf("Detail Kategori Baru (termasuk ID dari DB): ID=%d, Nama=%s, Deskripsi=%s\n",
		newCategory.ID, newCategory.Name, newCategory.Description)
}

func GetCategoryByID(handler handler.HandlerCategory) {
	var id int
	fmt.Print("Enter Category ID: ")
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return
	}

	category, err := handler.GetCategoryByID(id)
	if err != nil {
		fmt.Println("Error fetching category:", err)
		return
	}

	fmt.Printf("Category Details - ID: %d, Name: %s, Description: %s\n", id, category.Name, category.Description)
}

func UpdateCategory(handler handler.HandlerCategory) {
	var id int
	fmt.Print("Enter Category ID to update: ")
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return
	}

	oldCategory, err := handler.GetCategoryByID(id)
	updatedCategory := &model.Category{
		Name:        oldCategory.Name,
		Description: oldCategory.Description,
	}
	if err != nil {
		fmt.Println("Error fetching category:", err)
		return
	}
	nameInput := utils.ReadLine("Enter new Category Name: ")
	descInput := utils.ReadLine("Enter new Category Description: ")
	if nameInput == "" {
		nameInput = oldCategory.Name
	}
	if descInput == "" {
		descInput = oldCategory.Description
	}
	updatedCategory = &model.Category{
		Name:        nameInput,
		Description: descInput,
	}

	_, err = handler.UpdateCategory(id, updatedCategory)
	if err != nil {
		fmt.Println("Error updating category:", err)
		return
	}

	fmt.Printf("Category updated successfully - ID: %d, Name: %s, Description: %s\n", id, updatedCategory.Name, updatedCategory.Description)
}

func DeleteCategory(handler handler.HandlerCategory) {
	var id int
	fmt.Print("Enter Category ID: ")
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return
	}
	// Menggunakan context dengan timeout (Rekomendasi)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = handler.DeleteCategory(ctx, id)
	if err != nil {
		fmt.Println("Error deleting category:", err)
		return
	}

	fmt.Printf("Success delete category - ID: %d\n", id)

}
