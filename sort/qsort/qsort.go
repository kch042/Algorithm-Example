package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	// s := []int{3, 3, 2, 4, 5, 1}
	s := randInt(100000000)
	qsort2(s, 0, len(s)-1)
	fmt.Println(sort.IntsAreSorted(s))
}

func randInt(size int) []int {
	rand.Seed(time.Now().UnixNano())
	ret := make([]int, size)
	for i := range ret {
		ret[i] = (rand.Int() % 1000000)
	}
	return ret
}

// quicksort works well with distinct elements
// the performance decreases significantly when repetitions are high
// init call qsort(a, 0, len(a)-1)
func qsort(a []int, p, q int) {
	if p < q {
		r := partition(a, p, q)
		qsort(a, p, r-1)
		qsort(a, r+1, q)
	}
}

// traditional partition
func partition(a []int, p, q int) int {
	piv := a[p]
	i := p
	for j := i + 1; j <= q; j++ {
		if a[j] <= piv {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[p], a[i] = a[i], a[p]
	return i
}

func qsort2(a []int, p, q int) {
	if p < q {
		l, r := partition2(a, p, q)
		qsort2(a, p, l-1)
		qsort2(a, r+1, q)
	}
}

// improved partition (optimize dealing with the duplicates values)
// Draw the array pic to better understand the procedure
// | piv | < piv | == piv | > piv | ? |    (the array pic)
// p             l        r       j
func partition2(a []int, p, q int) (int, int) {
	piv := a[p]
	l, r := p, p // boundary index of the region where values == piv
	for j := p + 1; j <= q; j++ {
		if a[j] < piv {
			a[j], a[r+1] = a[r+1], a[j]
			l, r = l+1, r+1
			a[r], a[l] = a[l], a[r]
		} else if a[j] == piv {
			r++
			a[r], a[j] = a[j], a[r]
		}
	}
	a[l], a[p] = a[p], a[l]
	return l, r
}

//
func partition3() {

}
