package main

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

func enableRawMode() (func(), error) {
	termios, err := unix.IoctlGetTermios(unix.Stdin, unix.TCGETS)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting terminal attributes: %v\n", err)
		return nil, err
	}
	orig_termios := *termios

	termios.Lflag &^= unix.ECHO | unix.ICANON | unix.IEXTEN | unix.ISIG
	termios.Iflag &^= unix.BRKINT | unix.ICRNL | unix.INPCK | unix.ISTRIP | unix.IXON
	termios.Oflag &^= unix.OPOST
	termios.Cflag &^= unix.CS8

	if err := unix.IoctlSetTermios(unix.Stdin, unix.TCSETS, termios); err != nil {
		fmt.Fprintf(os.Stderr, "Error setting terminal attributes: %v\r\n", err)
		return nil, err
	}

	return func() {
		if err := unix.IoctlSetTermios(unix.Stdin, unix.TCSETS, &orig_termios); err != nil {
			fmt.Fprintf(os.Stderr, "Error restoring terminal attributes: %v\r\n", err)
			os.Exit(1)
		}

		os.Exit(0)
	}, nil
}
