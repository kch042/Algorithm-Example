package wiggle

// leetcode #280
// Input: []int
// requirement: a[0] <= a[1] >= a[2] <= a[3] >= ....

func wiggle(a []int) {
	for i := 1; i < len(a); i++ {
		if (i%2 == 1 && a[i] <= a[i-1]) || (i%2 == 0 && a[i] >= a[i-1]) {
			a[i], a[i-1] = a[i-1], a[i]
		}
	}
}
