package env

import (
	"errors"
	"fmt"
	basestrings "strings"
	"syscall"

	"github.com/ppcamp/go-lib/strings"
)

var (
	ErrFlagRequired   error = errors.New("the flag is required")
	ErrUnexpectedType error = errors.New("unexpected type")
)

func fromEnv(envVar string) (string, bool) {
	envVar = basestrings.TrimSpace(envVar)
	return syscall.Getenv(envVar)
}

type Flag interface {
	Apply() error
}

// Parse the passed flags
func Parse(flags []Flag) error {
	for _, v := range flags {
		if err := v.Apply(); err != nil {
			return err
		}
	}
	return nil
}

type BaseFlagTypes interface {
	string | int | int64 | int32 | float32 | float64
}

// BaseFlag can be used for the BaseFlagTypes only.
//
// If you don't pass a Default value, the variable will be mandatory
type BaseFlag[T BaseFlagTypes] struct {
	// Value is the address of some variable, which can be some pkg variable, for example
	Value *T

	// Default is the default value to assign to this variable
	Default T

	// EnvName is the name of the environment variable that will try to fetch this data
	EnvName string
}

// isEmpty check if the value has the same value as an unitialized variable
func (s *BaseFlag[T]) isEmpty(value T) bool {
	var r T
	return r == value
}

func (s *BaseFlag[T]) Apply() error {
	valueFromEnv, exist := fromEnv(s.EnvName)
	// check if the flag don't exist and if there's no default value
	if !exist && s.isEmpty(s.Default) {
		return fmt.Errorf("flag %s is not defined: %w", s.EnvName, ErrFlagRequired)
	}

	// creates a pointer of the type T pointing to the response object and switch basing on the ptrs
	var response T
	switch p := any(&response).(type) {
	case *int:
		*p = strings.ToInt[int](valueFromEnv)
	case *int32:
		*p = strings.ToInt[int32](valueFromEnv)
	case *int64:
		*p = strings.ToInt[int64](valueFromEnv)
	case *string:
		*p = valueFromEnv
	case *float32:
		*p = strings.ToFloat[float32](valueFromEnv)
	case *float64:
		*p = strings.ToFloat[float64](valueFromEnv)
	default:
		return fmt.Errorf("type %T: %w", p, ErrUnexpectedType)
	}

	// update the value of the passed variable
	if s.isEmpty(response) {
		*s.Value = s.Default
	} else {
		*s.Value = response
	}

	return nil
}
