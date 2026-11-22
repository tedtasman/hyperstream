package source

import (
	"encoding/json"
	"fmt"

	"github.com/tedtasman/sportstream/internal/contract"
)

type constructor func(raw json.RawMessage) (contract.Source, error)

// single registration point for new source types
var constructors = map[string]constructor{
	"rest": NewRestSource,
}

// Generic source constructor
func New(kind string, raw json.RawMessage) (contract.Source, error) {
	constructor, ok := constructors[kind]
	if !ok {
		return nil, fmt.Errorf("unknown source kind %q", kind)
	}
	return constructor(raw)
}
