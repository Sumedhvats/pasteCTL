package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Sumedhvats/pastectl/internal/config"
)

var apiClient = &http.Client{
	Timeout: time.Second * 15,
}

type Paste struct {
	ID        string     `json:"id"`
	Content   string     `json:"content"`
	Language  string     `json:"language"`
	CreatedAt time.Time  `json:"created_at"`
	ExpireAt  *time.Time `json:"expire_at,omitempty"`
	Views     int        `json:"views"`
}

type CreatePasteRequest struct {
	Content  string `json:"content"`
	Language string `json:"language"`
	Expire   string `json:"expire"`
}

type UpdatePasteRequest struct {
	ID       string `json:"id"`
	Content  string `json:"content"`
	Language string `json:"language"`
}
func getAPIURL() (string, error) {
	url := config.Get("backend_url")
	if url == "" {


		return "", fmt.Errorf("backend_url is not set. Please use 'pastectl config set backend_url <url>'")
		
	}
	return url, nil
}

func CreatePaste(content, language, expiry string) (*Paste, error) {
	backendURL, err := getAPIURL()
	if err != nil {
		return nil, err
	}

	reqData := CreatePasteRequest{
		Content:  content,
		Language: language,
		Expire:   expiry,
	}
	jsonData, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := apiClient.Post(backendURL+"/api/pastes", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("api returned non-OK status: %s", resp.Status)
	}

	var paste Paste
	if err := json.NewDecoder(resp.Body).Decode(&paste); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return &paste, nil
}

func GetPaste(id string) (*Paste, error) {
	backendURL, err := getAPIURL()
	if err != nil {
		return nil, err
	}

	resp, err := apiClient.Get(fmt.Sprintf("%s/api/pastes/%s", backendURL, id))
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("api returned non-OK status: %s", resp.Status)
	}

	var paste Paste
	if err := json.NewDecoder(resp.Body).Decode(&paste); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return &paste, nil
}

func GetPasteRaw(id string) (string, error) {
	backendURL, err := getAPIURL()
	if err != nil {
		return "", err
	}

	resp, err := apiClient.Get(fmt.Sprintf("%s/api/pastes/%s/raw", backendURL, id))
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("api returned non-OK status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var content string
	if err := json.Unmarshal(body, &content); err != nil {
		return "", fmt.Errorf("failed to unmarshal raw content from JSON string: %w", err)
	}

	return content, nil
}

func UpdatePaste(id, content, language string) (*Paste, error) {
	backendURL, err := getAPIURL()
	if err != nil {
		return nil, err
	}

	reqData := UpdatePasteRequest{
		ID:       id,
		Content:  content,
		Language: language,
	}
	jsonData, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/api/pastes/%s", backendURL, id), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create PUT request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := apiClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("api returned non-OK status: %s", resp.Status)
	}

	var paste Paste
	if err := json.NewDecoder(resp.Body).Decode(&paste); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return &paste, nil
}