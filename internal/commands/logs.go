package commands

import (
	"fmt"
	"strconv"
	"time"

	"srvctl/internal/runner"
)

func HandleLogs(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: srvctl logs <service> [--lines N]")
	}

	service := args[0]
	lines := 100

	// Парсим очень просто: ищем --lines N
	if len(args) >= 3 && args[1] == "--lines" {
		n, err := strconv.Atoi(args[2])
		if err != nil || n <= 0 {
			return fmt.Errorf("invalid --lines value: %s", args[2])
		}
		lines = n
	}

	res, err := runner.Run(
		10*time.Second,
		"journalctl",
		"-u", service,
		"-n", strconv.Itoa(lines),
		"--no-pager",
	)
	if err != nil {
		return err
	}

	fmt.Print(res.Stdout)
	return nil
}
