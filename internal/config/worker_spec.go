package config

import "encoding/json"

type WorkerSpec struct {
	Kind   string          `json:"kind"`
	Config json.RawMessage `json:"config"`
}
