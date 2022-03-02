package main

import (
	"fmt"
	"sort"
)

// In our implementation of B&B method
// we use bfs to explore the decision tree

// ref: https://www.geeksforgeeks.org/implementation-of-0-1-knapsack-using-branch-and-bound/

type item struct {
	w float32 // weight
	v int     // value
}

type items []item

// Note that we don't use pointer in knapsack func
// in case when both the scenarios including v and not including v
// have greater value than maxProfit
// so we need to enqueue both cases
type node struct {
	weight float32

	level int // level in the decision tree (i.e. index in items[])

	// profit by all nodes selected from those on the path from root to now
	// note that some node(s) on the path is/are not be selected
	profit int

	// the upper bound for the max profit in the subtree tree rooted at this node u
	// plus u.profit
	bound int // can be removed
}

// use greedy algorithm to cal. upper bound
// assuming all items are sorted in decreasing v/w
func bound(v node, its items, W float32) int {
	if v.weight > W {
		return 0
	}

	// initialize the bound to be profit of already selected nodes
	// same for totWeight
	profitBound := v.profit
	totWeight := v.weight

	j := v.level + 1
	for j < len(its) && (totWeight+its[j].w) < W {
		totWeight += its[j].w
		profitBound += its[j].v
		j++
	}

	// extract the max value with the remaining weight greedily
	// it's ok to slightly overestimate since
	// overestimating just relaxes the conditions for allowing
	// to go further down, and thus will never discard the subtree
	// that potentially has optimal sol.
	if j < len(its) {
		profitBound += int((W-totWeight)/its[j].w) * its[j].v
	}

	return profitBound
}

func knapsack(its items, W float32) int {
	sort.Sort(its)

	q := make([]node, 0, len(its)) // queue for bfs

	u := node{ // initialize u as dummy node
		weight: 0,
		profit: 0,
		level:  -1,
	}
	v := node{} // v is child of u

	q = append(q, u)
	maxProfit := 0 // max profit searched so far

	for len(q) > 0 {
		// extract from q
		u := q[0]
		q = q[1:]

		if u.level == len(its)-1 {
			continue // reach the bottom of the tree
		}

		v.level = u.level + 1
		v.weight = u.weight + its[v.level].w
		v.profit = u.profit + its[v.level].v
		if v.weight <= W && v.profit > maxProfit {
			maxProfit = v.profit
		}

		v.bound = bound(v, its, W)
		if v.bound > maxProfit {
			q = append(q, v)
		}

		// consider not including v in our sol.
		v.weight = u.weight
		v.profit = u.profit
		v.bound = bound(v, its, W)
		if v.bound > maxProfit {
			q = append(q, v)
		}
	}
	return maxProfit
}

func main() {
	its := items{
		{2, 40},
		{3.14, 50},
		{1.98, 100},
		{5, 95},
		{3, 30},
	}
	W := float32(10)

	fmt.Println("The max possible value: ")
	fmt.Println(knapsack(its, W))
}

// helper func to use sort package to sort the items
// in the decreasing order of v/w
func (its items) Less(i, j int) bool {
	return (float32(its[i].v) / its[i].w) > ((float32(its[j].v)) / its[j].w)
}

func (its items) Swap(i, j int) {
	its[i], its[j] = its[j], its[i]
}

func (its items) Len() int {
	return len(its)
}
