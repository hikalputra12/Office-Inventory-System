package main

import (
	"context"
	"inventory-system/cmd"
	"inventory-system/database"
	"inventory-system/handler"
	"inventory-system/repository"
	"inventory-system/service"
	"log"
)

// init function
func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background())
	repository := repository.NewRepositoryCategory(db)
	service := service.NewServiceCategory(&repository)
	handler := handler.NewHandlerCategory(&service)
	cmd.HomePage(handler)
}
