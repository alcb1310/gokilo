package main

import (
	"fmt"
	"os"
)

func editorRefreshScreen() {
	fmt.Fprintf(os.Stdout, "\x1b[2J")
	fmt.Fprintf(os.Stdout, "\x1b[H")

	editorDrawRows()

	fmt.Fprintf(os.Stdout, "\x1b[H")
}

func editorDrawRows() {
	for range 24 {
		fmt.Fprintf(os.Stdout, "~\r\n")
	}
}
