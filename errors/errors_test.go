package errors_test

import (
	"errors"
	"testing"

	xtenderr "github.com/ppcamp/go-lib/errors"
	"github.com/stretchr/testify/assert"
)

func TestWraps(t *testing.T) {
	a := errors.New("some error")
	j := xtenderr.Wraps("err b", a)
	err := xtenderr.Wraps("err c", j)

	assert := assert.New(t)
	assert.ErrorIs(err, j)
	assert.ErrorIs(err, a)
}

func TestMust(t *testing.T) {
	assert := assert.New(t)

	fn := func(e error) (bool, error) { return true, e }

	assert.NotPanics(func() {
		k := xtenderr.Must(fn(nil))
		var expect bool
		assert.IsType(expect, k)
	})

	assert.Panics(func() {
		_ = xtenderr.Must(fn(errors.New("some err")))
	})
}

func TestPanicsIferror(t *testing.T) {
	assert := assert.New(t)

	fn := func(e error) error { return e }

	assert.NotPanics(func() { xtenderr.PanicIfError(fn(nil)) })
	assert.Panics(func() { xtenderr.PanicIfError(fn(errors.New("some err"))) })
}
