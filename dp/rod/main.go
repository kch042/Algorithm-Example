package main

import "fmt"

func main() {
	p := []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	for i := 1; i <= 10; i++ {
		rodCutting(i, p)
	}
}

func rodCutting(n int, p []int) {
	dp := make([]int, n+1)
	s := make([]int, n+1)

	for j := 1; j <= n; j++ {
		for i := 1; i <= j; i++ {
			if dp[j] < p[i]+dp[j-i] {
				dp[j] = p[i] + dp[j-i]
				s[j] = i
			}
		}
	}
	fmt.Printf("Max Profit of len %d: %d\n", n, dp[n])
	printRodCutting(n, s)
}

func printRodCutting(n int, s []int) {
	fmt.Println("Cut into pieces of len:")
	for n > 0 {
		fmt.Printf("%d ", s[n])
		n = n - s[n]
	}
	fmt.Println()
}
