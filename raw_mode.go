package main

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

func enableRawMode() error {
	termios, err := unix.IoctlGetTermios(unix.Stdin, unix.TCGETS)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting terminal attributes: %v\n", err)
		return err
	}

	termios.Lflag &^= unix.ECHO

	if err := unix.IoctlSetTermios(unix.Stdin, unix.TCSETS, termios); err != nil {
		fmt.Fprintf(os.Stderr, "Error setting terminal attributes: %v\n", err)
		return err
	}

	return nil
}
