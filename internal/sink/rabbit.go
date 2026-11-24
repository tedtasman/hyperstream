package sink

type RabbitConfig struct{}

// implements contract.Sink
type RabbitSink struct{ cfg RabbitConfig }

func (s *RabbitSink) Write(payload []byte) error {
	return nil
}
