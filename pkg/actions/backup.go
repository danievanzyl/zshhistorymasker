package actions

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/urfave/cli/v3"
)

func Backup(ctx context.Context, cmd *cli.Command) (context.Context, error) {
	fmt.Println("backing up current", cmd.String("orig-history"))

	history, err := os.Open(cmd.String("orig-history"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer history.Close()

	backupFile, err := os.Create(cmd.String("bak-history"))
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
	return ctx, nil
}
