package main

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func updateTagsWithAI(c *fiber.Ctx, description string, reviewTag *ReviewTag, reviewID string) error {
	fmt.Println("Updating tags with AI with description", description)
	aiResponse, err := queryAi(description) // Implement this function as needed
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "AI query failed"})
	}

	// Append AI-generated tags
	reviewTag.Tags = append(reviewTag.Tags, aiResponse.Tags...) // Assuming aiResponse.Tags is []string

	// Update the review tag with the new tags
	updateQuery := "UPDATE review_tags SET tags = $1 WHERE review_id = $2"
	if _, err := db.Exec(updateQuery, reviewTag.Tags, reviewID); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update review tags"})
	}

	return nil
}
