package main

import (
	"fmt"
)

// Problem:
// Given a number N and an infinite supply of valued coins
// S = {s1, s2, ..., sk}
// How many ways can we make change of N cents
// using the given coins?

// Analysis
//   Optimal substructure
//   Let c(m, n) be the number of ways we can make changes
//   for n cents using the proceeding m types of coins
//   then we have
//	     c(m, n) = c(m, n-S[m]) + c(m-1, n)
//             (using m-th coin) (not using m-th coin)
//   base case: c(m, 0) = 1
//				c(k, n) = 0, if n < 0 or k = 0

// Plot the table to help understanding
func Coin(s []int, n int) int {
	// dp[i][j] = c(n, m)
	// i: $$; j: the number of types of available coins
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, len(s)+1) // m = len(s)
		dp[i][0] = 0
	}
	for j := range dp[0] {
		dp[0][j] = 1
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= len(s); j++ {
			dp[i][j] = dp[i][j-1]
			if i-s[j-1] >= 0 { // note s[j-1] not s[j] !!
				dp[i][j] += dp[i-s[j-1]][j]
			}
		}
	}
	return dp[n][len(s)]
}

// improved version
// remember c(m, n) = c(m-1, n) + c(m, n-s[m])
// At each iteration of j
// c with smaller n has been updated from j-1 to j
// and then c with larger n is then being updated (from j-1 to j)
// by these "small n" c
func Coin2(s []int, n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for j := range s {
		for i := s[j]; i <= n; i++ {
			dp[i] += dp[i-s[j]] // orginal: c(j-1, n) -> update to c(j, n)
		}
	}
	return dp[n]
}

func main() {
	s := []int{10, 11, 30, 20, 1}
	n := 100
	fmt.Println(Coin(s, n))
	fmt.Println(Coin2(s, n))
}
