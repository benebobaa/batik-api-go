package handlers

import (
	"github.com/benebobaa/batik-api-go/database"
	"github.com/benebobaa/batik-api-go/models"
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	sayHello := "Hello, World!"
	return c.Status(200).JSON(fiber.Map{
		"message": sayHello,
	})
}

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}

	database.DB.Db.Find(&facts)

	return c.Status(200).JSON(facts)
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)

	if err := c.BodyParser(&fact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad Request. Invalid request body",
		})
	}

	if fact.Question == "" || fact.Answer == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Question and answer cannot be empty",
		})
	}
	database.DB.Db.Create(&fact)

	return c.Status(201).JSON(fact)
}

func DetailFact(c *fiber.Ctx) error {
	fact := models.Fact{}

	id := c.Params("id")

	//if not found return 404 message not found
	err := database.DB.Db.First(&fact, id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Not Found",
		})
	}
	return c.Status(200).JSON(fact)
}

func UpdateFact(c *fiber.Ctx) error {
	fact := models.Fact{}

	id := c.Params("id")

	err := database.DB.Db.First(&fact, id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Not Found",
		})
	}

	if err := c.BodyParser(&fact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad Request. Invalid request body",
		})

	}

	if fact.Question == "" || fact.Answer == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Question and answer cannot be empty",
		})
	}

	database.DB.Db.Save(&fact)

	return c.Status(200).JSON(fact)
}

func DeleteFact(c *fiber.Ctx) error {
	fact := models.Fact{}
	id := c.Params("id")

	if err := database.DB.Db.First(&fact, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Not Found",
		})
	}

	database.DB.Db.Delete(&fact)

	return c.Status(200).JSON(fiber.Map{
		"message": "Fact successfully deleted",
	})

}
