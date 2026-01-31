package main

import (
	"fmt"
	"time"
)

// ============ Mediator Interface ============

// ChatRoomMediator defines the interface for communication
type ChatRoomMediator interface {
	SendMessage(message string, user User)
	AddUser(user User)
}

// ============ Concrete Mediator ============

// ChatRoom is the concrete mediator
type ChatRoom struct {
	users []User
}

func NewChatRoom() *ChatRoom {
	return &ChatRoom{
		users: make([]User, 0),
	}
}

func (c *ChatRoom) AddUser(user User) {
	c.users = append(c.users, user)
	fmt.Printf("📢 [System] %s joined the chat room\n", user.GetName())
}

func (c *ChatRoom) SendMessage(message string, sender User) {
	timestamp := time.Now().Format("15:04:05")

	// Send message to all users except the sender
	for _, user := range c.users {
		if user.GetName() != sender.GetName() {
			user.ReceiveMessage(message, sender.GetName(), timestamp)
		}
	}
}

// ============ Colleague Interface ============

// User is the colleague interface
type User interface {
	SendMessage(message string)
	ReceiveMessage(message string, from string, timestamp string)
	GetName() string
}

// ============ Concrete Colleagues ============

// ChatUser is a concrete colleague
type ChatUser struct {
	name     string
	mediator ChatRoomMediator
}

func NewChatUser(name string, mediator ChatRoomMediator) *ChatUser {
	user := &ChatUser{
		name:     name,
		mediator: mediator,
	}
	mediator.AddUser(user)
	return user
}

func (u *ChatUser) GetName() string {
	return u.name
}

func (u *ChatUser) SendMessage(message string) {
	fmt.Printf("💬 [%s]: %s\n", u.name, message)
	u.mediator.SendMessage(message, u)
}

func (u *ChatUser) ReceiveMessage(message string, from string, timestamp string) {
	fmt.Printf("📨 [%s] received from [%s] at %s: %s\n", u.name, from, timestamp, message)
}

func main() {
	fmt.Println("===== Mediator Pattern Example: Chat Room =====\n")

	// Create the mediator (chat room)
	chatRoom := NewChatRoom()

	// Create users (colleagues)
	alice := NewChatUser("Alice", chatRoom)
	bob := NewChatUser("Bob", chatRoom)
	charlie := NewChatUser("Charlie", chatRoom)
	diana := NewChatUser("Diana", chatRoom)

	fmt.Println()
	fmt.Println("--- Chat Conversation ---")

	// Users send messages through the mediator
	alice.SendMessage("Hi everyone! 👋")
	fmt.Println()

	bob.SendMessage("Hey Alice! How are you?")
	fmt.Println()

	charlie.SendMessage("Good morning all! ☀️")
	fmt.Println()

	diana.SendMessage("Welcome everyone! 🎉")
	fmt.Println()

	alice.SendMessage("I'm great, thanks Bob!")
	fmt.Println()

	fmt.Println("--- Benefits of Mediator Pattern ---")
	fmt.Println("✅ Users don't know about each other")
	fmt.Println("✅ All communication goes through ChatRoom")
	fmt.Println("✅ Easy to add new users without modifying existing ones")
	fmt.Println("✅ Centralized message routing logic")
}
}
