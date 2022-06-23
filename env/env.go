package env

import (
	"fmt"
	"strconv"
	"strings"
	"syscall"
)

func fromEnv(envVar string) (string, bool) {
	envVar = strings.TrimSpace(envVar)
	return syscall.Getenv(envVar)
}

type Flag interface{ Apply() }

func Parse(flags []Flag) error {
	for _, v := range flags {
		v.Apply()
	}
	return nil
}

type types interface {
	string | int | int64 | int32 | float32 | float64
}

type BaseFlag[T types] struct {
	Pos       *T
	Default   T
	EnvName   string
	Mandatory bool
}

func toInt[T int | int32 | int64](v string) T {
	a, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	return T(a)
}

func toFloat[T float32 | float64](v string) T {
	a, err := strconv.ParseFloat(v, 64)
	if err != nil {
		panic(err)
	}
	return T(a)
}

func (s *BaseFlag[T]) Apply() {
	v, exist := fromEnv(s.EnvName)

	if !exist {
		if s.Mandatory {
			panic(fmt.Sprintf("flag %s is not defined", s.EnvName))
		}

		s.Pos = &s.Default
		return
	}

	var ret T
	switch p := any(&ret).(type) {
	case *int:
		*p = toInt[int](v)
	case *int32:
		*p = toInt[int32](v)
	case *int64:
		*p = toInt[int64](v)
	case *string:
		*p = v
	case *float32:
		*p = toFloat[float32](v)
	case *float64:
		*p = toFloat[float64](v)
	}

	s.Pos = &ret
}
