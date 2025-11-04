package main

import (
	"fmt"
	"os"
)

func isCtrl(b byte) bool {
	return b <= 0x1f || b == 0x7f
}

func ctrlKey(b int) int {
	return b & 0x1f
}

// editorProcessKeypress processes a single keypress from the terminal
// it's job is to wait for one keypress and return it.
//
// @returns false when the user wants to exit
func (e *EditorConfig) editorProcessKeypress() {
	c, err := e.term.editorReadKey()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\r\n", err)
		E.safeExit(err)
		os.Exit(1)
	}

	switch c {
	case ctrlKey('q'):
		E.safeExit(nil)
		os.Exit(0)
	case ARROW_UP, ARROW_DOWN, ARROW_LEFT, ARROW_RIGHT:
		editorMoveCursor(c)
	}
}

func editorMoveCursor(key int) {
	switch key {
	case ARROW_LEFT:
		if E.cx != 0 {
			E.cx--
		}
	case ARROW_DOWN:
		if E.cy != E.term.ws.Row-1 {
			E.cy++
		}
	case ARROW_UP:
		if E.cy != 0 {
			E.cy--
		}
	case ARROW_RIGHT:
		if E.cx != E.term.ws.Col-1 {

			E.cx++
		}
	}
}
