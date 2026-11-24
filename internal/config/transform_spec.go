package config

import "encoding/json"

type TransformSpec struct {
	Kind   string
	Config json.RawMessage
}
