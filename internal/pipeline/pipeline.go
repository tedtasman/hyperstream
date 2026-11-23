package pipeline

import (
	"github.com/tedtasman/hyperstream/internal/contract"
	"github.com/tedtasman/hyperstream/internal/sink"
	"github.com/tedtasman/hyperstream/internal/source"
	"github.com/tedtasman/hyperstream/internal/transform"
)

// implements contract.Runner
type Pipeline struct {
	Source    contract.Source
	Transform contract.Transform
	Sink      contract.Sink
}

func New(stream contract.Stream) (*Pipeline, error) {
	newSource, err := source.New(stream.SourceSpec)
	if err != nil {
		return nil, err
	}
	newTransform := transform.New(stream.TransformSpec)
	newSink := sink.New(stream.SinkSpec)

	return &Pipeline{
		newSource,
		newTransform,
		newSink,
	}, nil
}

func (p *Pipeline) Run() error {
	return nil
}
