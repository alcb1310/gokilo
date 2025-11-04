package main

func (e *EditorConfig) editorAppendRow(s []byte, len int) {
	row := Erow{
		size:  len,
		chars: make([]byte, len),
	}
	copy(row.chars, s)

	e.rows = append(e.rows, row)
	e.numrows++
}
