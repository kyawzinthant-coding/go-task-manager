package routes

import (
	"task-manager-api/internal/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app * fiber.App , db * gorm.DB) {
	app.Get("/tasks", func(c * fiber.Ctx) error {
		var tasks []models.Task
		db.Find(&tasks)
		return c.JSON(tasks);
	})

	app.Post("/tasks", func(c *fiber.Ctx)error {
		var task models.Task
		if err := c.BodyParser((&task));
			err != nil {
				return c.Status(400).JSON(
					fiber.Map{"error" : "Bad requests"})
			}
			db.Create(&task)

			return c.Status(201).JSON(task)
	})

	app.Get("/tasks", func(c *fiber.Ctx) error {
		var task []models.Task

		db.Find(&task)
		return c.JSON(task)
	})

	app.Get("/tasks/:id" , func(c * fiber.Ctx) error {
		id := c.Params("id")
		var task models.Task
		if err := db.First(&task,id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Task not found"})
		}

		return c.JSON(task)
	})

	app.Put("/tasks/:id", func(c *fiber.Ctx) error {
    	id := c.Params("id")
    	var task models.Task

    	if err := db.First(&task, id).Error; err != nil {
       		 return c.Status(404).JSON(fiber.Map{"error": "Task not found"})
   		 }

		if err := c.BodyParser(&task); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Bad request"})
		}

		db.Save(&task)
		return c.JSON(task)
	})

	app.Delete("/tasks/:id", func(c *fiber.Ctx) error {
        id := c.Params("id")
        var task models.Task
        if err := db.First(&task, id).Error; err != nil {
            return c.Status(404).JSON(fiber.Map{"error": "Task not found"})
        }
        db.Delete(&task)
        return c.Status(204).SendString("") 
    })
}