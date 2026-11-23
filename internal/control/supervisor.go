package control

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/tedtasman/hyperstream/internal/contract"
	"github.com/tedtasman/hyperstream/internal/pipeline"
	"github.com/tedtasman/hyperstream/internal/worker"
)

type entry struct {
	worker contract.Worker
	cancel context.CancelFunc
}

func Run(ctx context.Context, streams []contract.Stream) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	workers, err := buildWorkers(streams)
	if err != nil {
		return err
	}

	// start all only after successful build
	entries := startWorkers(ctx, workers)

	<-ctx.Done()

	for _, e := range entries {
		e.cancel()
	}
	for _, e := range entries {
		<-e.worker.Done()
	}
	return nil
}

func buildWorkers(streams []contract.Stream) ([]contract.Worker, error) {
	workers := make([]contract.Worker, 0, len(streams))
	for _, stream := range streams {
		pl, err := pipeline.New(stream)
		if err != nil {
			return nil, fmt.Errorf("build pipeline %s: %w", stream.Name, err)
		}
		w, err := worker.New(stream.WorkerSpec, pl, stream.Name)
		if err != nil {
			return nil, fmt.Errorf("build worker %s: %w", stream.Name, err)
		}
		workers = append(workers, w)
	}
	return workers, nil
}

func startWorkers(ctx context.Context, workers []contract.Worker) []entry {
	entries := make([]entry, 0, len(workers))
	for _, w := range workers {
		wctx, cancel := context.WithCancel(ctx)
		entries = append(entries, entry{worker: w, cancel: cancel})
		go w.Start(wctx)
	}
	return entries
}
