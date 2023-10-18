package main

import (
	"fmt"

	"crabot/crabtalk"
)

var (
	message string
	err     error
)

// Standin (for now) main function
func main() {
	// Convert 'Hello World!' to morse and print it
	message, err = crabtalk.Get("Hello World!")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(message)
}
