package worker

import (
	"encoding/json"
	"fmt"

	"github.com/tedtasman/hyperstream/internal/config"
	"github.com/tedtasman/hyperstream/internal/contract"
)

type constructor func(raw json.RawMessage, runner contract.Runner, name string) (contract.Worker, error)

// single registration point for new source types
var constructors = map[string]constructor{
	"poll": NewPollWorker,
}

// Generic worker constructor
func New(spec config.WorkerSpec, runner contract.Runner, name string) (contract.Worker, error) {
	constructor, ok := constructors[spec.Kind]
	if !ok {
		return nil, fmt.Errorf("unknown worker kind %s", spec.Kind)
	}
	return constructor(spec.Config, runner, name)
}
