package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	workerpool "github.com/ppcamp/go-lib/workerpool"
)

func main() {
	executor := new(job)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
	// defer cancel()

	pool := workerpool.NewWorker(
		executor,
		workerpool.WithContext(ctx),
		workerpool.WithLogging(),
		workerpool.WithSetJobs(2),
	)
	pool.Process()

	<-ctx.Done()
	pool.Shutdown()
}

type job struct{}

func (s *job) Process(ctx context.Context) {
	id, ok := ctx.Value(workerpool.KeyId).(workerpool.InfoValue)
	if !ok {
		panic("fail")
	}

	// Forcing shutdown
	// err := mylib.ForceContextClose(ctx, func() error {
	// 	log.Printf(" -> [#%d] Waiting\n", id)
	// 	time.Sleep(3 * time.Second)
	// 	log.Printf(" -> [#%d] Waited\n", id)
	// 	return nil
	// })
	// if err != context.Canceled {
	// 	log.Printf("Forced due to ctx canceled: %v\n", err)
	// } else if err != nil {
	// 	log.Println(err)
	// } else {
	// 	log.Printf(" -> [#%d] Finished!!\n", id)
	// }

	select {
	case <-ctx.Done():
		return
	default:
		log.Printf(" -> [#%d] Waiting\n", id)
		time.Sleep(3 * time.Second)
		log.Printf(" -> [#%d] Waited\n", id)
	}
}
