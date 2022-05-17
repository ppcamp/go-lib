package timed

import (
	"context"
)

// ContextBlock will ensure that, whether finish first (context close, or function finishes), stop
// the execution.
//
// Example
// 	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
//	defer cancel()
//
// 	err := ContextBlock(ctx, func() error {
//		time.Sleep(4 * time.Second)
//		return nil
//	})
//	if err != nil {
//		print("context canceled, therefore, the block shouldn't be considered")
//	}
func ContextBlock(ctx context.Context, fn func() error) error {
	err := make(chan error, 1)
	go func() { err <- fn() }()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case e := <-err:
		return e
	}
}
