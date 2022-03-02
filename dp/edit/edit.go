package main

import "fmt"

// edit distance
// leetcode 72

// dp(i, j) = dp(i-1, j-1), if word1[i] == word2[j]
//          = min{dp(i-1, j), dp(i, j-1), dp(i-1, j-1)} + 1, else
// base case: dp(i, 0) = dp(0, i) = i, for all i != 0
//            dp(0, 0) = 0

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}

	// base case init
	for i := 1; i <= len(word1); i++ {
		dp[i][0] = i
	}
	for j := 1; j <= len(word2); j++ {
		dp[0][j] = j
	}
	dp[0][0] = 0

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	fmt.Println(len(word1), len(word2))
	return dp[len(word1)][len(word2)]
}

func min(a, b, c int) int {
	if a > b {
		a = b
	}
	if a > c {
		return c
	}
	return a
}

func main() {
	s1 := ""
	s2 := "a"
	fmt.Println(minDistance(s1, s2))
}
