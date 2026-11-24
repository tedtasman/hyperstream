package config

import "encoding/json"

type SinkSpec struct {
	Kind string
	cfg  json.RawMessage
}
