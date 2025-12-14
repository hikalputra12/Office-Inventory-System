package cmd

import (
	"fmt"
	"inventory-system/handler"
	"inventory-system/model"
	"inventory-system/utils"
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
