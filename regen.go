package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

func regenerateTags(c *fiber.Ctx) error {
	var requestBody RequestBody
	// Parse the request body
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Get the existing review tag
	var reviewTag ReviewTag
	query := "SELECT tags FROM review_tags WHERE review_id = $1"
	err := db.Get(&reviewTag, query, requestBody.ReviewID)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "ReviewTag not found"})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	if requestBody.Description == nil {
		fmt.Println("Description is nil")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Description is needed"})
	}

	// Clear existing tags
	reviewTag.Tags = pq.StringArray{} // Reset the tags

	if requestBody.Description != nil {
		prompt := *requestBody.Description

		// Call your AI function to get new tags
		aiResponse, err := queryAi(prompt) // Implement this function as needed
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "AI query failed"})
		}

		// Add new tags to the reviewTag
		reviewTag.Tags = append(reviewTag.Tags, aiResponse.Tags...) // Assuming aiResponse.Tags is of type []string
	}

	// Update the review tag in the database
	updateQuery := "UPDATE review_tags SET tags = $1, updated_at = NOW() WHERE review_id = $2"
	_, err = db.Exec(updateQuery, reviewTag.Tags, requestBody.ReviewID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update review tags"})
	}

	return c.JSON(fiber.Map{"reviewTag": reviewTag.Tags})
}
