package formatting

// Surround Markdown
func Italic(input string) string       { return "*" + input + "*" }
func Bold(input string) string         { return "**" + input + "**" }
func BoldItalic(input string) string   { return "***" + input + "***" }
func Strikethough(input string) string { return "~~" + input + "~~" }

func BlockQuote(input string) string { return "> " + input }
