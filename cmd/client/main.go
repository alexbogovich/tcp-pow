package main

import (
	"fmt"
	"net"
	"os"
	"tcp-pow/challenge"
)

func main() {
	address, ok := os.LookupEnv("TCP_TARGET")
	if !ok {
		panic("TCP_TARGET env variable is not set")
	}

	conn, err := net.Dial("tcp", address)
	if err != nil {
		panic(fmt.Errorf("failed to connect to %s: %w", address, err))
	}

	buf := make([]byte, 1024)
	challengeN, err := conn.Read(buf)
	if err != nil {
		panic(fmt.Errorf("failed to read challenge: %w", err))
	}

	result, err := challenge.HasherSolver(buf[:challengeN])
	if err != nil {
		panic(fmt.Errorf("failed to process challenge: %w", err))
	}

	_, err = conn.Write(result)
	if err != nil {
		panic(fmt.Errorf("failed to write solution: %w", err))
	}

	buf = make([]byte, 256)
	n, err := conn.Read(buf)

	if err != nil {
		panic(fmt.Errorf("failed to read quote: %w", err))
	}

	println(string(buf[:n]))
}
