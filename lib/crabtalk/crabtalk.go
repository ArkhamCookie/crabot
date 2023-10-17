package crabtalk

import (
	"errors"

	"github.com/gSpera/morse"
)

func Talk(text string) (string, error) {
	// Place holder text
	if text == "" {
		return "", errors.New("no text given")
	}

	// Convert to morse code
	textInMorse := morse.ToMorse(text)

	return textInMorse, nil
}
