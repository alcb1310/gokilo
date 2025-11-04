package main

import (
	"bufio"
	"fmt"
	"os"
)

func (e *EditorConfig) editorOpen(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\r\n", err)
		E.safeExit(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Bytes()
	linelen := len(line)

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\r\n", err)
		E.safeExit(err)
		os.Exit(1)
	}

	e.row.size = linelen
	e.row.chars = make([]byte, linelen)
	copy(e.row.chars, line)
	e.numrows++
}
