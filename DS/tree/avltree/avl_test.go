package avl

func Examplepreorder() {
	seed := []int{2, 5, 1, 8, 9, 10, 12, 3}
	var root *node
	for _, s := range seed {
		root = insert(s, root)
	}
	preorder(root)
	// Output:
	// val: 1, Height: 0
	// val: 3, Height: 0
	// val: 5, Height: 1
	// val: 2, Height: 2
	// val: 9, Height: 0
	// val: 12, Height: 0
	// val: 10, Height: 1
	// val: 8, Height: 3
}
