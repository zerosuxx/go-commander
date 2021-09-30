package utility

import (
	"os/exec"
)

type Shell struct {
}

func (s Shell) Exec(command string, args []string) ([]byte, error) {
	cmd := exec.Command(command, args...)

	defer func() {
		_ = cmd.Wait()
	}()

	return cmd.CombinedOutput()
}
