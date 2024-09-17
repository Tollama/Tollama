package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ollama/ollama/api"
)

type model struct {
	choices      []string
	cursor       int
	selected     string
	client       *api.Client
	httpClient   *http.Client
	chatMode     bool
	chatInput    string
	chatHistory  []string
	models       []api.ListModelResponse
	currentModel string
}

func initialModel() model {
	baseURL, _ := url.Parse("http://localhost:11434")
	return model{
		choices:      []string{"Chat", "List Models", "Quit"},
		client:       api.NewClient(baseURL, http.DefaultClient),
		httpClient:   http.DefaultClient,
		chatMode:     false,
		chatInput:    "",
		chatHistory:  []string{},
		models:       []api.ListModelResponse{},
		currentModel: "",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) listModels() tea.Msg {
	resp, err := m.client.List(context.Background())
	if err != nil {
		return errMsg{err}
	}
	return modelsMsg{resp.Models}
}

func (m model) downloadModel(modelName string) tea.Msg {
	// Simulate downloading the model
	// Call the appropriate API endpoint to download the model
	m.currentModel = modelName
	return nil
}

func (m model) sendMessageToAPI(message string) (string, error) {
	url := "http://localhost:11434/api/chat"
	payload := map[string]string{"message": message, "model": m.currentModel}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := m.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var response map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	return response["reply"], nil
}
