package stringencoder

import (
	"encoding/binary"
	"errors"
)

// Errors
var (
	ErrMalformed = errors.New("decoder: malformed input")
)

// Decoder decodes strings.
type Decoder struct {
	err error
	off int
	buf []byte
	val []byte
}

// NewDecoder with the buffer.
func NewDecoder(buf []byte) *Decoder {
	return &Decoder{
		buf: buf,
	}
}

// Err encountered during decoding.
func (d *Decoder) Err() error {
	return d.err
}

// Next value.
func (d *Decoder) Next() bool {
	if d.off >= len(d.buf) {
		return false
	}

	size, n := binary.Uvarint(d.buf[d.off:])

	if n <= 0 {
		d.err = ErrMalformed
		return false
	}

	d.off += n
	d.val = d.buf[d.off : d.off+int(size)]
	d.off += int(size)

	return true
}

// Bytes value.
func (d *Decoder) Bytes() []byte {
	return d.val
}

// String value.
func (d *Decoder) String() string {
	return string(d.Bytes())
}
