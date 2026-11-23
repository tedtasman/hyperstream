package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/tedtasman/hyperstream/internal/contract"
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

type PollConfig struct {
	IntervalSeconds    int `json:"intervalSeconds"`
	BackoffMaxSeconds  int `json:"backoffMaxSeconds"`
	BackoffMaxFailures int `json:"backoffMaxFailures"`
}

// implements contract.Worker
type PollWorker struct {
	done   chan struct{}
	runner contract.Runner
	cfg    PollConfig
	Name   string
}

// construct new polling worker from stream with done channel
func NewPollWorker(raw json.RawMessage, runner contract.Runner, name string) (contract.Worker, error) {

	var cfg PollConfig
	if err := json.Unmarshal(raw, &cfg); err != nil {
		return nil, err
	}

	return &PollWorker{
		done:   make(chan struct{}),
		runner: runner,
		cfg:    cfg,
		Name:   name,
	}, nil
}

func (w *PollWorker) Start(ctx context.Context) {
	defer close(w.done)

	delay := time.Duration(0)
	failures := 0

	for {
		// wait for delay or context cancel
		if !sleep(ctx, delay) {
			fmt.Printf("Stopping stream %s\n", w.Name)
			return
		}

		// attempt pipeline run
		err := w.runner.Run()
		// reset failure count on success
		if err == nil {
			failures = 0
			delay = w.interval()
			// backoff on error
		} else if failures < w.cfg.BackoffMaxFailures {
			failures++
			delay = w.backoff(failures)
			// stop on error out
		} else {
			fmt.Fprintf(os.Stderr, "Stopping stream %s after %d read failures\n", w.Name, failures)
			return
		}
	}
}

// convert stream interval to time.Duration
func (w *PollWorker) interval() time.Duration {
	return time.Second * time.Duration(w.cfg.IntervalSeconds)
}

// exponential backoff function
func (w *PollWorker) backoff(failures int) time.Duration {
	maxBackoff := time.Second * time.Duration(w.cfg.BackoffMaxSeconds)
	d := w.interval() << min(failures, 16)
	if d > maxBackoff || d <= 0 {
		d = maxBackoff
	}
	return d
}

// signal worker finished
func (w *PollWorker) Done() <-chan struct{} {
	fmt.Printf("Stream %s finished\n", w.Name)
	return w.done
}
