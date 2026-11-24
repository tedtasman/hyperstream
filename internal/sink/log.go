package sink

import (
	"encoding/json"
	"fmt"

	"github.com/tedtasman/hyperstream/internal/contract"
)

type LogConfig struct{}

// implements contract.Sink
type LogSink struct{ cfg LogConfig }

func NewLogSink(raw json.RawMessage) (contract.Sink, error) {
	var cfg LogConfig
	if err := json.Unmarshal(raw, &cfg); err != nil {
		return nil, err
	}
	return &LogSink{cfg: cfg}, nil
}

func (s *LogSink) Write(payload []byte) error {
	fmt.Printf("[LogSink]: %s\n", payload)
	return nil
}
