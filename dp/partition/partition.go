package main

// problem:
// Given a slice of int a[0...n-1]
// Determine if a[0...n-1] can be partitioned into two slices
// which have equal sum

// Heuristic: 0-1 Knapsack problem

// Algo: DP
// dp[s][i]: whether s can be the sum of a subset of a[0...i]
// dp[s][i] = dp[s-a[i]][i-1] || dp[s][i-1], if s - a[i] >= 0
//          = dp[s][i-1] else
// Draw the dp table to understand better

// Time: pseudo-polynomial

func partition(a []int) bool {
	sum := 0
	for _, v := range a {
		sum += v
	}
	if sum&1 == 1 { // odd sum, cannot divide evenly
		return false
	}

	sum /= 2
	dp := make([]bool, sum+1)
	tmp := make([]bool, sum+1)
	dp[0], tmp[0] = true, true

	// use tmp to store dp[i-1][s] to build dp[i][s]
	for i := 0; i < len(a); i++ {
		// new iteration, tmp takes old dp (i.e. dp[i-1][s])
		// to compute dp[i][s]
		tmp, dp = dp, tmp

		for s := 1; s <= sum; s++ {
			if s-a[i] >= 0 {
				dp[s] = tmp[s-a[i]] || tmp[s]
			}
		}
	}

	return dp[sum]
}
