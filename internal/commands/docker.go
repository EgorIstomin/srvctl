package commands

import (
	"fmt"
	"time"

	"srvctl/internal/runner"
)

func HandleDocker(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: srvctl docker <ps>")
	}

	switch args[0] {
	case "ps":
		res, err := runner.Run(10*time.Second, "docker", "ps")
		if err != nil {
			return err
		}
		fmt.Print(res.Stdout)
		return nil

	default:
		return fmt.Errorf("unknown docker subcommand: %s", args[0])
	}
}
