Since Go is a statically typed language and does not have a direct equivalent to React, it's not possible to directly convert this React code into Go. However, I can provide you with a basic structure of a Go program that you can use as a starting point to implement similar functionality.

```go
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type ChatRole string

const (
	Assistant ChatRole = "Assistant"
	User      ChatRole = "User"
)

type ChatMessage struct {
	Content string    `json:"content"`
	Role    ChatRole  `json:"role"`
}

type ChatConfig struct {
	Stream bool `json:"stream"`
}

type ChatGPTProps struct {
	Prompts    []ChatMessage `json:"prompts"`
	Config     ChatConfig    `json:"config"`
	FetchPath  string        `json:"fetchPath"`
}

func main() {
	// Initialize your ChatGPTProps and other variables here
	chatGPTProps := ChatGPTProps{
		Prompts:   []ChatMessage{},
		Config:    ChatConfig{Stream: false},
		FetchPath: "https://api.example.com/chat",
	}

	// Implement your main logic here
}

func requestMessage(url string, messages []ChatMessage, prompts []ChatMessage, config ChatConfig) (*http.Response, error) {
	requestBody, err := json.Marshal(map[string]interface{}{
		"messages": messages,
		"prompts":  prompts,
		"config":   config,
	})

	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func fetchMessage(chatGPTProps ChatGPTProps, messages []ChatMessage) (string, error) {
	response, err := requestMessage(chatGPTProps.FetchPath, messages, chatGPTProps.Prompts, chatGPTProps.Config)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if chatGPTProps.Config.Stream {
		// Implement streaming logic here
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return "", err
		}

		var result map[string]interface{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			return "", err
		}

		message, ok := result["message"].(string)
		if !ok {
			return "", errors.New("message not found in response")
		}

		return message, nil
	}

	return "", nil
}

func onSend(chatGPTProps ChatGPTProps, message ChatMessage, messages []ChatMessage) ([]ChatMessage, error) {
	newMessages := append(messages, message)
	fetchedMessage, err := fetchMessage(chatGPTProps, newMessages)
	if err != nil {
		return nil, err
	}

	newMessages = append(newMessages, ChatMessage{
		Content: fetchedMessage,
		Role:    Assistant,
	})

	return newMessages, nil
}

func onClear() []ChatMessage {
	return []ChatMessage{}
}
```

This Go code provides a basic structure for implementing a chatbot similar to the given React code. You can further modify and expand this code to suit your specific requirements.