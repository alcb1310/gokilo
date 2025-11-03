package main

import (
	"fmt"
	"os"
)

type EditorConfig struct {
	exitFunction func()
	term         *terminal
}

var E EditorConfig

func (g *EditorConfig) safeExit(err error) {
	fmt.Fprintf(os.Stdout, "\x1b[2J")
	fmt.Fprintf(os.Stdout, "\x1b[H")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\r\n", err)
	}
	if g.exitFunction != nil {
		g.exitFunction()
	}
	os.Exit(0)
}

func initEditor() {
	E = EditorConfig{
		exitFunction: nil,
		term:         NewTerminal(),
	}
}

func main() {
	initEditor()
	var err error
	E.exitFunction, err = enableRawMode()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error enabling raw mode: %v\r\n", err)
		defer E.safeExit(err)
	}
	defer E.safeExit(nil)

	for {
		E.editorRefreshScreen()
		E.editorProcessKeypress()
	}
}
