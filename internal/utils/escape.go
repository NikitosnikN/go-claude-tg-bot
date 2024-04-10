package utils

func EscapeMarkdown(text string) string {
	var escapedText string
	for _, char := range text {
		switch char {
		case '_', '*', '[', ']', '(', ')', '~', '`', '>', '#', '+', '-', '=', '|', '{', '}', '.', '!':
			escapedText += "\\" + string(char)
		default:
			escapedText += string(char)
		}
	}
	return escapedText
}
