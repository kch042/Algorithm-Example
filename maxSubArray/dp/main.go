package main

import "fmt"

// ref
// 1. https://medium.com/@rsinghal757/kadanes-algorithm-dynamic-programming-how-and-why-does-it-work-3fd8849ed73d
// 2. CLRS P75 Exercise 4.1-5

func main() {
	s := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	fmt.Println(maxSubArray2(s))
}

// dp solution
func maxSubArray(a []int) {
	M, r := a[0], -1

	n := len(a)
	dp := a[0]
	for i := 1; i < n; i++ {
		dp = max(a[i]+dp, a[i])
		if M < dp {
			M = dp
			r = i
		}
	}

	// trace back to find l
	l, sum := r, M
	for l >= 0 && sum != 0 {
		sum -= a[l]
		l--
	}
	l++

	fmt.Printf("Max Profit: $%d\n", M)
	fmt.Printf("From %d to %d\n", l, r)
}

// return l, r, max s.t. max = a[l] + a[l-1] + ... + a[r]
func maxSubArray2(a []int) (int, int, int) {
	var nl int // stores the start location of possible maxSubarray

	l, r := 0, 0
	ret, dp := a[0], a[0]
	for i := 1; i < len(a); i++ {
		if dp < 0 { // the max sum including a[i-1] is negative, don't include them
			nl, dp = i, a[i]
		} else {
			dp += a[i]
		}
		if ret < dp {
			ret = dp
			l, r = nl, i
		}
	}
	return l, r, ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
