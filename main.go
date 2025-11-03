package main

import (
	"bufio"
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

	reader := bufio.NewReader(os.Stdin)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input: %v\r\n", err)
			os.Exit(1)
		}

		if isCtrl(b) {
			fmt.Fprintf(os.Stdout, "%d\r\n", b)
		} else {
			fmt.Fprintf(os.Stdout, "%d ('%c')\r\n", b, b)
		}

		if b == ctrlKey('q') {
			break
		}
	}
}

func isCtrl(b byte) bool {
	return b <= 0x1f || b == 0x7f
}

func ctrlKey(b byte) byte {
	return b & 0x1f
}
