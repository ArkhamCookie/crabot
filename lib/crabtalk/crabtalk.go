package crabtalk

import (
	"fmt"

	"github.com/gSpera/morse"
)

func Talk() {
	// Place holder text
	text := "Hello World"

	// Convert to morse code
	textInMorse := morse.ToMorse(text)
	fmt.Println(textInMorse)
}
