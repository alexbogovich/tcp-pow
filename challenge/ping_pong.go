package challenge

type PingPongChallenger struct{}

func (p PingPongChallenger) Problem() Problem {
	return Problem{
		Challenge:      []byte("ping"),
		ExpectBytesLen: 4,
	}
}

func (p PingPongChallenger) Verify(solution []byte) bool {
	return string(solution) == "pong"
}
