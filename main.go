package main

import (
	"fmt"
	"log/slog"
	"os"
)

const (
	KILO_VERSION = "0.0.1"
)

const (
	ARROW_LEFT = iota + 1000
	ARROW_RIGHT
	ARROW_UP
	ARROW_DOWN
	HOME_KEY
	END_KEY
	PAGE_UP
	PAGE_DOWN
)

type EditorConfig struct {
	cx, cy       uint16
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

func init() {
	var f *os.File
	var err error
	userTempDir, _ := os.UserConfigDir()
	if f, err = createLoggerFile(userTempDir); err != nil {
		E.safeExit(err)
	}

	handlerOptions := &slog.HandlerOptions{}
	handlerOptions.Level = slog.LevelDebug

	loggerHandler := slog.NewTextHandler(f, handlerOptions)
	slog.SetDefault(slog.New(loggerHandler))

	E = EditorConfig{
		exitFunction: nil,
		term:         NewTerminal(),
		cx:           0,
		cy:           0,
	}
}

func main() {
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
