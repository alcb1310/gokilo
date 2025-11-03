package main

import (
	"fmt"
	"os"
)

type Global struct {
	f func()
}

var E Global

func (g *Global) safeExit(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\r\n", err)
	}
	if g.f != nil {
		g.f()
	}
	os.Exit(0)
}

func initEditor() {
	E = Global{
		f: nil,
	}
}

func main() {
	initEditor()
	var err error
	E.f, err = enableRawMode()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error enabling raw mode: %v\r\n", err)
		defer E.safeExit(err)
	}
	defer E.safeExit(nil)

	term := NewTerminal()

	for {
		editorRefreshScreen()
		editorProcessKeypress(term)
	}
}
