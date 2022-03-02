package main

import (
	"reflect"
	"testing"
)

// usage: "go test -v"

var tcs [][]int = [][]int{
	{4, 5, 1, 2, 3},
	{10, 9, 2, 5, 3, 7, 101, 18},
	{7, 7, 7, 7, 7, 7},
	{0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15},
	{5, 8, 3, 7, 9, 1},
	{0, 1, 0, 3, 2, 3},
}

var ans = [][]int{
	{1, 2, 3},
	{2, 3, 7, 18},
	{7},
	{0, 2, 6, 9, 11, 15},
	{3, 7, 9},
	{0, 1, 2, 3},
}

func TestLis2(t *testing.T) {
	for i := range tcs {
		res := lis2(tcs[i])
		if res != len(ans[i]) {
			t.Errorf("Want: %d\nGet: %d", len(ans[i]), res)
		}
	}
}

func TestLis2_construction(t *testing.T) {
	for i, tc := range tcs {
		t.Run("123", func(t *testing.T) {
			res := lis2_construction(tc)
			if !reflect.DeepEqual(res, ans[i]) { // compare slices
				t.Error("want:", ans[i], "\nget: ", res)
			}
		})
	}
}
