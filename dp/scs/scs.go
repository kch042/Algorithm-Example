package main

import "fmt"

// problem: Shortest Common Supersequence
// Given two string A, B
// Find the length of shortest string that contains both A and B as subsequences
// For example:
// A = "AXA"; B = "XB"
// SCS = "AXAB"

// Algo: DP
// dp[i, j] = length of scs of a[0...i-1] and b[0...j-1]
// note: dp[i, j] = length of scs of prefix of a with len i and prefix of b with length j

// dp[i, j] = dp[i-1, j-1] + 1, if a[i] = b[j]
//          = min{dp[i-1, j], dp[i, j-1]} + 1, else

func main() {
	a := "AXA"
	b := "XB"

	fmt.Println(scs(a, b))
}

func scs(a, b string) int {
	dp := make([][]int, len(a)+1)
	for i := range dp {
		dp[i] = make([]int, len(b)+1)
		dp[i][0] = i
	}
	for j := range dp[0] {
		dp[0][j] = j
	}

	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}

	// fmt.Println(len(dp))
	// fmt.Println(len(dp[0]))
	// for j := 0; j <= len(b); j++ {
	// 	for i := 0; i <= len(a); i++ {
	// 		fmt.Printf("%d ", dp[i][j])
	// 	}
	// 	fmt.Println()
	// }

	return dp[len(a)][len(b)]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
a = "AXA"; B = "XB"
    a   ∅   A   AX   AXA
b       0   1    2    3
∅   0   0   1    2    3
B   1   1   2    3    4
BX  2   2   3    3    4

*/

// SCS is closely related to LCS
// In fact, len(scs) = m + n - len(scs), where  m, n = len(a), len(b)
// Proof: We can create scs by taking lcs and insert non-lcs symbol of a and b
// into lcs while preserving the original order

// for lcs detail, check algo note!
func lcs(a, b string) int {
	m, n := len(a), len(b)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = 0
	}
	for j := range dp[0] {
		dp[0][j] = 0
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

func checkSCS(a, b string) int {
	return len(a) + len(b) - lcs(a, b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
