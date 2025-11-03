package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	safe, err := enableRawMode()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error enabling raw mode: %v\r\n", err)
		os.Exit(1)
	}
	defer safe()

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
