package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	n, k := 100000000, 500
	a := gen(n, k)
	// a = countingsort(a, k)
	sort.Ints(a)

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

func countingsort(a []int, k int) []int {
	c := make([]int, k+1) // 0 <= a[i] <= k
	for _, ele := range a {
		c[ele]++
	}
	// do prefix sum
	for i := 1; i <= k; i++ {
		c[i] += c[i-1]
	}

	b := make([]int, len(a))
	for j := len(a) - 1; j >= 0; j-- {
		b[c[a[j]]-1] = a[j]
		c[a[j]]--
	}
	return b
}
