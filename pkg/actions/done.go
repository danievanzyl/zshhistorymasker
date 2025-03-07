package actions

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func Done(ctx context.Context, cmd *cli.Command) error {

	fmt.Println("Remember to start a new terminal session!")
	return nil
}
