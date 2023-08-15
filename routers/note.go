package routers

import (
	"fiber-api/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitNoteRouter(app *fiber.App, db *gorm.DB) {
	router := app.Group("/task")

	router.Use(func(c *fiber.Ctx) error {
		c.Response().Header.Add("x-response-tag", "wope")
		return c.Next()
	})

	router.Get("/", func(c *fiber.Ctx) error {
		var tasks []models.Task

		result := db.WithContext(c.Context()).Find(&tasks)

		if result.Error != nil {
			return result.Error
		}

		return c.JSON(tasks)
	})

	router.Get("/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return err
		}

		task := new(models.Task)
		result := db.WithContext(c.Context()).First(&task, id)

		if result.Error != nil {
			return result.Error
		}

		return c.JSON(task)
	})

	router.Post("/", func(c *fiber.Ctx) error {
		task := new(models.Task)
		if err := c.BodyParser(task); err != nil {
			return err
		}

		result := db.WithContext(c.Context()).Create(&task)
		if result.Error != nil {
			return result.Error
		}

		return c.JSON(map[string]any{"affected": result.RowsAffected})

	})
	router.Patch("/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return err
		}
		temp := models.Task{Model: gorm.Model{ID: uint(id)}}

		body := new(models.Task)
		if err := c.BodyParser(body); err != nil {
			return err
		}

		result := db.WithContext(c.Context()).Model(&temp).Updates(body)
		if result.Error != nil {
			return result.Error
		}

		return c.JSON(map[string]any{"affected": result.RowsAffected})
	})

	router.Delete("/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return err
		}

		result := db.WithContext(c.Context()).Delete(&models.Task{}, id)

		if result.Error != nil {
			return result.Error
		}

		return c.JSON(map[string]any{"affected": result.RowsAffected})
	})
}
