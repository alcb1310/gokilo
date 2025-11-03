package main

import (
	"fmt"
	"os"
)

func editorRefreshScreen() {
	fmt.Fprintf(os.Stdout, "\x1b[2J")
}
