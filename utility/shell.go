package utility

import (
	"context"
	"os/exec"
	"time"
)

type Shell struct {
}

func (s Shell) Exec(command string, args []string, waitInSeconds time.Duration) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), waitInSeconds*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, command, args...)

	defer func() {
		_ = cmd.Wait()
	}()

	return cmd.CombinedOutput()
}
