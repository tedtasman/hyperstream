package contract

import (
	"context"
)

type Worker interface {
	Start(ctx context.Context)
	Done() <-chan struct{}
}
