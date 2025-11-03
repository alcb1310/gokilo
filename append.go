package main

type AppendBuffer struct {
	b   []byte
	len int
}

func NewAppendBuffer() *AppendBuffer {
	return &AppendBuffer{
		b:   make([]byte, 0),
		len: 0,
	}
}

func (ab *AppendBuffer) Append(b []byte, l int) {
	ab.b = append(ab.b, b...)
	ab.len += l
}
