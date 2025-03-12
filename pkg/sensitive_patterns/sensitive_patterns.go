package sensitivepatterns

import (
	"fmt"
	"regexp"
	"strings"
)

var Patterns = []*regexp.Regexp{
	// Generic API key patterns (adjust these patterns based on your needs)
	regexp.MustCompile(`(?i)api[_-]key[=:]\s*['"]([^'"]+)['"]`),
	regexp.MustCompile(`(?i)api[_-]key[=:]\s*([^\s]+)`),
	// AWS Access Key (20 characters)
	regexp.MustCompile(`AKIA[0-9A-Z]{16}`),
	// Generic alphanum API keys
	regexp.MustCompile(`([a-zA-Z0-9]{32,})`),
	// Bearer tokens
	regexp.MustCompile(`Bearer\s+([a-zA-Z0-9._\-]+)`),
	// GitHub tokens
	regexp.MustCompile(`gh[ps]_[0-9a-zA-Z]{36}`),
}

func UpdatePatterns(patterns []string) {
	for _, pattern := range patterns {
		Patterns = append(Patterns, regexp.MustCompile(pattern))
	}

	for _, p := range Patterns {
		fmt.Println("using the following pattern:", p.String())
	}
}

func MaskSensitiveInfo(text string) string {
	maskedText := text
	for _, pattern := range Patterns {
		maskedText = pattern.ReplaceAllStringFunc(maskedText, func(match string) string {
			// Find the API key part in the match
			submatch := pattern.FindStringSubmatch(match)
			if len(submatch) > 1 {
				// Replace the API key with asterisks, keeping the first and last 4 chars
				key := submatch[1]
				if len(key) > 8 {
					masked := key[:4] + strings.Repeat("*", len(key)-8) + key[len(key)-4:]
					return strings.Replace(match, key, masked, 1)
				}
				return strings.Replace(match, key, strings.Repeat("*", len(key)), 1)
			}
			return strings.Repeat("*", len(match))
		})
	}
	return maskedText
}
