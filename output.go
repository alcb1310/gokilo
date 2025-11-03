package main

import (
	"fmt"
	"os"
)

func (e *EditorConfig) editorRefreshScreen() {
	fmt.Fprintf(os.Stdout, "\x1b[2J")
	fmt.Fprintf(os.Stdout, "\x1b[H")

	e.editorDrawRows()

	fmt.Fprintf(os.Stdout, "\x1b[H")
}

func (e *EditorConfig) editorDrawRows() {
	var y uint16

	for y = 0; y < e.term.ws.Row; y++ {
		fmt.Fprintf(os.Stdout, "~")

		if y < e.term.ws.Row-1 {
			fmt.Fprintf(os.Stdout, "\r\n")
		}
	}
}
