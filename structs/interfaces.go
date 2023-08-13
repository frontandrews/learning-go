package main

import "fmt"

// Message interface defines the GetMessage method
type Message interface {
	GetMessage() string
}

// SimpleMessage struct represents a simple message
type SimpleMessage struct {
	Content string
}

// GetMessage returns the content of the simple message
func (m SimpleMessage) GetMessage() string {
	return m.Content
}

// SpecialMessage struct represents a special message
type SpecialMessage struct {
	Content  string
	Priority int
}

// GetMessage returns the content of the special message
func (s SpecialMessage) GetMessage() string {
	return "Priority: " + s.Content
}

func main() {
	message1 := SimpleMessage{Content: "Hello, World!"}
	message2 := SpecialMessage{Content: "Important Update", Priority: 1}

	// Using the Message interface to create a slice of messages
	messages := []Message{message1, message2}

	for _, msg := range messages {
		fmt.Println(msg.GetMessage())
	}
}
