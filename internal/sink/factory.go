package sink

import "github.com/tedtasman/hyperstream/internal/contract"

func New(spec contract.SinkSpec) contract.Sink {
	return &RabbitSink{}
}
