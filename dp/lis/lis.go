package main

import "fmt"

// longest increasing subsequence

// O(n^2)
// dp[i] = length of longest subsequence of
// a[0...i] containing a[i]
// so max is not necessarily a[n-1]
func lis1(a []int) int {
	n := len(a)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if a[j] < a[i] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
	}

	ret := 0
	for _, x := range dp {
		if ret < x {
			ret = x
		}
	}
	return ret
}

// O(nlgn)
// https://www.geeksforgeeks.org/longest-monotonically-increasing-subsequence-size-n-log-n/
func lis2(a []int) int {
	length := 0
	// dp stores the elements that are "potentially" an element of a lis
	dp := make([]int, len(a))

	dp[0] = a[0]
	length++

	for i := 1; i < len(a); i++ {
		if a[i] > dp[length-1] {
			length++
			dp[length-1] = a[i]
		} else {
			dp[binarySearch(a[i], length, dp)] = a[i]
		}
	}

	return length
}

// O(nlgn) sol with construction of
// https://www.geeksforgeeks.org/construction-of-longest-monotonically-increasing-subsequence-n-log-n/
// lis2_construction returns "a" longest increasing subsequence
func lis2_construction(a []int) []int {
	n := len(a)
	length := 0

	// t stores "indices" of elements that are "potentially" members of a lis
	// prev[i] stores the "index" of the prev element of a[i] in a lis
	// for example:
	// a = [4, 5, 1]
	// then t = [4] (here we write elements but in the code t stores indices)
	//      -> t = [4, 5]
	//      -> t = [1, 5] (4 is replaced by 1 since 4 is the smallest element larger than 1)
	// 	    prev = [4] (still write in elements, not indices for conv.)
	//      -> prev = [4, 4]
	//      -> prev = [4, 4, 1]
	// think of it as a linked list, element 5 points to 4 which points to itself
	// while element 1 points to itself

	t, prev := make([]int, n), make([]int, n)

	t[0], prev[0] = 0, 0
	length++

	for i := 1; i < n; i++ {
		if k := a[i]; k > a[t[length-1]] {
			t[length] = i
			prev[i] = t[length-1]
			length++
		} else if k <= a[t[0]] {
			// include k = a[t[0]] for reason that
			// otherwise m might be 0 which causes t[m-1] to index -1 -> panic
			// for example: a = [7, 7]
			t[0] = i
			prev[0] = i // may be a new start for longer inc subseq
		} else {
			// replace the smallest element >= a[i]
			// feels greedy!
			m := binarySearch2(k, length, a, t)
			t[m] = i
			// the invariant: a[t[0]...t[m-1]] < a[t[m]]
			// hence set prev[i] = t[m-1]
			prev[i] = t[m-1]
		}
	}

	ret := make([]int, length)
	j := t[length-1]
	for length > 0 {
		ret[length-1] = a[j]
		j = prev[j]
		length--
	}
	return ret
}

// since we use indices in t
// to get the value, we need a[]
func binarySearch2(k, length int, a, t []int) int {
	l, r := 0, length
	for l < r {
		m := l + ((r - l) >> 1)
		if cmp := a[t[m]]; cmp < k {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

// find the position of the smallest element
// that is greater or equal k
func binarySearch(k, length int, a []int) int {
	l, r := 0, length
	for l != r {
		m := l + (r-l)/2

		if a[m] < k {
			// if a[m] < k, then a[0...m] < k -> a[0...m] are of no use
			// start l at m+1
			l = m + 1
		} else {
			// a[m] >= k, m might be a candidate
			// but we are looking for the "first" element larger than k
			// so a[m+1...n-1] are useless -> start r at m
			r = m
		}
	}
	return l
}

func main() {
	// a := []int{7, 7, 8}
	// fmt.Println(lis2_construction(a))
	fmt.Println(lis1([]int{0, 1, 0, 3, 2, 3}))
	fmt.Println(lis1([]int{10, 9, 2, 5, 3, 4}))
}

// some example to think about
// a = [4, 5, 1, 2, 3] -> lis = [1, 2, 3]
// a = [4, 5, 11, 7, 8] -> lis = [4, 5, 7, 8]

// In lis2_construct
// t doesn't necessarily stores the actual lis
// for example
// a = [4, 5, 1] -> t = [1, 5] != lis = [4, 5]
// or
// a = [4, 5, 6, 1, 2] -> t = [1, 2, 6]
