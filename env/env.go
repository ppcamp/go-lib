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

type Flag interface{ Apply() }

func Parse(flags []Flag) {
	for _, v := range flags {
		v.Apply()
	}
}

type types interface {
	string | int | int64 | int32 | float32 | float64
}

type BaseFlag[T types] struct {
	Pos     *T
	Default T
	EnvName string
}

func (s *BaseFlag[T]) Apply() {
	v, exist := fromEnv(s.EnvName)
	var response T

	if !exist {
		// check if there's no default value
		if s.Default == response {
			panic(fmt.Sprintf("flag %s is not defined", s.EnvName))
		}
		s.Pos = &s.Default
		return
	}

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

	s.Pos = &response
}
