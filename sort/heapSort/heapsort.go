package heapsort

import (
	"math/rand"
	"time"
)

func hsort(nums []int) []int {
	// build a max heap
	// which is O(n)
	n := len(nums)
	for i := n/2 - 1; i >= 0; i-- { // heapify bottom-up
		heapify(n, i, nums)
	}

	// i represents the boundary of the heap
	// when the heap reaches 0, the sort is complete
	// time complexity O(n*logn)
	// since we run a loop with O(n)
	// in each loop we perform a heapify with O(logn)
	for i := n - 1; i > 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		heapify(i, 0, nums) // heapify top-down
	}

	return nums
}

// heapify at i-th position given an array of size n
// More generally, it's a "sift-down" operation
// time complexity for one heapify: O(logn)
func heapify(n int, i int, nums []int) {
	if i >= n/2 {
		return // leaf node
	}

	lar := i
	if l := 2*i + 1; l < n && nums[l] > nums[lar] {
		lar = l
	}
	if r := 2*i + 2; r < n && nums[r] > nums[lar] {
		lar = r
	}
	if lar != i {
		nums[i], nums[lar] = nums[lar], nums[i]
		heapify(n, lar, nums)
	}
}

func siftDown(a []int, root, hi int) {
	// 1. two children duel first (if two exist)
	// 2. challenge the parent
	//    if suc: swap and sift down
	//    else : return
	for 2*root+1 < hi {
		child := 2*root + 1
		if 2*root+2 < hi && a[child] < a[2*root+2] {
			child = 2*root + 2
		}
		if a[root] < a[child] {
			a[root], a[child] = a[child], a[root]
		} else {
			return
		}
		root = child
	}
}

func generate(n int) []int {
	rand.Seed(time.Now().Unix())
	return rand.Perm(n)
}
