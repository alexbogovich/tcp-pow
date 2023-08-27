package tcppow

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net"
	"strings"
	"tcp-pow/challenge"
	"testing"
)

func TestHasherPositive(t *testing.T) {
	tcpServerStop, err := TcpServerStart(HasherProvider, "localhost:8888")
	require.NoError(t, err)

	t.Cleanup(func() {
		tcpServerStop()
	})

	conn, err := net.Dial("tcp", "localhost:8888")
	require.NoError(t, err)

	buf := make([]byte, 1024)
	challengeN, err := conn.Read(buf)
	require.NoError(t, err, "expected to read challenge")

	result, err := challenge.HasherSolver(buf[:challengeN])
	require.NoError(t, err, "expected to solve challenge")

	_, err = conn.Write(result)
	require.NoError(t, err, "expected to write solution")

	buf = make([]byte, 256)
	n, err := conn.Read(buf)
	require.NoError(t, err, "expected to read quote")

	assert.Contains(t, quotes, string(buf[:n]), "expected one of the quotes on passing the challenge")

	conn.Close()
}

func TestHasherNegative(t *testing.T) {
	tcpServerStop, err := TcpServerStart(HasherProvider, "localhost:8888")
	require.NoError(t, err)

	t.Cleanup(func() {
		tcpServerStop()
	})

	conn, err := net.Dial("tcp", "localhost:8888")
	require.NoError(t, err)

	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	require.NoError(t, err, "expected to read challenge")

	wrong := strings.Join([]string{"yollo", "bollo"}, "||")

	_, err = conn.Write([]byte(wrong))
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
