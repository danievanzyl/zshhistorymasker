package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	sensitivepatterns "github.com/danievanzyl/zshhistorymasker/pkg/sensitive_patterns"
)

var historyFileLocation = fmt.Sprintf("%s/.zsh_history", os.Getenv("HOME"))

var bakHistoryFileLocation = fmt.Sprintf("%s_bak", historyFileLocation)

func backup() {
	history, err := os.Open(historyFileLocation)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer history.Close()

	backupFile, err := os.Create(bakHistoryFileLocation)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating backup file: %v\n", err)
		os.Exit(1)
	}

	defer backupFile.Close()

	uu, err := io.Copy(backupFile, history)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error copying backup file: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("written bytes", uu)
}

func main() {

	fmt.Println("scanning and updating", historyFileLocation)
	backup()
	fmt.Println("backed up current .zsh_history (this will contain sensitive information on first run)")

	// var commands []string
	var currentCommand strings.Builder
	history, err := os.Open(bakHistoryFileLocation)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer history.Close()

	newHistory, err := os.Create(historyFileLocation)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating new history file: %v\n", err)
		os.Exit(1)
	}
	defer newHistory.Close()

	scanner := bufio.NewScanner(history)
	for scanner.Scan() {
		line := scanner.Text()

		line = sensitivepatterns.MaskSensitiveInfo(line)

		// If line starts with : and we have a previous command, append it
		if strings.HasPrefix(line, ":") {
			if currentCommand.Len() > 0 {
				// commands = append(commands, strings.TrimSpace(currentCommand.String()))
				newHistory.WriteString(fmt.Sprintf("%s\n", strings.TrimSpace(currentCommand.String())))
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
		newHistory.WriteString(strings.TrimSpace(currentCommand.String()))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
