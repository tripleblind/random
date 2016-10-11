package random

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/stretchr/testify/assert"
)

func TestKMSGen(t *testing.T) {

	const reps = 16

	assert := assert.New(t)

	source := NewKMS(kms.New(session.New(), aws.NewConfig().WithRegion("eu-west-1")))

	buf, err := source.Generate(16)

	assert.NoError(err)
	assert.Len(buf, 16)

	var (
		set = make(map[string]struct{}, reps)
		val = struct{}{}
	)

	for i := 0; i < reps; i++ {

		buf, err = source.Generate(16)
		assert.NoError(err)

		_, ok := set[string(buf)]
		assert.False(ok)

		set[string(buf)] = val

	}

}
