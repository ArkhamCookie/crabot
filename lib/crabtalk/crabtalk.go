package crabtalk

import (
	"errors"
	"strings"

	"github.com/gSpera/morse"
)

var crabConverter = morse.NewConverter(
	crabMorse,

	morse.WithCharSeparator("_"),
	morse.WithWordSeparator(" / "),
	morse.WithLowercaseHandling(true),
	morse.WithHandler(morse.IgnoreHandler),
	morse.WithTrailingSeparator(false),
)

// Convert the Morse code to "crab"
func convert(text string) string {
	text = strings.ReplaceAll(text, ".", "click")
	text = strings.ReplaceAll(text, "-", "clack")
	text = strings.ReplaceAll(text, "_ _", " / ") // Temporary until fix word separators
	return text
}

// Get the *text* in "crab"
func Get(text string) (string, error) {
	// Confirm that text is given
	if text == "" {
		return "", errors.New("no text given")
	}

	// Convert to morse code
	text = crabConverter.ToMorse(text)

	// Convert to crab
	text = convert(text)

	return text, nil
}
