package heapsort

import (
	"fmt"
	"sort"
	"testing"
)

func TestHsort(t *testing.T) {
	tcs := make([][]int, 10)
	for i := range tcs {
		tcs[i] = generate(1000000)
	}
	for ind, tc := range tcs {
		t.Run(fmt.Sprintf("%d", ind), func(t *testing.T) {
			if !sort.IntsAreSorted(hsort(tc)) {
				t.Error("Something wrong!")
			}
		})
	}
}
