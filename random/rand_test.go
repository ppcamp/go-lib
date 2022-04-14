package random_test

import (
	"testing"

	"github.com/ppcamp/go-lib/random"
	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	assert := require.New(t)

	assert.NotPanics(func() {
		_ = random.String(30)
	})

	assert.Len(random.String(30), 30)
}
