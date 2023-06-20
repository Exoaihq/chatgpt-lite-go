package main

import (
	"fmt"
)

type Message struct {
	Role    Role
	Content string
}

type ChatConfig struct {
	Model  *ChatGPTVersion
	Stream *bool
}

type ChatGPTVersion string

const (
	GPT_35_turbo ChatGPTVersion = "gpt-35-turbo"
	GPT_4        ChatGPTVersion = "gpt-4"
	GPT_4_32K    ChatGPTVersion = "gpt-4-32k"
)

type Role string

const (
	Assistant Role = "assistant"
	User      Role = "user"
)

func main() {
	fmt.Println("Go code generated")
}