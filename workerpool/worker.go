package pool

import (
	"context"
	"log"
	"runtime"
	"sync/atomic"
)

type InfoKey string
type InfoValue int32

type worker struct {
	close   chan bool
	logging bool
	job     Job
	workers int32
	ctx     context.Context
}

func NewWorker(job Job, opts ...Option) Worker {
	option := &opt{
		Jobs:    int32(runtime.NumCPU()),
		Logging: false,
		Context: context.Background(),
	}

	for _, o := range opts {
		o(option)
	}

	return &worker{
		close:   make(chan bool),
		logging: option.Logging,
		job:     job,
		workers: option.Jobs,
		ctx:     option.Context,
	}
}

func (s *worker) Shutdown() {
	if s.logging {
		log.Println("Closing processor")
	}

	for i := int32(0); i < s.workers; i++ {
		s.close <- true // block here until some goroutine is free to close
	}

	if s.logging {
		log.Println("Processor closed!")
	}

	close(s.close)
}

func (s *worker) Process() {
	if s.logging {
		log.Printf("Starting %d jobs", s.workers)
	}

	for i := int32(0); i < s.workers; i++ {
		go s.do(atomic.LoadInt32(&i))
	}

	if s.logging {
		log.Printf("%d jobs started", s.workers)
	}

}

func (s *worker) do(tid int32) {
	if s.logging {
		log.Printf("[#%d] Started worker\n", tid)
	}
	for {
		select {
		case <-s.ctx.Done():
			if s.logging {
				log.Printf("[#%d] Closing the worker\n", tid)
			}
			<-s.close
			return

		case _, closed := <-s.close:
			if s.logging {
				if closed {
					log.Printf("[#%d] Closing the worker\n", tid)
				} else {
					log.Printf("[#%d] Closing the worker. Channel already closed!\n", tid)
				}
			}
			return

		default:
			if s.logging {
				log.Printf("[#%d] Processing some job\n", tid)
			}

			ctx := context.WithValue(s.ctx, InfoKey("id"), InfoValue(tid))
			s.job.Process(ctx)
		}
	}
}
