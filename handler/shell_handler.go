package handler

import (
	"errors"
	"fmt"
	"github.com/creack/pty"
	"golang.org/x/net/websocket"
	"io"
	"net/url"
	"os"
	"os/exec"
	"strconv"
)

func ShellServer(ws *websocket.Conn) {
	command := exec.Command("bash")

	winSize := pty.Winsize{}
	query := ws.Request().URL.Query()
	if colsUint, colsErr := GetUintFromValues(query, "cols"); colsErr == nil {
		winSize.Cols = colsUint
	}

	ptyTerminal, ptyErr := pty.StartWithSize(command, &winSize)

	defer func() {
		_ = command.Wait()
		_ = command.Process.Release()
		_ = command.Process.Signal(os.Interrupt)
		_ = ptyTerminal.Close()
	}()

	if ptyErr != nil {
		_, _ = ws.Write([]byte(fmt.Sprintf("Error creating pty: %s\r\n", ptyErr)))
		_ = ws.Close()
		return
	}

	go func() {
		_, _ = io.Copy(ws, ptyTerminal)
	}()
	_, _ = io.Copy(ptyTerminal, ws)
	_ = ws.Close()
}

func GetUintFromValues(values url.Values, key string) (uint16, error) {
	value := values.Get(key)
	if value == "" {
		return 0, errors.New("empty " + key)
	}

	valueInt, err := strconv.ParseUint(value, 16, 16)
	if err != nil {
		return 0, err
	}

	return uint16(valueInt), nil
}
