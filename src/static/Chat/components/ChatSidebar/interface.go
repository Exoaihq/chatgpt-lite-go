package main

import (
	"github.com/google/uuid"
)

type ChatRole string

const (
	Admin   ChatRole = "Admin"
	User    ChatRole = "User"
	AI      ChatRole = "AI"
)

type ChatMessage struct {
	Role     ChatRole
	Content  string
}

type Persona struct {
	Role    ChatRole
	Avatar  *string
	Name    *string
	Prompt  *string
}

type Chat struct {
	ID       string
	Persona  *Persona
	Messages []ChatMessage
}

type ChatSidebarProps struct {
	IsActive      *bool
	ChatList      []Chat
	CurrentChatId *string
	ChangeChat    func(chat Chat)
	CloseChat     func(chat Chat)
	NewChat       func(persona Persona)
	Settings      func()
}

func main() {
	// This is just a placeholder for the main function.
	// You can add your own logic here.
}