package contract

type Transform interface {
	Apply(input []byte) ([]byte, error)
}
