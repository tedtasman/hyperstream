package transform

import (
	"github.com/tedtasman/hyperstream/internal/config"
	"github.com/tedtasman/hyperstream/internal/contract"
)

// implements contract.Transform
type JsonTransform struct{ cfg config.TransformSpec }

func NewJsonTransform(spec config.TransformSpec) (contract.Transform, error) {
	return &JsonTransform{cfg: spec}, nil
}

// TODO: actually transform
func (t *JsonTransform) Apply(input []byte) ([]byte, error) {
	return input, nil
}
