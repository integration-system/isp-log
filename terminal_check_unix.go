package log

import (
	"golang.org/x/sys/unix"
	"io"
	"os"
)

const ioctlReadTermios = unix.TCGETS

func checkIfTerminal(w io.Writer) bool {
	switch v := w.(type) {
	case *os.File:
		return isTerminal(int(v.Fd()))
	default:
		return false
	}
}

func isTerminal(fd int) bool {
	_, err := unix.IoctlGetTermios(fd, ioctlReadTermios)
	return err == nil
}
