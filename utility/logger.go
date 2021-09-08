package utility

import (
    "log"
    "os"
    "runtime/debug"
)

func LogError(err error) {
    log.Print(err)
    _, _ = os.Stderr.Write(debug.Stack())
}