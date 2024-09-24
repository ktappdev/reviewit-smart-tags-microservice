package main

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func getTags(c *fiber.Ctx) error {
	fmt.Println("Getting tags")
	var requestBody RequestBody

	// Parse the request body
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Attempt to find existing review tags
	var reviewTag ReviewTag
	query := "SELECT tags FROM review_tags WHERE review_id = $1"
	if err := db.Get(&reviewTag, query, requestBody.ReviewID); err == nil {
		// If found, return the existing tags
		return c.JSON(fiber.Map{"reviewTag": reviewTag.Tags})
	}

	// If no tags found, check if the review description is provided
	if requestBody.Description == nil {
		return fetchAndSaveTags(c, requestBody.ReviewID)
	}

	// If no tags and no description provided, return an error
	return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "No data found inside get tags"})
}
