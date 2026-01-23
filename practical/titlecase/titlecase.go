package titlecase

import "strings"

// TitleCase capitalizes the first character of each word.
func TitleCase(s string) string {
	// Split the string into words
	words := strings.Fields(s)

	// Iterate over each word
	for i, w := range words {
		// Capitalize the first "character" of each word and keep the rest unchanged
		words[i] = strings.ToUpper(w[:1]) + w[1:]
	}

	// Join the words back into a single string
	return strings.Join(words, " ")
}
