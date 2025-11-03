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
	fmt.Fprintf(os.Stdout, "Rows %d\r\n", e.term.ws.Row)
	for range e.term.ws.Row {
		fmt.Fprintf(os.Stdout, "~\r\n")
	}
}
