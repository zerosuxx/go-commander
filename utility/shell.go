package utility

import (
	"io"
	"net/http"
	"os/exec"
)

type Shell struct {

}

func (s Shell) Run(command string, args []string, output io.Writer) error {
	cmd := exec.Command(command, args...)

	stdoutPipe, _ := cmd.StdoutPipe()
	stderrPipe, _ := cmd.StderrPipe()
	if f, ok := output.(http.Flusher); ok {
		f.Flush()
	}
	go func() {
		_, _ = io.Copy(output, stdoutPipe)
		_, _ = io.Copy(output, stderrPipe)
	}()

	return cmd.Run()
}
