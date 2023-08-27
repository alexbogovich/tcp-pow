package challenge

type Problem struct {
	Challenge      []byte
	ExpectBytesLen int
}

type Challenger interface {
	Problem() Problem
	Verify([]byte) bool
}
