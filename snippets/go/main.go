// Install: go mod init app && go mod tidy
// Run:     go run main.go

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model     string    `json:"model"`
	Messages  []Message `json:"messages"`
	MaxTokens int       `json:"max_tokens"`
}

type ChatResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

func main() {
	apiKey := os.Getenv("AGENTROUTER_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: AGENTROUTER_API_KEY environment variable not set")
		os.Exit(1)
	}

	reqBody := ChatRequest{
		Model: "gpt-5.6-sol",
		Messages: []Message{
			{Role: "user", Content: "Hi, I'm a vibecoder."},
		},
		MaxTokens: 256,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", "https://agentrouter.org/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Originator", "codex_cli_rs")
	req.Header.Set("Version", "0.101.0")
	req.Header.Set("User-Agent", "codex_cli_rs/0.101.0 (Mac OS 26.0.1; arm64) Apple_Terminal/464")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("HTTP Error %d: %s\n", resp.StatusCode, string(body))
		return
	}

	var chatResp ChatResponse
	if err := json.Unmarshal(body, &chatResp); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	if len(chatResp.Choices) > 0 {
		fmt.Println(chatResp.Choices[0].Message.Content)
	}
}
