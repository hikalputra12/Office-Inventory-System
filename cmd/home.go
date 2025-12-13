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
		fmt.Println("3. Add Item")
		fmt.Println("4. View Items")
		fmt.Println("5. Exit")

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

		case 3:

		case 4:

		case 5:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
