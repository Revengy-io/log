package rlog

import (
	"fmt"
)

// Error prints out an error message to StdErr.
func Error(v ...any) {
	msg := fmt.Sprint(v...)
	msg = ColorMessage(msg, ErrorColor)

	_ = err.Output(2, msg)
}

// Errorf prints out a formatted error message to StdErr.
func Errorf(msg string, args ...any) {
	output := fmt.Sprintf(msg, args...)
	output = ColorMessage(output, ErrorColor)

	_ = err.Output(2, output)
}
