package main

import (
	"task-manager-api/internal/models"
	"task-manager-api/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

    app := fiber.New();

    db, err := models.ConnectDB()

    if err != nil {
        panic("Failed to connect to the db")
    }

    routes.SetupRoutes(app,db)

   

    app.Listen(":3000")
}