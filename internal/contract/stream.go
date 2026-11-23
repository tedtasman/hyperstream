package contract

import "encoding/json"

type SourceSpec struct {
	Kind   string          `json:"kind"`
	Config json.RawMessage `json:"config"`
}

type WorkerSpec struct {
	Kind   string          `json:"kind"`
	Config json.RawMessage `json:"config"`
}

type TransformSpec struct{}

type SinkSpec struct{}

type Stream struct {
	Name          string
	WorkerSpec    WorkerSpec
	SourceSpec    SourceSpec
	TransformSpec TransformSpec
	SinkSpec      SinkSpec
}
