package formatting_test

import (
	"testing"

	"crabot/formatting"
)

var (
	input  string
	want   string
	output string
)

// Italic Test
func TestMarkdownItalic(t *testing.T) {
	input = "crabot/formatting italic test!"
	want = "*" + input + "*"

	output = formatting.Italic(input)
	if want != output {
		t.Fatalf(`formatting.Italic(%q) = %q, want match for %#q`, input, output, want)
	}
}

// Bold Test
func TestMarkdownBold(t *testing.T) {
	input = "crabot/formatting bold test!"
	want = "**" + input + "**"

	output = formatting.Bold(input)
	if want != output {
		t.Fatalf(`formatting.Bold(%q) = %q, want match for %#q`, input, output, want)
	}
}

// Bold+Italic Test
func TestMarkdownBoldItalic(t *testing.T) {
	input = "crabot/formatting bold italic test!"
	want = "***" + input + "***"

	output = formatting.BoldItalic(input)
	if want != output {
		t.Fatalf(`formatting.BoldItalic(%q) = %q, want match for %#q`, input, output, want)
	}
}

// Strikethough Test
func TestMarkdownStrikethough(t *testing.T) {
	input = "crabot/formatting strikethough test!"
	want = "~~" + input + "~~"

	output = formatting.Strikethough(input)
	if want != output {
		t.Fatalf(`formatting.Strikethough(%q) = %q, want match for %#q`, input, output, want)
	}
}

// Blockquote Test
func TestMarkdownBlockquote(t *testing.T) {
	input = "crabot/formatting foo test!"
	want = "> " + input

	output = formatting.Blockquote(input)
	if want != output {
		t.Fatalf(`formatting.Foo(%q) = %q, want match for %#q`, input, output, want)
	}
}

// TestMarkdown Template
/*
func TestMarkdownFoo(t *testing.T) {
	input = "crabot/formatting foo test!"
	want = input

	output = formatting.Foo(input)
	if want != output {
		t.Fatalf(`formatting.Foo(%q) = %q, want match for %#q`, input, output, want)
	}
}
*/
