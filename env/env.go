package env

import (
	"fmt"
	basestrings "strings"
	"syscall"

	"github.com/ppcamp/go-lib/strings"
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

func (s *BaseFlag[T]) Apply() error {
	var response T

	v, exist := fromEnv(s.EnvName)
	if !exist {
		// check if there's no default value
		if s.Default == response {
			return fmt.Errorf("flag %s is not defined", s.EnvName)
		}
		s.Value = &s.Default
	}

	// creates a pointer of the type T pointing to the response object and switch basing on the ptrs
	switch p := any(&response).(type) {
	case *int:
		*p = strings.ToInt[int](v)
	case *int32:
		*p = strings.ToInt[int32](v)
	case *int64:
		*p = strings.ToInt[int64](v)
	case *string:
		*p = v
	case *float32:
		*p = strings.ToFloat[float32](v)
	case *float64:
		*p = strings.ToFloat[float64](v)
	}

	// update the value of the passed variable
	*s.Value = response
	return nil
}
