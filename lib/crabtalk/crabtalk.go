package crabtalk

import (
	"errors"
	"strings"

	"github.com/gSpera/morse"
)

func convert(text string) string {
	text = strings.ReplaceAll(text, ".", "click")
	text = strings.ReplaceAll(text, "-", "clack")
	return text
}

func Talk(text string) (string, error) {
	// Confirm that text is given
	if text == "" {
		return "", errors.New("no text given")
	}

	// Convert to morse code
	text = morse.ToMorse(text)

	// Convert to crab
	text = convert(text)

	return text, nil
}
