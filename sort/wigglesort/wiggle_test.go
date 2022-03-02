package wiggle

import (
	"math/rand"
	"testing"
	"time"
)

func TestWiggle(t *testing.T) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		input := rand.Perm(10000)
		wiggle(input)

		c := 0
		for c < len(input)-1 {
			if c%2 == 0 && input[c] > input[c+1] {
				t.Error("Failed!\n")
			} else if c%2 == 1 && input[c] < input[c+1] {
				t.Error("Failed!\n")
			}
			c++
		}
	}
}
