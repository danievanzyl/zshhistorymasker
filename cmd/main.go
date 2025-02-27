package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var sensitivePatterns = []*regexp.Regexp{
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

func maskSensitiveInfo(text string) string {
	maskedText := text
	for _, pattern := range sensitivePatterns {
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

func main() {
	file, err := os.Open(os.Getenv("HOME") + "/.zsh_history")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var commands []string
	var currentCommand strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = maskSensitiveInfo(line)

		// If line starts with : and we have a previous command, append it
		if strings.HasPrefix(line, ":") {
			if currentCommand.Len() > 0 {
				commands = append(commands, strings.TrimSpace(currentCommand.String()))
				currentCommand.Reset()
			}
			currentCommand.WriteString(line)
		} else {
			// If there's already content, add a newline
			if currentCommand.Len() > 0 {
				currentCommand.WriteString("\n")
			}
			currentCommand.WriteString(line)
		}
	}

	// Don't forget to append the last command
	if currentCommand.Len() > 0 {
		commands = append(commands, strings.TrimSpace(currentCommand.String()))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	for _, cmd := range commands {
		fmt.Println("xxxx>  --- ", cmd)
	}
}
