package util

import "strings"

// TrimQuotes removes single and double quotes from string
func TrimQuotes(s string) string {
	s = strings.Replace(s, "\"", "", -1)
	s = strings.Replace(s, "'", "", -1)
	return s
}
