package source

import (
	"encoding/json"
	"fmt"

	"github.com/tedtasman/hyperstream/internal/contract"
)

type constructor func(raw json.RawMessage) (contract.Source, error)

// single registration point for new source types
var constructors = map[string]constructor{
	"rest": NewRestSource,
}

// Generic source constructor
func New(spec contract.SourceSpec) (contract.Source, error) {
	constructor, ok := constructors[spec.Kind]
	if !ok {
		return nil, fmt.Errorf("unknown source kind %q", spec.Kind)
	}
	return constructor(spec.Config)
}
