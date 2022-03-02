package main

import "fmt"

// Given a sequence of numbers called A
// A bitonic subsequence of A is such that
// it first increases and then decreases

// For example:
// A = [1, 11, 2, 10, 4, 5, 2, 1]
// then bitonic = [1, 2, 10, 5, 2, 1]

// another example:
// A = [8, 6, 5, 4, 11, 0]
// then bitonic = [8, 6, 5, 4, 0]
// [8] is trivially an increasing sequence

// To find the length of the bitonic subsequence
// we need to calculate
// 1. LIS[i]: the length of longest increasing subsequence
//			  of A "ending" at A[i]
// 2. LDS[i]: the length of longest decreasing subsequence
//			  of A "starting" at A[i]
// then the length of longest bitonic subseq =
// 			max{LIS[i] + LDS[i] - 1, for all i s.t. 0 <= i < n}

func bitonic(A []int) int {
	n := len(A)

	LIS, LDS := make([]int, n), make([]int, n)

	// base case
	// Every number is a longest inc/dec subsequence of length 1
	for i := range LIS {
		LIS[i], LDS[i] = 1, 1
	}

	// LIS
	for i := range LIS {
		for j := 0; j < i; j++ {
			if A[i] > A[j] && LIS[j]+1 > LIS[i] {
				LIS[i] = LIS[j] + 1
			}
		}
	}

	// LDS
	for i := n - 2; i >= 0; i-- {
		for j := n - 1; j > i; j-- {
			if A[i] > A[j] && LDS[i] < LDS[j]+1 {
				LDS[i] = LDS[j] + 1
			}
		}
	}

	ret := LIS[0] + LDS[0] - 1
	for i := range A {
		if ret < LIS[i]+LDS[i]-1 {
			ret = LIS[i] + LDS[i] - 1
		}
	}

	return ret
}

func main() {
	// A := []int{1, 11, 2, 10, 4, 5, 2, 1}
	A := []int{80, 60, 30, 40, 20, 10}
	B := []int{12, 11, 40, 5, 3, 1}
	fmt.Println(bitonic(A))
	fmt.Println(bitonic(B))
}
