package util

import "strings"

// TrimQuotes removes single and double quotes from string
func TrimQuotes(strs []string) []string {
	for _, s := range strs {
		s = strings.Replace(s, "\"", "", -1)
		s = strings.Replace(s, "'", "", -1)
	}
	return strs
}
