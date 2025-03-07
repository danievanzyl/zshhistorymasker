package actions

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	sensitivepatterns "github.com/danievanzyl/zshhistorymasker/pkg/sensitive_patterns"
	"github.com/urfave/cli/v3"
)

func Mask(ctx context.Context, cmd *cli.Command) error {

	var currentCommand strings.Builder
	history, err := os.Open(cmd.String("bak-history"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		return err
	}
	defer history.Close()

	newHistory, err := os.Create(cmd.String("orig-history"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating new history file: %v\n", err)
		return err
	}
	defer newHistory.Close()

	scanner := bufio.NewScanner(history)
	for scanner.Scan() {
		line := scanner.Text()

		line = sensitivepatterns.MaskSensitiveInfo(line)

		// If line starts with : and we have a previous command, append it
		if strings.HasPrefix(line, ":") {
			if currentCommand.Len() > 0 {
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
		return err
	}

	return nil
}
