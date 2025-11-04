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

	msg := fmt.Sprintf("\x1b[%d;%dH", e.cy+1, e.cx+1)
	msglen := len(msg)
	ab.Append([]byte(msg), msglen)

	ab.Append([]byte("\x1b[?25h"), 6)

	fmt.Fprintf(os.Stdout, "%s", ab.b)
}

func (e *EditorConfig) editorDrawRows(ab *AppendBuffer) {
	var y uint16

	for y = 0; y < e.term.ws.Row; y++ {
		if y >= uint16(e.numrows) {
			if e.numrows == 0 && y == e.term.ws.Row/3 {
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
		} else {
			len := min(e.row.size, (int)(e.term.ws.Col))
			ab.Append(e.row.chars[:len], len)
		}

		ab.Append([]byte("\x1b[K"), 3)
		if y < e.term.ws.Row-1 {
			ab.Append([]byte("\r\n"), 2)
		}
	}
}
