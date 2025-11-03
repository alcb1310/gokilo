package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	err := enableRawMode()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error enabling raw mode: %v\n", err)
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			os.Exit(1)
		}

		if b == 'q' {
			break
		}
	}
}
