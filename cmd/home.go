package cmd

import (
	"fmt"
	"inventory-system/handler"
	"os"
)

// HomePage displays the main menu and handles user input.
func HomePage(handler handler.HandlerCategory) {
	for {
		fmt.Println("\n1. View Categories")
		fmt.Println("2. Add Category")
		fmt.Println("3. View Category by ID")
		fmt.Println("4. Update Category by ID")
		fmt.Println("5. Delete Category")
		fmt.Println("6. Exit")

		var choise int
		fmt.Print("enter your input: ")
		_, err := fmt.Scanln(&choise)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			return
		}
		ClearScreen()
		switch choise {
		case 1:
			GetAllCategory(handler)
		case 2:
			CreateCategory(handler)

		case 3:
			GetCategoryByID(handler)

		case 4:
			UpdateCategory(handler)
		case 5:
			DeleteCategory(handler)

		case 6:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
