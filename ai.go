package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type ResponseMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type RequestBodyAi struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type ResponseBody struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func queryAi(prompt string) (TagsResponse, error) {
	url := "https://openrouter.ai/api/v1/chat/completions"
	apiKey := os.Getenv("OPEN_ROUTER_API_KEY")
	if apiKey == "" {
		return TagsResponse{}, fmt.Errorf("OPEN_ROUTER_API_KEY environment variable is not set")
	}

	// Models to try in order (primary first, then fallback)
	models := []string{
		"meta-llama/llama-3.2-1b-instruct",
		"google/gemma-3-4b-it",
	}

	var lastErr error
	
	for i, model := range models {
		response, err := tryModel(url, apiKey, model, prompt)
		if err == nil {
			return response, nil
		}
		
		lastErr = err
		fmt.Printf("Model %s failed: %v\n", model, err)
		
		// If this isn't the last model, continue to next
		if i < len(models)-1 {
			fmt.Printf("Trying fallback model: %s\n", models[i+1])
		}
	}
	
	return TagsResponse{}, fmt.Errorf("all models failed, last error: %w", lastErr)
}

func tryModel(url, apiKey, model, prompt string) (TagsResponse, error) {
	requestBody, err := json.Marshal(map[string]interface{}{
		"model": model,
		"messages": []map[string]string{
			{"role": "user", "content": direction + prompt},
		},
	})
	if err != nil {
		return TagsResponse{}, fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return TagsResponse{}, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return TagsResponse{}, fmt.Errorf("failed to make HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return TagsResponse{}, fmt.Errorf("API request failed with status: %d, response: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return TagsResponse{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var aiResponse AiResponse
	if err := json.Unmarshal(body, &aiResponse); err != nil {
		return TagsResponse{}, fmt.Errorf("failed to unmarshal AI response: %w", err)
	}

	if len(aiResponse.Choices) == 0 {
		return TagsResponse{}, fmt.Errorf("no choices returned from AI API")
	}

	var tagsResponse TagsResponse
	if err := json.Unmarshal([]byte(aiResponse.Choices[0].Message.Content), &tagsResponse); err != nil {
		return TagsResponse{}, fmt.Errorf("failed to unmarshal tags response: %w", err)
	}

	return tagsResponse, nil
}
