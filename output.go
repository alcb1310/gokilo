package main

import (
	"fmt"
	"os"
)

func (e *EditorConfig) editorRefreshScreen() {
	ab := NewAppendBuffer()

	ab.Append([]byte("\x1b[?25l"), 6)
	ab.Append([]byte("\x1b[H"), 3)

	e.editorDrawRows(ab)

	ab.Append([]byte("\x1b[H"), 3)
	ab.Append([]byte("\x1b[?25h"), 6)

	fmt.Fprintf(os.Stdout, "%s", ab.b)
}

func (e *EditorConfig) editorDrawRows(ab *AppendBuffer) {
	var y uint16

	for y = 0; y < e.term.ws.Row; y++ {
		if y == e.term.ws.Row/3 {
			welcome := fmt.Sprintf("Kilo editor -- version %s", KILO_VERSION)
			welcomelen := (uint16)(len(welcome))
			if welcomelen > e.term.ws.Col {
				welcomelen = e.term.ws.Col
				welcome = welcome[:welcomelen]
			}

			var padding uint16 = (e.term.ws.Col - welcomelen) / 2
			if padding > 0 {
				ab.Append([]byte("~"), 1)
				padding--
			}

			for ; padding > 0; padding-- {
				ab.Append([]byte(" "), 1)
			}

			ab.Append([]byte(welcome), len(welcome))
		} else {
			ab.Append([]byte("~"), 1)
		}

		ab.Append([]byte("\x1b[K"), 3)
		if y < e.term.ws.Row-1 {
			ab.Append([]byte("\r\n"), 2)
		}
	}
}
