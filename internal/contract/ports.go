package contract

import (
	"context"
)

type Source interface {
	Read(ctx context.Context) ([]byte, error)
}

type Sink interface {
	Write()
}

type Worker interface {
	Start(ctx context.Context)
	Done() <-chan struct{}
}

type Transform interface {
	Apply()
}

type Runner interface {
	Run() error
}
