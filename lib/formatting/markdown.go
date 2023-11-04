package formatting

func surround(markdownData string, text string) string { return markdownData + text + markdownData }

// Surround Markdown
func Italic(input string) string       { return surround("*", input) }
func Bold(input string) string         { return surround("**", input) }
func BoldItalic(input string) string   { return surround("***", input) }
func Strikethough(input string) string { return surround("~~", input) }

func BlockQuote(input string) string { return "> " + input }
