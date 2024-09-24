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

	response, err := queryAi(formData.Description)
	if err != nil {
		fmt.Println("Error querying AI:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to query AI"})
	}
	return c.JSON(response)
}
