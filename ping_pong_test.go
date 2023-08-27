package tcppow

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net"
	"testing"
)

func TestPingPongPositive(t *testing.T) {
	tcpServerStop, err := tcpServerStart(pingPongProvider)
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

func TestPingPongNegative(t *testing.T) {
	tcpServerStop, err := tcpServerStart(pingPongProvider)
	require.NoError(t, err)

	t.Cleanup(func() {
		tcpServerStop()
	})

	conn, err := net.Dial("tcp", "localhost:8888")
	require.NoError(t, err)

	buf := make([]byte, 1024)
	challengeN, err := conn.Read(buf)
	require.NoError(t, err, "expected to read challenge")
	assert.Equal(t, "ping", string(buf[:challengeN]), "expected challenge: ping")

	_, err = conn.Write([]byte("wrong"))
	require.NoError(t, err, "expected to write solution")

	buf = make([]byte, 256)
	_, err = conn.Read(buf)

	switch {
	case err.Error() == "EOF":
		assert.Error(t, err, "expected to fail on wrong solution")
	default:
		assert.ErrorContains(t, err, "read: connection reset by peer", "expected to fail on wrong solution")
	}

	conn.Close()
}
