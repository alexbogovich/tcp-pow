package challenge

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

type HasherChallenger struct {
	ByteLen int

	value  string
	secret string
	count  int
}

func NewHasherChallenger(byteLen int, count int) *HasherChallenger {
	secret := strconv.Itoa(rand.Int())

	return &HasherChallenger{
		ByteLen: byteLen,
		secret:  secret,
		value:   strconv.Itoa(rand.Int()),
		count:   count,
	}
}

func (p HasherChallenger) Problem() Problem {
	hash := sha256.Sum256([]byte(p.value + "$" + p.secret))

	challenge := []byte(strconv.Itoa(p.count) + "||" + p.value + "||")
	challenge = append(challenge, hash[:p.ByteLen]...)

	return Problem{
		Challenge:      challenge,
		ExpectBytesLen: 512,
	}
}

func (p HasherChallenger) Verify(solution []byte) bool {
	chunks := strings.Split(string(solution), "||")
	if len(chunks) != p.count {
		println(fmt.Sprintf("wrong number of chunks: %d, expected: %d", len(chunks), p.count))
		return false
	}

	original := sha256.Sum256([]byte(p.value + "$" + p.secret))
	want := original[:p.ByteLen]

	for i := 0; i < len(chunks); i++ {
		hash := sha256.Sum256([]byte(p.value + "$" + chunks[i]))
		if !bytes.Equal(hash[:p.ByteLen], want) {
			return false
		}
	}

	return true
}

func HasherSolver(problem []byte) ([]byte, error) {
	chunks := strings.Split(string(problem), "||")
	if len(chunks) != 3 {
		return nil, fmt.Errorf("wrong number of chunks: %d, expected: %d", len(chunks), 3)
	}
	count, err := strconv.Atoi(chunks[0])
	if err != nil {
		return nil, fmt.Errorf("wrong count: %s", chunks[0])
	}

	value := chunks[1]
	want := []byte(chunks[2])
	bytesLen := len(want)

	var ans []string

	// brute force
	for i := math.MinInt; i < math.MaxInt; i++ {
		hash := sha256.Sum256([]byte(value + "$" + strconv.Itoa(i)))
		if bytes.Equal(hash[:bytesLen], want) {
			ans = append(ans, strconv.Itoa(i))
			if len(ans) == count {
				break
			}
		}
	}

	if len(ans) != count {
		return nil, fmt.Errorf("wrong count: %d, expected: %d", len(ans), count)
	}

	return []byte(strings.Join(ans, "||")), nil
}
