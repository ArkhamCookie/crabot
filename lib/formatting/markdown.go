package formatting

// Format input into italic markdown style
func Italic(input string) string       { return "*" + input + "*" }

// Format input into Bold markdown style
func Bold(input string) string         { return "**" + input + "**" }

// Format input into Bold+Italic markdown style
func BoldItalic(input string) string   { return "***" + input + "***" }

// Format input into Strikethough markdown style
func Strikethough(input string) string { return "~~" + input + "~~" }
