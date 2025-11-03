package main

import (
	"bufio"
	"os"
)

type terminal struct {
	reader *bufio.Reader
}

func NewTerminal() *terminal {
	return &terminal{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (t *terminal) editorReadKey() (byte, error) {
	return t.reader.ReadByte()
}
