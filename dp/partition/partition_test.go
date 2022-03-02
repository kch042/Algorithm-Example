package main

import "testing"

func TestPartition(t *testing.T) {
	input := []struct {
		nums []int
		ans  bool
	}{
		{[]int{1, 5, 11, 5}, true},
		{[]int{1, 5, 3}, false},
		{[]int{3, 1, 1, 2, 2, 1}, true},
		{[]int{1000, 5, 994, 0}, false},
	}

	for _, in := range input {
		t.Run("partition", func(t *testing.T) {
			if res := partition(in.nums); res != in.ans {
				t.Fail()
			}
		})
	}
}
