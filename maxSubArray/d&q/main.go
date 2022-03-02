package main

import "fmt"

// divide and conquer solution
// ref: CLRS P70-71

// T(1) = 1
// T(n) = 2T(n/2) + theta(n), for n > 1
// Time complexity: O(nlgn)

func main() {
	s := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	l, r, max := maxSubArray(s, 0, len(s)-1)
	fmt.Printf("max = [%d - %d] = %d\n", l, r, max)
}

func find(a []int, lo, hi int) (int, int, int) {
	m := int(uint(lo+hi) >> 1) // m = ceil(lo+hi)/2
	l, r := m, m

	sum, lmax := a[l], a[l]
	for i := l - 1; i >= 0; i-- {
		sum += a[i]
		if lmax < sum {
			lmax = sum
			l = i
		}
	}

	sum, rmax := 0, 0
	for j := r + 1; j <= hi; j++ { // j <"=" hi to avoid like two-element cases bug
		sum += a[j]
		if sum > rmax {
			rmax = sum
			r = j
		}
	}
	return l, r, (lmax + rmax)
}

// initial call: maxSubArray(a, 0, len(a)-1)
func maxSubArray(a []int, lo, hi int) (int, int, int) {
	if lo == hi {
		return lo, hi, a[lo] // base case: only one element
	}
	l, r, max := find(a, lo, hi)

	m := int(uint(lo+hi) >> 1) // m = ceil(lo+hi)/2
	if ll, lr, lmax := maxSubArray(a, lo, m); lmax > max {
		l, r, max = ll, lr, lmax
	}
	if rl, rr, rmax := maxSubArray(a, m+1, hi); rmax > max {
		l, r, max = rl, rr, rmax
	}
	return l, r, max
}
