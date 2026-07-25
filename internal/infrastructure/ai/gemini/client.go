// Package gemini provides an image-generation client for the Gemini API.
package gemini

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"emedic-bk/internal/domain/service"
)

const baseURL = "https://generativelanguage.googleapis.com/v1beta"

// Config holds Gemini API configuration.
type Config struct {
	APIKey string
	Model  string
}

// ImageService implements service.ImageGenerator using the Gemini Interactions API.
type ImageService struct {
	apiKey string
	model  string
	client *http.Client
}

// NewImageService creates a new Gemini-backed image generator.
func NewImageService(cfg Config) service.ImageGenerator {
	model := cfg.Model
	if model == "" {
		model = "gemini-3.1-flash-image"
	}
	return &ImageService{
		apiKey: cfg.APIKey,
		model:  model,
		client: &http.Client{Timeout: 60 * time.Second},
	}
}

type interactionRequest struct {
	Model          string         `json:"model"`
	Input          []inputPart    `json:"input"`
	ResponseFormat responseFormat `json:"response_format"`
}

type inputPart struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type responseFormat struct {
	Type        string `json:"type"`
	MimeType    string `json:"mime_type"`
	AspectRatio string `json:"aspect_ratio"`
	ImageSize   string `json:"image_size"`
}

type interactionResponse struct {
	Steps []struct {
		Type    string `json:"type"`
		Content []struct {
			Type     string `json:"type"`
			Data     string `json:"data"`
			MimeType string `json:"mime_type"`
		} `json:"content"`
	} `json:"steps"`
}

// GenerateImage requests a single image from the given text prompt.
func (s *ImageService) GenerateImage(ctx context.Context, prompt string) ([]byte, string, error) {
	body, err := json.Marshal(interactionRequest{
		Model: s.model,
		Input: []inputPart{{Type: "text", Text: prompt}},
		ResponseFormat: responseFormat{
			Type:        "image",
			MimeType:    "image/jpeg",
			AspectRatio: "16:9",
			ImageSize:   "1K",
		},
	})
	if err != nil {
		return nil, "", err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, baseURL+"/interactions", bytes.NewReader(body))
	if err != nil {
		return nil, "", err
	}
	req.Header.Set("x-goog-api-key", s.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(io.LimitReader(resp.Body, 20<<20)) // images are base64'd, allow headroom
	if err != nil {
		return nil, "", err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("gemini image generation failed: %s: %s", resp.Status, string(data))
	}

	var res interactionResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, "", err
	}
	for _, step := range res.Steps {
		for _, c := range step.Content {
			if c.Type != "image" || c.Data == "" {
				continue
			}
			imgBytes, err := base64.StdEncoding.DecodeString(c.Data)
			if err != nil {
				return nil, "", err
			}
			mimeType := c.MimeType
			if mimeType == "" {
				mimeType = "image/png"
			}
			return imgBytes, mimeType, nil
		}
	}
	return nil, "", fmt.Errorf("gemini response contained no image")
}
