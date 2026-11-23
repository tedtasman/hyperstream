package transform

import "github.com/tedtasman/hyperstream/internal/contract"

func New(spec contract.TransformSpec) contract.Transform {
	return &JsonTransform{}
}
