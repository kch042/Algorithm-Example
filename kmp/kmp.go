package main

import "fmt"

func main() {
	T := "applepenapple"
	P := "apple"
	// P := "AAACAAAA"
	// fmt.Println(longestPrefix(P))
	fmt.Println(kmp(T, P))
}

func kmp(T, P string) []int {
	ret := make([]int, 0)
	p := longestPrefix(P)

	k := 0
	for x := range T {
		for k > 0 && P[k] != T[x] {
			k = p[k-1]
		}
		if P[k] == T[x] {
			k++
		}
		if k == len(P) {
			ret = append(ret, x-len(P)+1)
			k = p[k-1]
		}
	}

	return ret
}

func longestPrefix(P string) []int {
	p := make([]int, len(P))
	p[0] = 0

	i, k := 1, 0
	for i < len(P) {
		for k != 0 && P[i] != P[k] {
			// this is the most tricky part
			k = p[k-1]
		}
		if P[i] == P[k] {
			k++
		}
		p[i] = k
		i++
	}

	return p
}
