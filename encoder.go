package stringencoder

import "encoding/binary"

// Encoder encodes strings.
type Encoder struct {
	buf []byte
}

// NewEncoder with initial size.
func NewEncoder(size int) *Encoder {
	return &Encoder{
		buf: make([]byte, size),
	}
}

// Bytes returns the encoded strings.
func (e *Encoder) Bytes() []byte {
	return e.buf
}

// WriteString writes a string.
func (e *Encoder) WriteString(s string) error {
	return e.WriteBytes([]byte(s))
}

// WriteBytes writes a string.
func (e *Encoder) WriteBytes(b []byte) error {
	var buf [10]byte
	n := binary.PutUvarint(buf[:], uint64(len(b)))
	e.buf = append(e.buf, buf[:n]...)
	e.buf = append(e.buf, b...)
	return nil
}
