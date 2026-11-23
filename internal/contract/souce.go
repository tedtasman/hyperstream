package contract

import (
	"context"
)

type Source interface {
	Read(ctx context.Context) ([]byte, error)
}
