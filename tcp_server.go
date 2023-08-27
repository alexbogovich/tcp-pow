package tcppow

import (
	"math/rand"
	"net"
	"tcp-pow/challenge"
	"time"
)

// Word of Wisdom quotes
var quotes = []string{
	"Be yourself; everyone else is already taken.",
	"Two things are infinite: the universe and human stupidity; and I'm not sure about the universe.",
	"So many books, so little time.",
	"A room without books is like a body without a soul.",
}

type ChallengerProvider func() challenge.Challenger

var pingPongProvider = func() challenge.Challenger {
	return challenge.PingPongChallenger{}
}

var hasherProvider = func() challenge.Challenger {
	// bitLen and count may be dynamic depending on current metrics or whatever
	return challenge.NewHasherChallenger(2, 4)
}

func tcpServerStart(cp ChallengerProvider) (func() error, error) {
	l, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		return nil, err
	}

	go tcpListenerRun(cp, l)

	return l.Close, nil
}

func tcpListenerRun(cp ChallengerProvider, listener net.Listener) {
	// TODO: panic handling

	for {
		conn, err := listener.Accept()
		if err != nil {
			return
		}

		// TODO: deadline config
		conn.SetDeadline(time.Now().Add(60 * time.Second))

		go func() {
			challenger := cp()
			problem := challenger.Problem()

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
