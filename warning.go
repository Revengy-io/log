package rlog

import (
	"fmt"
)

// Warn prints out a warning message to StdOut.
func Warn(v ...any) {
	msg := fmt.Sprint(v...)
	msg = ColorMessage(msg, WarningColor)

	_ = warning.Output(2, msg)
}

// Warnf prints out a formatted warning message to StdOut.
func Warnf(msg string, args ...any) {
	output := fmt.Sprintf(msg, args...)
	output = ColorMessage(output, WarningColor)

	_ = warning.Output(2, output)
}
