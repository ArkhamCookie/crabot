package formatting

import (
	"errors"
)

// Discord's markdown syntax
var discordMarkdownData = map[string]string{
	"italic":     "*",
	"bold":       "**",
	"boldItalic": "***",

	"strikethough": "~~",
}

// Convert *test* into Discord markdown syntax
// based on style given
func DiscordMarkdown(text, style string) (string, error) {
	// Confirm that text is given
	if text == "" {
		return "", errors.New("no text given")
	}
	// Confirm that style type is given
	if style == "" {
		return "", errors.New("no style given")
	}
	// Confirm style type exists
	_, ok := discordMarkdownData[style]
	if !ok {
		return "", errors.New("given style type doesn't exist")
	}

	// Transform based on style
	output := discordMarkdownData[style] + text + discordMarkdownData[style]

	return output, nil
}
