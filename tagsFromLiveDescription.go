package main

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func getTagsFromDescription(c *fiber.Ctx) error {
	var formData FormData
	if err := c.BodyParser(&formData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if formData.Description == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Description is required"})
	}

	response, err := queryAi(formData.Description)
	if err != nil {
		fmt.Printf("Error querying AI: %v\n", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to query AI",
			"details": err.Error(),
		})
	}
	return c.JSON(response)
}
