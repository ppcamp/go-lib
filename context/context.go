package context

import (
	"context"
)

// ForceContextClose will ensure that, whether finish first (context close, or function finishes),
// stop the execution.
//
// Note
//
// With this approach, if the context are cancelled, you're state will be unkown or corrupted.
// Remember that, if you wanna make sure that the application guarantee the state, you must use
// just a for loop within a select block.
//
// Basic example of using this decorator
// 	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
//	defer cancel()
//
// 	err := ForceContextClose(ctx, func() error {
//		time.Sleep(4 * time.Second)
//		return nil
//	})
//	if err != nil {
//		print("context canceled, therefore, the block state will be unknown")
//	}
func ForceContextClose(ctx context.Context, fn func() error) error {
	err := make(chan error, 1)
	go func() { err <- fn() }()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case e := <-err:
		return e
	}
}
