package random_test

import (
	"testing"

	"github.com/ppcamp/go-lib/random"
	"github.com/stretchr/testify/require"
)

func TestRandString(t *testing.T) {
	assert := require.New(t)

	assert.NotPanics(func() {
		_ = random.RandString(30)
	})

	assert.Len(random.RandString(30), 30)
}
