package timed

import (
	"context"
)

func TimedBlock(ctx context.Context, fn func() error) error {
	err := make(chan error, 1)
	go func() { err <- fn() }()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case e := <-err:
		return e
	}
}
