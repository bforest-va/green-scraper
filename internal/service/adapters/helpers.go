package adapters

import (
	"html"
	"strconv"

	strip "github.com/grokify/html-strip-tags-go"
)

// convertChars swaps unix code values into standard characters
func convertChars(s string) string {
	return html.UnescapeString(s)
}

// stripHTML strips all HTML from a supplied string and returned the newly stripped string
func stripHTML(s string) string {
	return strip.StripTags(s)
}

// Atob converts a supplied ascii character (string) into a boolean with 0 being false and all other valid values being true
// If atob is given a non int convertable string it will be interpellated as a 0 value and false will be returned
func atob(s string) bool {
	i, _ := strconv.Atoi(s)
	return i != 0
}
