package main

import "fmt"

// Message represents a generic message
type Message struct {
	Content string
}

// GetMessage returns the content of the message
func (m Message) GetMessage() string {
	return m.Content
}

// SpecialMessage represents a special type of message
type SpecialMessage struct {
	Message  // Embedding Message
	Priority int
}

// GetMessage overrides the GetMessage method for SpecialMessage
func (s SpecialMessage) GetMessage() string {
	return "Special: " + s.Content
}

func main() {
	message := Message{Content: "Hello, World!"}
	specialMessage := SpecialMessage{
		Message:  Message{Content: "Important Update"},
		Priority: 1,
	}

	fmt.Println("Message:", message.GetMessage())                    // Output: Message: Hello, World!
	fmt.Println("Special Message:", specialMessage.GetMessage())     // Output: Special: Important Update
	fmt.Println("Embedded Message:", specialMessage.Message.Content) // Output: Embedded Message: Important Update
}
