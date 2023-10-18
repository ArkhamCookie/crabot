package formatting_test

import (
	"regexp"
	"testing"

	"crabot/formatting"
)

func TestDiscordMarkdown(t *testing.T) {
	input := "crabot/formatting test!"
	style := "italic"
	want := regexp.MustCompile(`\*` + input + `\*`)

	output, err := formatting.DiscordMarkdown(input, style)
	if !want.MatchString(output) || err != nil {
		t.Fatalf(`DiscordItalics("crabot/formatting test!", "italic") = %q, %v, want match for %#q, nil`, output, err, want)
	}
}
