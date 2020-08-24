// +build darwin windows

package log

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return false
}
