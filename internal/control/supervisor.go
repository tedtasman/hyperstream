package control

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/tedtasman/hyperstream/internal/contract"
)

type entry struct {
	worker *Worker
	cancel context.CancelFunc
}

func Run(ctx context.Context, streams []contract.Stream) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	entries := buildWorkers(ctx, streams)

	<-ctx.Done()

	for _, e := range entries {
		e.cancel()
	}
	for _, e := range entries {
		<-e.worker.Done()
	}
	return nil
}

func buildWorkers(ctx context.Context, streams []contract.Stream) []entry {
	entries := make([]entry, 0, len(streams))
	for _, st := range streams {
		wctx, cancel := context.WithCancel(ctx)
		w := NewWorker(st)
		entries = append(entries, entry{worker: w, cancel: cancel})
		go w.Start(wctx)
	}
	return entries
}
