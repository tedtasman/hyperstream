package transform

import (
	"encoding/json"

	"github.com/tedtasman/hyperstream/internal/contract"
)

type IdentityConfig struct{}

// implements contract.Transform
type IdentityTransform struct{ cfg IdentityConfig }

func NewIdentityTransform(raw json.RawMessage) (contract.Transform, error) {
	var cfg IdentityConfig
	if err := json.Unmarshal(raw, &cfg); err != nil {
		return nil, err
	}
	return &IdentityTransform{cfg: cfg}, nil
}

// TODO: actually transform
func (t *IdentityTransform) Apply(input []byte) ([]byte, error) {
	return input, nil
}
