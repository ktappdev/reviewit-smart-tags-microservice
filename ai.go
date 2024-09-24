package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
		fmt.Println("Error: No API key found")
	}

	requestBody, _ := json.Marshal(map[string]interface{}{
		"model": "openai/gpt-3.5-turbo",
		"messages": []map[string]string{
			{"role": "user", "content": direction + prompt},
		},
	})

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	req.Header.Set("Authorization", apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	var aiResponse AiResponse
	var tagsResponse TagsResponse
	body, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal([]byte(body), &aiResponse)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
	}
	data := aiResponse.Choices[0].Message.Content
	_ = json.Unmarshal([]byte(data), &tagsResponse)
	return tagsResponse, nil
}
