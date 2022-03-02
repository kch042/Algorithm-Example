package main

import "fmt"

type activity struct {
	s   int // start time
	f   int // finish time
	val int // profit
}

type dpInfo struct {
	j   int
	val int
}

// Weighted activity selection problem:
//   We have n jobs(activities),
//   each of which has an time interval and profit made
//   Suppose we cannot do more than one activity at any time point (compatibility)
//   Give an algorithm to find max value we could make
//   Moreover, list the compatible jobs that have max value

// Ref: CLRS Exercise 16.1-5

// weighted activity selection
func WAS(acts []activity) int {
	// sort acts by f

	dp := make([]dpInfo, len(acts)+1)
	dp[0] = dpInfo{val: acts[0].val, j: -1}
	for i := 1; i < len(acts); i++ {
		tmp, j := acts[i].val, find(acts, i)
		if j != -1 {
			tmp += dp[j].val
		}
		if tmp > dp[i-1].val {
			dp[i] = dpInfo{j: j, val: tmp} // include
		} else {
			dp[i] = dp[i-1] // not include
		}
	}
	fmt.Printf("The max profit: %d\n", dp[len(acts)-1].val)
	printActivities(dp, len(acts)-1, acts)
	return dp[len(acts)-1].val
}

func printActivities(dp []dpInfo, n int, acts []activity) {
	sp, stack := -1, make([]int, 0, len(dp))
	i := n
	for i >= 0 {
		val := dp[i].val
		for i > 0 && dp[i-1].val == val {
			i--
		}
		sp++
		stack, i = append(stack, i), dp[i].j
	}

	fmt.Println("Choose the job:")
	for sp > -1 {
		i := stack[sp]
		fmt.Printf("(%d, %d, %d)\n", acts[i].s, acts[i].f, acts[i].val)
		sp--
	}
	fmt.Println()
}

func sortAct(acts []activity) {
	n := len(acts)
	siftDown := func(i int, hi int) {
		for 2*i+1 < hi {
			child := 2*i + 1
			if child+1 < hi && acts[child].val < acts[child+1].val {
				child++
			}
			if acts[child].val > acts[i].val {
				acts[child], acts[i] = acts[i], acts[child]
				i = child
			} else {
				return
			}
		}
	}

	for i := n/2 - 1; i >= 0; i-- {
		siftDown(i, n)
	}
	j := n - 1
	for j > 0 {
		acts[0], acts[j] = acts[j], acts[0]
		siftDown(0, j+1)
		j--
	}
}

// e.g. Given A = [1, 2, 3, 4, 7, 8, 9] and A[n].s = 8
// return 5 since A[5] <= 8 and A[6] > 8
func find(acts []activity, n int) int {
	l, r := 0, n
	for l <= r {
		m := (l + r) / 2
		if acts[m].f <= acts[n].s {
			if acts[m+1].f <= acts[n].s {
				l = m + 1
			} else {
				return m
			}
		} else {
			r = m - 1
		}
	}
	return -1
}

func main() {
	jobs := []activity{
		{3, 5, 10},
		{1, 4, 30},
		{5, 9, 50},
		{0, 6, 60},
		{7, 8, 10},
		{5, 7, 30},
	}
	WAS(jobs)
}
