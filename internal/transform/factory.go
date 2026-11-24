package transform

import (
	"github.com/tedtasman/hyperstream/internal/config"
	"github.com/tedtasman/hyperstream/internal/contract"
)

func New(spec config.TransformSpec) contract.Transform {
	return &JsonTransform{cfg: spec}
}
