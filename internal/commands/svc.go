package commands

import (
	"fmt"
	"strings"
	"time"

	"srvctl/internal/runner"
)

func HandleSvc(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: srvctl svc <status|restart> <service>")
	}

	action := args[0]
	service := args[1]

	switch action {
	case "status":
		// Быстро и читаемо
		res, err := runner.Run(5*time.Second, "systemctl", "is-active", service)
		if err != nil {
			return err
		}
		state := strings.TrimSpace(res.Stdout)
		if state == "active" {
			fmt.Printf("[OK] %s: active\n", service)
		} else {
			fmt.Printf("[FAIL] %s: %s\n", service, state)
		}
		return nil

	case "restart":
		// Важно: без sudo это сработает только если у вас права.
		// Если будете запускать под root на сервере — ок.
		_, err := runner.Run(10*time.Second, "systemctl", "restart", service)
		if err != nil {
			return err
		}
		// После рестарта проверим
		return HandleSvc([]string{"status", service})

	default:
		return fmt.Errorf("unknown svc action: %s", action)
	}
}
