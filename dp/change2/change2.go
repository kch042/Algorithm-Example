package main

import (
	"fmt"
	"math"
)

// Give n and S = {s1, s2...}
// Find the fewest number such that n can be summed up
// by a linear combination of si

// For example:
// S = {9, 6, 5, 1}, n = 11
// output = 2 (11 = 6 + 5)

// Recurrence
// def: c(n) = fewest number of si that constitute n
// c(n) = min{c(n-s[j]) + 1}, for all j s.t. s[j] <= n
// c(0) = 0

func fewestCoins(s []int, n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	for i := 1; i <= n; i++ {
		for j := range s {
			if s[j] <= i {
				dp[i] = min(dp[i], 1+dp[i-s[j]])
			}
		}
	}
	if dp[n] == math.MaxInt32 {
		return -1
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	s := []int{1}
	n := 1
	fmt.Println(fewestCoins(s, n))
}
