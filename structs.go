package main

// AI response structures
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

// API request/response structures
type FormData struct {
	Description string `json:"description"`
}

type TagsResponse struct {
	Tags []string `json:"tags"`
}

type TitleFormData struct {
	Body        string `json:"body"`
	Rating      int    `json:"rating"`
	ProductName string `json:"productName"`
}

type TitleResponse struct {
	Title string `json:"title"`
}
