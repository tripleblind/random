package randomness

import (
	"crypto/rand"
	"fmt"
	"io"
)

var Crypto Source = crypto{}

type crypto struct{}

func (c crypto) Fill(buf []byte) error {

	n, err := io.ReadFull(rand.Reader, buf)

	if err != nil {
		return fmt.Errorf("failed to fill buffer: %s", err)
	}

	if l := len(buf); l != n {
		return fmt.Errorf("failed to fill buffer with %d (wrote %d)", l, n)
	}

	return nil

}

func (c crypto) Generate(bytes uint32) ([]byte, error) {

	buf := make([]byte, int(bytes))

	return buf, c.Fill(buf)

}
