package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func generateTitle(c *fiber.Ctx) error {
	var formData TitleFormData
	if err := c.BodyParser(&formData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	body := strings.TrimSpace(formData.Body)
	if body == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Body is required"})
	}
	formData.Body = body

	// Rating is int — whole stars only. If half-star ratings are needed later,
	// change Rating to float64 in structs.go and update this comparison.
	if formData.Rating < 1 || formData.Rating > 5 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Rating must be between 1 and 5"})
	}

	prompt := directionTitle + "\n\nBody: " + formData.Body + "\nRating: " + fmt.Sprintf("%d", formData.Rating) + " stars"
	if trimmed := strings.TrimSpace(formData.ProductName); trimmed != "" {
		formData.ProductName = trimmed
		prompt += "\nProduct: " + formData.ProductName
	}

	rawContent, err := queryAI(prompt)
	if err != nil {
		fmt.Printf("Error querying AI for title: %v\n", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate title"})
	}

	var titleResponse TitleResponse
	if err := json.Unmarshal([]byte(rawContent), &titleResponse); err != nil {
		fmt.Printf("Error unmarshalling title response: %v\n", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate title"})
	}

	// Reject empty or whitespace-only titles from AI
	if strings.TrimSpace(titleResponse.Title) == "" {
		fmt.Printf("Error: AI returned empty title\n")
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate title"})
	}

	// Enforce max 80 characters (rune-aware for multi-byte characters)
	if runes := []rune(titleResponse.Title); len(runes) > 80 {
		fmt.Printf("Warning: AI title exceeded 80 characters, truncating\n")
		titleResponse.Title = string(runes[:80])
	}

	return c.JSON(titleResponse)
}
