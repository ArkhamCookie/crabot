package crabtalk

import (
	"errors"
	"strings"

	"github.com/gSpera/morse"
)

var crabMorse = morse.NewConverter(
	morse.DefaultMorse,

	morse.WithCharSeparator("_"),
	morse.WithWordSeparator(" / "),
	morse.WithLowercaseHandling(true),
	morse.WithHandler(morse.IgnoreHandler),
	morse.WithTrailingSeparator(false),
)

func convert(text string) string {
	text = strings.ReplaceAll(text, ".", "click")
	text = strings.ReplaceAll(text, "-", "clack")
	return text
}

func Get(text string) (string, error) {
	// Confirm that text is given
	if text == "" {
		return "", errors.New("no text given")
	}

	// Convert to morse code
	text = crabMorse.ToMorse(text)

	// Convert to crab
	text = convert(text)

	return text, nil
}
