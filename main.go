package main

import (
	"fmt"

	"crabot/crabtalk"
)

func main() {
	// Convert message to morse and print it
	message, err := crabtalk.Talk("Hello World")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(message)
}
