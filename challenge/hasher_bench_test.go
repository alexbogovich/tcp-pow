package challenge

import (
	"fmt"
	"testing"
)

var table = []struct {
	count int
	mask  int
}{
	//{count: 1, mask: 1},
	//{count: 4, mask: 1},
	//{count: 16, mask: 1},
	//{count: 32, mask: 1},

	{count: 1, mask: 2},
	{count: 4, mask: 2},
	{count: 16, mask: 2},
	{count: 32, mask: 2},

	{count: 1, mask: 3},
	{count: 4, mask: 3},
	{count: 16, mask: 3},
	{count: 32, mask: 3},
}

func BenchmarkPrimeNumbers(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("count_%d_mask_%d", v.count, v.mask), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				challenger := NewHasherChallenger(v.mask, v.count)
				solution, err := HasherSolver(challenger.Problem().Challenge)
				if err != nil {
					b.Fatal(err)
				}

				if !challenger.Verify(solution) {
					b.Fatal("wrong solution")
				}
			}
		})
	}
}
