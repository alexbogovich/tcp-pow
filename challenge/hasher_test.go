package challenge

import (
	"bytes"
	"crypto/sha256"
	"testing"
)

func TestHasherChallenger_Verify(t *testing.T) {
	challenger := NewHasherChallenger(2, 3)
	solution, err := HasherSolver(challenger.Problem().Challenge)
	if err != nil {
		t.Fatal(err)
	}

	if !challenger.Verify(solution) {
		t.Fatal("wrong solution")
	}
}

func TestHasherChallenger_Problem(t *testing.T) {
	t.Run("correct mask", func(t *testing.T) {
		c := HasherChallenger{value: "421", secret: "78", ByteLen: 4, count: 3}

		sum256 := sha256.Sum256([]byte("421" + "$" + "78"))
		wantHash := sum256[:4]

		want := []byte("3||421||")
		want = append(want, wantHash...)

		got := c.Problem().Challenge
		if !bytes.Equal(got, want) {
			t.Fatalf("wrong challenge: %s, expected: %s", got, want)
		}
	})
}
