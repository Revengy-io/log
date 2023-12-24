package rlog

import (
	"fmt"
)

// Debug prints out a debug message to StdOut.
func Debug(v ...any) {
	msg := fmt.Sprint(v...)
	msg = ColorMessage(msg, DebugColor)

	_ = debug.Output(2, msg)
}

// Debugf prints out a formatted debug message to StdOut.
func Debugf(msg string, args ...any) {
	output := fmt.Sprintf(msg, args...)
	output = ColorMessage(output, DebugColor)

	_ = debug.Output(2, output)
}
