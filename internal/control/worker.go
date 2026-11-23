package control

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/tedtasman/hyperstream/internal/contract"
	"github.com/tedtasman/hyperstream/internal/pipeline"
)

// select between timer with duration d or context cancelled
func sleep(ctx context.Context, d time.Duration) bool {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-t.C:
		return true
	case <-ctx.Done():
		return false
	}
}

type Worker struct {
	stream contract.Stream
	done   chan struct{}
}

// construct new worker from stream with done channel
func NewWorker(stream contract.Stream) *Worker {
	return &Worker{stream: stream, done: make(chan struct{})}
}

func (w *Worker) Start(ctx context.Context) {
	defer close(w.done)

	delay := time.Duration(0)
	failures := 0

	for {
		// wait for delay or context cancel
		if !sleep(ctx, delay) {
			fmt.Printf("Stopping stream %s\n", w.stream.Name)
			return
		}

		// attempt pipeline run
		err := pipeline.Run(w.stream)
		// reset failure count on success
		if err == nil {
			failures = 0
			delay = w.interval()
			// backoff on error
		} else if failures < w.stream.BackoffMaxFailures {
			failures++
			delay = w.backoff(failures)
			// stop on error out
		} else {
			fmt.Fprintf(os.Stderr, "Stopping stream %s after %d read failures\n", w.stream.Name, failures)
			return
		}
	}
}

// convert stream interval to time.Duration
func (w *Worker) interval() time.Duration {
	return time.Second * time.Duration(w.stream.IntervalSeconds)
}

// exponential backoff function
func (w *Worker) backoff(failures int) time.Duration {
	maxBackoff := time.Second * time.Duration(w.stream.BackoffMaxSeconds)
	d := w.interval() << min(failures, 16)
	if d > maxBackoff || d <= 0 {
		d = maxBackoff
	}
	return d
}

// signal worker finished
func (w *Worker) Done() <-chan struct{} {
	fmt.Printf("Stream %s finished\n", w.stream.Name)
	return w.done
}
