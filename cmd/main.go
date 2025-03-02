package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	sensitivepatterns "github.com/danievanzyl/zshhistorymasker/pkg/sensitive_patterns"
)

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
		line = sensitivepatterns.MaskSensitiveInfo(line)

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
