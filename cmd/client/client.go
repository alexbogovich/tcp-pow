package main

import (
	"net"
	"tcp-pow/challenge"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 1024)
	challengeN, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}

	result, err := challenge.HasherSolver(buf[:challengeN])
	if err != nil {
		panic(err)
	}

	_, err = conn.Write(result)
	if err != nil {
		panic(err)
	}

	buf = make([]byte, 256)
	n, err := conn.Read(buf)

	if err != nil {
		panic(err)
	}

	println(string(buf[:n]))
}
