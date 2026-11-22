package contract

import "encoding/json"

type SourceSpec struct {
	Kind   string          `json:"kind"`
	Config json.RawMessage `json:"config"`
}

type Target struct {
	source SourceSpec
}
