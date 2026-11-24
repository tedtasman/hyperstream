package contract

type Sink interface {
	Write(payload []byte) error
}
