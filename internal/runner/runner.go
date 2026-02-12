package runner

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"time"
)

type Result struct {
	Stdout string
	Stderr string
}

func Run(timeout time.Duration, name string, args ...string) (Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, name, args...)

	var outBuf bytes.Buffer
	var errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	err := cmd.Run()

	res := Result{
		Stdout: outBuf.String(),
		Stderr: errBuf.String(),
	}

	if ctx.Err() == context.DeadlineExceeded {
		return res, fmt.Errorf("command timed out: %s %v", name, args)
	}

	if err != nil {
		// Пробрасываем stderr, чтобы было понятно “почему”
		if res.Stderr != "" {
			return res, fmt.Errorf("%v: %s", err, res.Stderr)
		}
		return res, err
	}

	return res, nil
}

