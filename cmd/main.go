package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/danievanzyl/zshhistorymasker/pkg/actions"
	"github.com/urfave/cli/v3"
)

var version = "v0.0.0"
var historyFileLocation = fmt.Sprintf("%s/.zsh_history", os.Getenv("HOME"))

var bakHistoryFileLocation = fmt.Sprintf("%s_bak", historyFileLocation)

func main() {

	cmd := &cli.Command{
		Version: version,
		Name:    "boom",
		Usage:   "make an explosive entrance",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "orig-history", Aliases: []string{"z"}, Value: historyFileLocation},
			&cli.StringFlag{Name: "bak-history", Aliases: []string{"Z"}, Value: bakHistoryFileLocation},
			&cli.StringSliceFlag{Name: "mask-pattern", Aliases: []string{"p"}, Value: []string{}},
			// &cli.StringSliceFlag{Name: "mask-exclude", Aliases: []string{"e"}, Value: []string{}},
		},
		Before: actions.Backup,
		Action: actions.Mask,
		After:  actions.Done,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}
