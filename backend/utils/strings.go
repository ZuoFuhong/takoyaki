package utils

import "strings"

// Contains reports whether substr is within s.
func Contains(s string, sub string) bool {
	if strings.TrimSpace(sub) != "" {
		return strings.Contains(s, strings.TrimSpace(sub))
	}
	return true
}
