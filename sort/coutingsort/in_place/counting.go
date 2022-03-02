package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// in place counting sort
// only allowed to use extra O(k) storage
// CLRS 8-2.e
// ref: https://my.oschina.net/pybug/blog/29286

func main() {
	n, k := 10000000, 10000000
	a := gen(n, k)
	countingsort(a, k)

	fmt.Println(sort.IntsAreSorted(a))
}

// generates n int over the range 0~k
func gen(n, k int) []int {
	a := make([]int, n)
	rand.Seed(time.Now().UnixNano())
	for i := range a {
		a[i] = rand.Int() % (k + 1)
	}

	return a
}

func countingsort(a []int, k int) {
	c := make([]int, k+1)
	for _, ele := range a {
		c[ele] += 1
	}

	// prefix sum
	ac, act := make([]int, k+1), make([]int, k+1)
	ac[0], act[0] = c[0], c[0]
	for i := 1; i < len(c); i++ {
		ac[i], act[i] = c[i]+ac[i-1], c[i]+act[i-1]
	}

	j := len(a) - 1
	for j >= 0 {
		val := a[j]

		// the j-th position is correct
		if j >= (ac[val]-c[val]) && j < ac[val] {
			act[val]--
			j-- // move to the (j-1)th position
			continue
		}

		// keep exchanging until the j-th postition has correct val
		a[act[val]-1], a[j] = val, a[act[val]-1]

		// the value originally at j-th postition is now in correct postition
		// so decrement by 1
		act[val]--
	}
}
