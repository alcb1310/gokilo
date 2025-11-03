package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

type terminal struct {
	reader *bufio.Reader
	ws     *unix.Winsize
}

func NewTerminal() *terminal {
	t := &terminal{
		reader: bufio.NewReader(os.Stdin),
	}
	t.getWindowSize()

	return t
}

func (t *terminal) editorReadKey() (byte, error) {
	return t.reader.ReadByte()
}

func (t *terminal) getWindowSize() (err error) {
	if t.ws, err = unix.IoctlGetWinsize(unix.Stdin, unix.TIOCGWINSZ); err != nil {
		fmt.Fprintf(os.Stderr, "getWindowSize: Error getting window size: %v\r\n", err)
		return
	}

	return
}
