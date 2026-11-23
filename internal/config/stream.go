package config

type Stream struct {
	Name          string
	WorkerSpec    WorkerSpec
	SourceSpec    SourceSpec
	TransformSpec TransformSpec
	SinkSpec      SinkSpec
}
