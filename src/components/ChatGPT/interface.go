```go
package main

import (
	"github.com/golang/protobuf/ptypes/any"
)

type ChatRole int

const (
	Assistant ChatRole = iota
	User
	System
)

type ChatGPTVersion int

const (
	GPT_35_turbo ChatGPTVersion = iota
	GPT_4
	GPT_4_32K
)

type Prompt struct {
	Title   *string
	Content *string
}

type ChatGPTProps struct {
	Header          *any.Any
	FetchPath       string
	Config          *ChatConfig
	Prompts         []*ChatMessage
	OnMessages      func(messages []*ChatMessage)
	OnSettings      func()
	OnChangeVersion func(version ChatGPTVersion)
}

type ChatMessage struct {
	Content string
	Role    ChatRole
}

type ChatMessageItemProps struct {
	Message *ChatMessage
}

type SendBarProps struct {
	Loading   bool
	Disabled  bool
	InputRef  *any.Any
	OnSettings func()
	OnSend    func(message *ChatMessage)
	OnClear    func()
	OnStop     func()
}

type ShowProps struct {
	Loading  *bool
	Fallback *any.Any
	Children *any.Any
}

type ChatGPInstance struct {
	SetPrompt     func(prompt *ChatMessage)
	SetChatContent func(prompt *Prompt)
	SetMessages    func(messages []*ChatMessage)
	GetMessages    func() []*ChatMessage
	ScrollDown     func()
}

type ChatConfig struct {
	Model  *ChatGPTVersion
	Stream *bool
}
```
