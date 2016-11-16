package stringencoder_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	stringencoder "github.com/tj/go-stringencoder"
)

func TestEncoder(t *testing.T) {
	enc := stringencoder.NewEncoder(0)
	assert.NoError(t, enc.WriteString("foo"))
	assert.NoError(t, enc.WriteString("bar"))
	assert.NoError(t, enc.WriteBytes([]byte("baz")))
	assert.NoError(t, enc.WriteString("whatever"))

	dec := stringencoder.NewDecoder(enc.Bytes())

	var vals []string

	for dec.Next() {
		vals = append(vals, dec.String())
	}

	assert.Equal(t, []string{"foo", "bar", "baz", "whatever"}, vals)
	assert.NoError(t, dec.Err())
}

func BenchmarkEncoder(b *testing.B) {
	enc := stringencoder.NewEncoder(0)
	buf := []byte("hello world")
	for i := 0; i < b.N; i++ {
		enc.WriteBytes(buf)
	}
}

func BenchmarkDecoder(b *testing.B) {
	enc := stringencoder.NewEncoder(0)
	buf := []byte("hello world")
	ops := 100000

	for i := 0; i < ops; i++ {
		enc.WriteBytes(buf)
	}

	b.ResetTimer()
	b.SetBytes(int64(ops * len(buf)))
	blob := enc.Bytes()

	for i := 0; i < b.N; i++ {
		dec := stringencoder.NewDecoder(blob)

		for dec.Next() {
			_ = dec.Bytes()
		}

		if err := dec.Err(); err != nil {
			b.Fatalf("error decoding: %s", err)
		}
	}
}
