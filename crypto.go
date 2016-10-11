package random

import (
	"crypto/rand"
	"fmt"
	"io"
)

var Crypto Source = cryptoSource{}

type cryptoSource struct{}

func (c cryptoSource) Fill(buf []byte) error {

	n, err := io.ReadFull(rand.Reader, buf)

	if err != nil {
		return fmt.Errorf("failed to fill buffer: %s", err)
	}

	if l := len(buf); l != n {
		return fmt.Errorf("failed to fill buffer with %d (wrote %d)", l, n)
	}

	return nil

}

func (c cryptoSource) Generate(bytes uint32) ([]byte, error) {

	buf := make([]byte, int(bytes))

	return buf, c.Fill(buf)

}
