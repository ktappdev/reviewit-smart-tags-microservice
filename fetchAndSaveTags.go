package main

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

func fetchAndSaveTags(c *fiber.Ctx, reviewID string) error {
	fmt.Printf("Fetching tags for %s\n", reviewID)
	// Fetch review from another source
	reviewFromReviewIt, err := getReviewItem(reviewID)
	if err != nil || reviewFromReviewIt.Data.Tags == nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "Review not found, check the review ID"})
	}

	// Create a new ReviewTag with fetched tags
	reviewTag := ReviewTag{
		Tags: pq.StringArray(reviewFromReviewIt.Data.Tags), // Assuming reviewFromReviewIt.Data.Tags is []string
	}

	// Save the new review tag
	insertQuery := "INSERT INTO review_tags (review_id, tags) VALUES ($1, $2)"
	if _, err := db.Exec(insertQuery, reviewID, reviewTag.Tags); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save review tags"})
	}

	// If description is available, use it to query AI for new tags
	if reviewFromReviewIt.Data.Description != "" {
		if err := updateTagsWithAI(c, reviewFromReviewIt.Data.Description, &reviewTag, reviewID); err != nil {
			return err
		}
	}

	// Send the response with the tags
	return c.JSON(fiber.Map{"reviewTag": reviewTag.Tags})
}
