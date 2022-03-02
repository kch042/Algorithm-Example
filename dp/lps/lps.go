package main

import "fmt"

// longest palindrome subsequence
// example: seq = "character" -> lps = "carac"

// time: O(n^2)

/*
					|-- 2 + lps(s[i+1...j-1]), if seq[i] == seq[j]
   lps(s[i...j]) = -|
					|-- max(lps(s[i+1...j]), lps(s[i...j-1])), otherwise
*/

func lps(seq string) int {
	// dp[j][i] stores the max length of possible palindromes
	// in s[i,...,j], i < j
	// we use dp[j][i] instead of dp[i][j] so that
	// we could allocate memory more conveniently since j > i
	dp := make([][]int, len(seq))

	for i := range dp {
		dp[i] = make([]int, i+1) // if i = 8, dp[8][0~8] -> allocate 9!
		dp[i][i] = 1             // the i-th character is a palindrome of lenth 1
	}
	// case for offset = 1
	for i := 0; i < len(seq)-1; i++ {
		if seq[i] == seq[i+1] {
			dp[i+1][i] = 2
		} else {
			dp[i+1][i] = 1
		}
	}

	for offset := 2; offset < len(seq); offset++ {
		for i := 0; i < len(seq)-offset; i++ {
			j := i + offset
			if seq[i] == seq[j] {
				dp[j][i] = dp[j-1][i+1] + 2
			} else {
				dp[j][i] = max(dp[j][i+1], dp[j-1][i])
			}
		}
	}
	return dp[len(seq)-1][0]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	s := "character"
	// s2 := "underqualified"
	fmt.Println(lps(s))
}
