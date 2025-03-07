package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

var version = ""
var historyFileLocation = fmt.Sprintf("%s/.zsh_history", os.Getenv("HOME"))

var bakHistoryFileLocation = fmt.Sprintf("%s_bak", historyFileLocation)

func main() {

	cmd := &cli.Command{
		Name:  "boom",
		Usage: "make an explosive entrance",
		Action: func(context.Context, *cli.Command) error {
			fmt.Println("boom! I say!")
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}
