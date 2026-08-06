// Install: go mod init app && go mod tidy
// Run:     go run claude.go
// Model:   claude-opus-5 | claude-opus-4-8

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

type ClaudeRequest struct {
	Model     string    `json:"model"`
	MaxTokens int       `json:"max_tokens"`
	Messages  []Message `json:"messages"`
}

type ClaudeResponse struct {
	Content []struct {
		Text string `json:"text"`
	} `json:"content"`
}

func main() {
	apiKey := os.Getenv("AGENTROUTER_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: AGENTROUTER_API_KEY not set")
		os.Exit(1)
	}

	reqBody := ClaudeRequest{
		Model:     "claude-opus-5", // or "claude-opus-4-8"
		MaxTokens: 256,
		Messages:  []Message{{Role: "user", Content: "Hi, who are you?"}},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", "https://agentrouter.org/v1/messages", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("anthropic-version", "2023-06-01")
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

	var claudeResp ClaudeResponse
	if err := json.Unmarshal(body, &claudeResp); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	if len(claudeResp.Content) > 0 {
		fmt.Println(claudeResp.Content[0].Text)
	}
}
