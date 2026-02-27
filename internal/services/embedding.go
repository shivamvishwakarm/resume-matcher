package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	ollamaEmbedURL   = "http://localhost:11434/api/embed"
	ollamaEmbedModel = "nomic-embed-text"
)

type ollamaEmbedRequest struct {
	Model string `json:"model"`
	Input string `json:"input"`
}

type ollamaEmbedResponse struct {
	Model      string      `json:"model"`
	Embeddings [][]float32 `json:"embeddings"`
	Error      string      `json:"error,omitempty"`
}

// GenerateEmbedding calls Ollama's local embedding API to convert text into a vector.
// Requires Ollama running locally with the nomic-embed-text model pulled.
// Install: https://ollama.com  |  Pull model: ollama pull nomic-embed-text
func GenerateEmbedding(text string) ([]float32, error) {
	reqBody := ollamaEmbedRequest{
		Model: ollamaEmbedModel,
		Input: text,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := http.Post(ollamaEmbedURL, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to call Ollama API (is Ollama running?): %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var embResp ollamaEmbedResponse
	if err := json.Unmarshal(body, &embResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if embResp.Error != "" {
		return nil, fmt.Errorf("Ollama API error: %s", embResp.Error)
	}

	if len(embResp.Embeddings) == 0 || len(embResp.Embeddings[0]) == 0 {
		return nil, fmt.Errorf("no embedding data returned from Ollama")
	}

	return embResp.Embeddings[0], nil
}
