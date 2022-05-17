package pool

import "context"

type Worker interface {
	Process()
	Shutdown()
}

type Job interface {
	Process(ctx context.Context)
}
