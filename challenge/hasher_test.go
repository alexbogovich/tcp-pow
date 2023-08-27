package challenge

import (
	"testing"
)

func Test(t *testing.T) {
	challenger := NewHasherChallenger(2, 3)
	solution, err := HasherSolver(challenger)
	if err != nil {
		t.Fatal(err)
	}

	if !challenger.Verify(solution) {
		t.Fatal("wrong solution")
	}
}
