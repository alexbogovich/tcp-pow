package tcppow

import (
	"math/rand"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Word of Wisdom quotes
var quotes = []string{
	"Be yourself; everyone else is already taken.",
	"Two things are infinite: the universe and human stupidity; and I'm not sure about the universe.",
	"So many books, so little time.",
	"A room without books is like a body without a soul.",
}

type Problem struct {
	Challenge      []byte
	ExpectBytesLen int
}

type Challenger interface {
	Challenge() Problem
	Verify([]byte) bool
}

type PingPongChallenger struct{}

func (p PingPongChallenger) Challenge() Problem {
	return Problem{
		Challenge:      []byte("ping"),
		ExpectBytesLen: 4,
	}
}

func (p PingPongChallenger) Verify(solution []byte) bool {
	return string(solution) == "pong"
}

func tcpServerStart() (func() error, error) {
	l, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		return nil, err
	}

	go tcpListenerRun(l)

	return l.Close, nil
}

func tcpListenerRun(listener net.Listener) {
	// TODO: panic handling

	challenger := PingPongChallenger{}

	for {
		conn, err := listener.Accept()
		if err != nil {
			return
		}

		// TODO: deadline config
		conn.SetDeadline(time.Now().Add(1 * time.Second))

		go func() {

			problem := challenger.Challenge()

			// Make challenge
			conn.Write(problem.Challenge)

			// Read Solution
			buf := make([]byte, problem.ExpectBytesLen)
			conn.Read(buf)
			if !challenger.Verify(buf) {
				println("wrong message")
				conn.Close()
			} else {
				msg := quotes[rand.Intn(len(quotes))]
				conn.Write([]byte(msg))
			}

			conn.Close()
		}()
	}
}

func TestYolo(t *testing.T) {
	tcpServerStop, err := tcpServerStart()
	require.NoError(t, err)

	t.Cleanup(func() {
		tcpServerStop()
	})

	conn, err := net.Dial("tcp", "localhost:8888")
	require.NoError(t, err)

	buf := make([]byte, 4)
	_, err = conn.Read(buf)
	require.NoError(t, err, "expected to read challenge")
	assert.Equal(t, "ping", string(buf), "expected challenge: ping")

	_, err = conn.Write([]byte("pong"))
	require.NoError(t, err, "expected to write solution")

	buf = make([]byte, 256)
	n, err := conn.Read(buf)
	require.NoError(t, err, "expected to read quote")

	assert.Contains(t, quotes, string(buf[:n]), "expected one of the quotes on passing the challenge")

	conn.Close()
}
