package main

import "testing"

func TestScs(t *testing.T) {
	input := []struct {
		a   string
		b   string
		ans int
	}{
		{"AXA", "XB", 4},
		{"geek", "eke", 5},
		{"AGGTAB", "GXTXAYB", 9}, // "AGXGTXAYB"
		{"ABCBDAB", "BDCABA", 9}, // ABCBDCABA, ABDCABDAB, and ABDCBDABA
		{"abcbdab", "bdcaba", 9}, // abdcabdab
	}

	for _, in := range input {
		t.Run("scs", func(t *testing.T) {
			if res, check := scs(in.a, in.b), checkSCS(in.a, in.b); res != check {
				t.Errorf("Get: %d; Want: %d\n", res, check)
			}
		})
	}
}
