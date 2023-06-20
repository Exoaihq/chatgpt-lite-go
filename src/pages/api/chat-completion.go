package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type ChatConfig struct {
	Model   string `json:"model"`
	Stream  bool   `json:"stream"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Usage struct {
		TotalTokens int `json:"total_tokens"`
	} `json:"usage"`
}

func main() {
	handler()
}

func handler() {
	reqBody := `{"messages": [{"role": "system", "content": "You are a helpful assistant."}, {"role": "user", "content": "Who won the world series in 2020?"}], "prompts": [], "config": {"model": "gpt-3.5-turbo", "stream": false}}`
	var requestData struct {
		Messages []Message   `json:"messages"`
		Prompts  []Message   `json:"prompts"`
		Config   ChatConfig  `json:"config"`
	}
	err := json.Unmarshal([]byte(reqBody), &requestData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	charLimit := 12000
	charCount := 0
	messagesToSend := []Message{}

	for _, message := range requestData.Messages {
		if charCount+len(message.Content) > charLimit {
			break
		}
		charCount += len(message.Content)
		messagesToSend = append(messagesToSend, message)
	}

	apiUrl := "https://api.openai.com/v1/chat/completions"
	apiKey := os.Getenv("OPENAI_API_KEY")
	model := requestData.Config.Model

	if requestData.Config.Stream == false {
		data, err := OpenAI(apiUrl, apiKey, model, messagesToSend, requestData.Prompts)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Response:", data)
	} else {
		fmt.Println("Streaming is not supported in this implementation.")
	}
}

func OpenAI(apiUrl, apiKey, model string, messages, prompts []Message) (string, error) {
	requestData := struct {
		Model            string    `json:"model"`
		FrequencyPenalty int       `json:"frequency_penalty"`
		MaxTokens        int       `json:"max_tokens"`
		Messages         []Message `json:"messages"`
		PresencePenalty  int       `json:"presence_penalty"`
		Stream           bool      `json:"stream"`
		Temperature      float64   `json:"temperature"`
		TopP             float64   `json:"top_p"`
	}{
		Model:            model,
		FrequencyPenalty: 0,
		MaxTokens:        4000,
		Messages:         append(prompts, messages...),
		PresencePenalty:  0,
		Stream:           false,
		Temperature:      0.7,
		TopP:             0.95,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("The OpenAI API has encountered an error with a status code of %d and message %s", resp.StatusCode, resp.Status))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var openAIResponse OpenAIResponse
	err = json.Unmarshal(body, &openAIResponse)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(openAIResponse.Choices[0].Message.Content), nil
}