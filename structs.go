package main

import "github.com/lib/pq"

// Response struct represents the overall structure of the response
type ReviewItResponse struct {
	Success bool `json:"success"`
	Status  int  `json:"status"`
	Data    Data `json:"data"`
}

// Data struct represents the data field in the response
type Data struct {
	ID           string   `json:"id"`
	Address      string   `json:"address"`
	CreatedDate  string   `json:"createdDate"`
	Description  string   `json:"description"`
	DisplayImage string   `json:"display_image"`
	Images       []string `json:"images"`
	Videos       []string `json:"videos"`
	Links        []string `json:"links"`
	Name         string   `json:"name"`
	Tags         []string `json:"tags"`
	OpeningHrs   string   `json:"openingHrs"`
	ClosingHrs   string   `json:"closingHrs"`
	Telephone    string   `json:"telephone"`
	Website      []string `json:"website"`
	Rating       int      `json:"rating"`
	HasOwner     bool     `json:"hasOwner"`
	OwnerID      string   `json:"ownerId"`
	CreatedByID  string   `json:"createdById"`
	IsDeleted    bool     `json:"isDeleted"`
	Email        *string  `json:"email"` // Use pointer to handle null values
	BusinessID   string   `json:"businessId"`
	Business     Business `json:"business"`
	Reviews      []Review `json:"reviews"`
}

// Business struct represents the business information
type Business struct {
	ID                 string  `json:"id"`
	OwnerID            string  `json:"ownerId"`
	SubscriptionStatus string  `json:"subscriptionStatus"`
	SubscriptionExpiry string  `json:"subscriptionExpiry"`
	CreatedDate        string  `json:"createdDate"`
	IsVerified         bool    `json:"isVerified"`
	OwnerName          *string `json:"ownerName"` // Use pointer to handle null values
}

// Review struct represents an individual review
type Review struct {
	ID             string   `json:"id"`
	Body           string   `json:"body"`
	CreatedDate    string   `json:"createdDate"`
	HelpfulVotes   int      `json:"helpfulVotes"`
	UnhelpfulVotes int      `json:"unhelpfulVotes"`
	Rating         int      `json:"rating"`
	Title          string   `json:"title"`
	ProductID      string   `json:"productId"`
	UserID         string   `json:"userId"`
	IsVerified     *bool    `json:"isVerified"` // Use pointer to handle null values
	VerifiedBy     *string  `json:"verifiedBy"` // Use pointer to handle null values
	IsPublic       bool     `json:"isPublic"`
	Images         []string `json:"images"`
	Videos         []string `json:"videos"`
	Links          []string `json:"links"`
	CreatedBy      string   `json:"createdBy"`
	IsDeleted      bool     `json:"isDeleted"`
	User           User     `json:"user"`
}

// User struct represents the user information for the review
type User struct {
	ID          string `json:"id"`
	UserName    string `json:"userName"`
	Avatar      string `json:"avatar"`
	CreatedDate string `json:"createdDate"`
	Email       string `json:"email"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	ClerkUserID string `json:"clerkUserId"`
	IsDeleted   bool   `json:"isDeleted"`
}

// Define the struct to match the JSON structure
type AiResponse struct {
	ID                string   `json:"id"`
	Provider          string   `json:"provider"`
	Model             string   `json:"model"`
	Object            string   `json:"object"`
	Created           int64    `json:"created"`
	Choices           []Choice `json:"choices"`
	SystemFingerprint *string  `json:"system_fingerprint,omitempty"`
	Usage             Usage    `json:"usage"`
}

type Choice struct {
	Logprobs     *interface{} `json:"logprobs"` // Use interface{} for nullable field
	FinishReason string       `json:"finish_reason"`
	Index        int          `json:"index"`
	Message      Message      `json:"message"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
	Refusal string `json:"refusal"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type ReviewTag struct {
	Tags pq.StringArray `json:"tags" db:"tags"`
}

type RequestBody struct {
	ReviewID    string  `json:"review_id"`
	Description *string `json:"description"`
}

type FormData struct {
	Description string `json:"description"`
}

type TagsResponse struct {
	Tags []string `json:"tags"`
}

type ReviewRequest struct {
	ID string `json:"id"`
}
