package random

type Source interface {
	Fill([]byte) error
	Generate(uint32) ([]byte, error)
}
