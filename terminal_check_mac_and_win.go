// +build !linux

package log

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return false
}
