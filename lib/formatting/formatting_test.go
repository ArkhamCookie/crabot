package formatting_test

import (
	"testing"

	"formatting/markdown"
)

var (
	input  string
	want   string
	output string
)

// Markdown pkg Testing Start

// Italic Test
func TestMarkdownItalic(t *testing.T) {
	input = "crabot/formatting italic test!"
	want = "*" + input + "*"

	output = markdown.Italic(input)
	if want != output {
		t.Fatalf(`markdown.Italic(%q) = %q, want match for %#q`, input, output, want)
	}
}

// Bold Test
func TestMarkdownBold(t *testing.T) {
	input = "crabot/formatting bold test!"
	want = "**" + input + "**"

	output = markdown.Bold(input)
	if want != output {
		t.Fatalf(`markdown.Bold(%q) = %q, want match for %#q`, input, output, want)
	}
}

// Bold+Italic Test
func TestMarkdownBoldItalic(t *testing.T) {
	input = "crabot/formatting bold italic test!"
	want = "***" + input + "***"

	output = markdown.BoldItalic(input)
	if want != output {
		t.Fatalf(`markdown.BoldItalic(%q) = %q, want match for %#q`, input, output, want)
	}
}

// Strikethough Test
func TestMarkdownStrikethough(t *testing.T) {
	input = "crabot/formatting strikethough test!"
	want = "~~" + input + "~~"

	output = markdown.Strikethough(input)
	if want != output {
		t.Fatalf(`markdown.Strikethough(%q) = %q, want match for %#q`, input, output, want)
	}
}

// Blockquote Test
func TestMarkdownBlockquote(t *testing.T) {
	input = "crabot/formatting foo test!"
	want = "> " + input

	output = markdown.Blockquote(input)
	if want != output {
		t.Fatalf(`markdown.Foo(%q) = %q, want match for %#q`, input, output, want)
	}
}

// Unordered List Test
func TestMarkdownUnorderedList(t *testing.T) {
	input = "crabot/formatting unordered list test!"
	want = "- crabot/formatting unordered list test!"

	output = markdown.UnorderedList(input, 0)
	if want != output {
		t.Fatalf(`markdown.UnorderedList(%q) = %q, want match for %#q`, input, output, want)
	}

	input = "1 indented item"
	want = "  - 1 indented item"
	output = markdown.UnorderedList(input, 1)
	if want != output {
		t.Fatalf(`markdown.UnorderedList(%q) = %q, want match for %#q`, input, output, want)
	}

	input = "2 indented item"
	want = "    - 2 indented item"
	output = markdown.UnorderedList(input, 2)
	if want != output {
		t.Fatalf(`markdown.UnorderedList(%q) = %q, want match for %#q`, input, output, want)
	}
}

// Markdown pkg Testing End
