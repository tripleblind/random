package randomness

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCryptoGen(t *testing.T) {

	const reps = 65536

	assert := assert.New(t)

	buf, err := Crypto.Generate(16)

	assert.NoError(err)
	assert.Len(buf, 16)

	var (
		set = make(map[string]struct{}, reps)
		val = struct{}{}
	)

	for i := 0; i < reps; i++ {

		buf, err = Crypto.Generate(16)
		assert.NoError(err)

		_, ok := set[string(buf)]
		assert.False(ok)

		set[string(buf)] = val

	}

}
