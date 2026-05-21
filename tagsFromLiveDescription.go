package main

import (
	"encoding/json"
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

	rawContent, err := queryAI(direction + formData.Description)
	if err != nil {
		fmt.Printf("Error querying AI: %v\n", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to query AI",
		})
	}

	var tagsResponse TagsResponse
	if err := json.Unmarshal([]byte(rawContent), &tagsResponse); err != nil {
		fmt.Printf("Error unmarshalling AI response: %v\n", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to query AI",
		})
	}

	return c.JSON(tagsResponse)
}
