package commands

import (
	"fmt"
	"strings"
	"time"

	"srvctl/internal/runner"
)

func HandleSys(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("sys requires subcommand: uptime|df|free|status")
	}

	switch args[0] {
	case "uptime":
		return printCmd("uptime")

	case "df":
		// Покажем корень
		return printCmd("df", "-h", "/")

	case "free":
		return printCmd("free", "-h")

	case "status":
		// Мини-статус: hostname + uptime + df + free
		if err := printCmd("hostname"); err != nil {
			return err
		}
		fmt.Println(strings.Repeat("-", 40))
		if err := printCmd("uptime"); err != nil {
			return err
		}
		fmt.Println(strings.Repeat("-", 40))
		if err := printCmd("df", "-h", "/"); err != nil {
			return err
		}
		fmt.Println(strings.Repeat("-", 40))
		if err := printCmd("free", "-h"); err != nil {
			return err
		}
		return nil

	default:
		return fmt.Errorf("unknown sys subcommand: %s", args[0])
	}
}

func printCmd(name string, args ...string) error {
	res, err := runner.Run(5*time.Second, name, args...)
	if err != nil {
		return err
	}
	if strings.TrimSpace(res.Stdout) != "" {
		fmt.Print(res.Stdout)
	}
	return nil
}
