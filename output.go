package main

import (
	"fmt"
	"os"
)

func (e *EditorConfig) editorRefreshScreen() {
	ab := NewAppendBuffer()
	ab.Append([]byte("\x1b[2J"), 4)
	ab.Append([]byte("\x1b[H"), 3)

	e.editorDrawRows(ab)

	ab.Append([]byte("\x1b[H"), 3)

	fmt.Fprintf(os.Stdout, "%s", ab.b)
}

func (e *EditorConfig) editorDrawRows(ab *AppendBuffer) {
	var y uint16

	for y = 0; y < e.term.ws.Row; y++ {
		ab.Append([]byte("~"), 1)

		if y < e.term.ws.Row-1 {
			ab.Append([]byte("\r\n"), 2)
		}
	}
}
