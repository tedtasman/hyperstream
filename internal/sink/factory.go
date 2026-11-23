package sink

import (
	"github.com/tedtasman/hyperstream/internal/config"
	"github.com/tedtasman/hyperstream/internal/contract"
)

func New(spec config.SinkSpec) contract.Sink {
	return &RabbitSink{}
}
