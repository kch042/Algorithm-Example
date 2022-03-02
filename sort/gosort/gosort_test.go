package gosort

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

// how many numbers in each sorting test case
const testNum = 100000

func Test_insertionsort(t *testing.T) {
	tsort(t, insertionsort, 0, testNum)
}

func Test_heapSort(t *testing.T) {
	tsort(t, heapSort, 10, 500)
}

func Test_quicksort(t *testing.T) {
	tsort(t, quicksort, 0, testNum)
}

// helper function for testing sort functions
func tsort(t *testing.T, fn func([]int, int, int), lo, hi int) {
	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		slice []int
		name  string
	}{
		{rand.Perm(testNum), "case 1"},
		{rand.Perm(testNum), "case 2"},
		{rand.Perm(testNum), "case 3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn(tt.slice, lo, hi)
			if !sort.IntsAreSorted(tt.slice[lo:hi]) {
				t.Errorf("Failed: %s", tt.name)
			}
		})
	}
}
