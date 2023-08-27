package tcppow

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

// Word of Wisdom quotes
var quotes = []string{
	"Be yourself; everyone else is already taken.",
	"Two things are infinite: the universe and human stupidity; and I'm not sure about the universe.",
	"So many books, so little time.",
	"A room without books is like a body without a soul.",
}

func TcpServerStart(cp ChallengerProvider, address string) (func() error, error) {
	l, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	go tcpListenerRun(cp, l)

	return l.Close, nil
}

func tcpListenerRun(cp ChallengerProvider, listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			return
		}

		conn.SetDeadline(time.Now().Add(60 * time.Second))

		go func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("panic in tcp connection processing", r)
				}
			}()

			// Make challenge
			challenger := cp()
			problem := challenger.Problem()
			conn.Write(problem.Challenge)

			// Read Solution
			buf := make([]byte, problem.ExpectBytesLen)
			conn.Read(buf)
			if challenger.Verify(buf) {
				msg := quotes[rand.Intn(len(quotes))]
				conn.Write([]byte(msg))
			} else {
				//println("wrong message")
			}

			conn.Close()
		}()
	}
}
