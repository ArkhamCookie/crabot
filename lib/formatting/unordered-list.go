package formatting

// Add an item to an Unordered List
func UnorderedList(input string, indent int) string {
	indentSpacing := ""
	// Add 2 spaces per indent level
	for i := indent; i > 0; {
		indentSpacing = "  " + indentSpacing
		i--
	}
	return indentSpacing + "- " + input
}
