package lcs

// longest common subsequence

// m = len(s1), n = len(s2)
// Time: O(mn)
// Space: O(m)

// returns the number of lcs
func lcsNum(s1, s2 string) int {
	dp := make([]int, len(s1)+1)
	for j := 1; j <= len(s2); j++ {
		prev := dp[0]
		for i := 1; i <= len(s1); i++ {
			if s1[i-1] == s2[j-1] {
				dp[i], prev = prev+1, dp[i]
			} else {
				dp[i], prev = max(dp[i], dp[i-1]), dp[i]
			}
		}
	}
	return dp[len(s1)]
}

func lcs(s1, s2 string) {
	dp := make([][]int, len(s1)+1)
	for i := range dp {
		dp[i] = make([]int, len(s2)+1)
	}

	for j := 1; j <= len(s2); j++ {
		for i := 1; i <= len(s1); i++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	tmp := ""
	i, j := len(s1), len(s2)
	for i > 0 && j > 0 {
		if s1[i] == s2[j] {
			// add to the ret string stack
		}
	}
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
