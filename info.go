package rlog

import (
	"fmt"
)

// Info prints out an informative message to StdOut.
func Info(v ...any) {
	msg := fmt.Sprint(v...)
	msg = ColorMessage(msg, InfoColor)

	_ = info.Output(2, msg)
}

// Infof prints out a formatted informative message to StdOut.
func Infof(msg string, args ...any) {
	output := fmt.Sprintf(msg, args...)
	output = ColorMessage(output, InfoColor)

	_ = info.Output(2, output)
}
