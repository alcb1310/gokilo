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

func (t *terminal) editorReadKey() (int, error) {
	var ru rune
	var err error

	ru, _, err = t.reader.ReadRune()
	if ru == '\x1b' {
		var runs [3]rune
		var size int

		runs[0], size, err = t.reader.ReadRune()
		if size == 0 || err != nil {
			return '\x1b', nil
		}
		runs[1], size, err = t.reader.ReadRune()
		if size == 0 || err != nil {
			return '\x1b', nil
		}

		if runs[0] == '[' {
			switch runs[1] {
			case 'A':
				return ARROW_UP, nil
			case 'B':
				return ARROW_DOWN, nil
			case 'C':
				return ARROW_RIGHT, nil
			case 'D':
				return ARROW_LEFT, nil
			}
		}

		return '\x1b', nil
	}

	return int(ru), err
}

func (t *terminal) getWindowSize() (err error) {
	if t.ws, err = unix.IoctlGetWinsize(unix.Stdin, unix.TIOCGWINSZ); err != nil {
		fmt.Fprintf(os.Stderr, "getWindowSize: Error getting window size: %v\r\n", err)
		return
	}

	return
}
