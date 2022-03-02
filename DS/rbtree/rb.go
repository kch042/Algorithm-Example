package main

type node struct {
	val   int
	left  *node
	right *node
	pa    *node
	black bool
}

type RBTree struct {
	root *node
}

func (rb *RBTree) insert(k int) {
	if rb.root == nil {
		rb.root = &node{val: k, black: true}
		return
	}
	rb.root.insert(k)
	rb.root.black = true
}

func (root *node) insert(k int) {
	x := root.bstInsert(k)
	if x.pa.black {
		return
	}

	for x != root && !x.black {
		if x.pa == x.pa.pa.left { // A: the pa of x is the left child of grandpa of x
			y := x.pa.pa.right // uncle
			if !y.black {
				// case 1: pa and uncle are both red
				x.pa.black, y.black, x.pa.pa.black = true, true, false
				x = x.pa.pa
			} else {
				// case 2: zigzag shape of x, x.pa, x.pa.pa
				if x == x.pa.right {
					x = x.pa
					x.leftRotate()
				}
				// case 3: either occurs singly or after case 2
				x = x.pa.pa.rightRotate()
				x.black, x.right.black = true, false
				return
			}
		} else { // B: the pa of x is the right child of grandpa of x
			y := x.pa.pa.left // uncle
			if !y.black {     // case 1
				x.pa.black, y.black, x.pa.pa.black = true, true, false
				x = x.pa.pa
			} else {
				// case 2
				if x == x.pa.left {
					x = x.pa
					x.rightRotate()
				}
				// case 3
				x = x.pa.pa.leftRotate()
				x.black, x.left.black = true, false
			}
		}
	}
}

func (root *node) bstInsert(k int) *node {
	if root.val < k {
		if root.left == nil {
			root.left = &node{val: k, pa: root, black: false}
			return root.left
		}
		return root.left.bstInsert(k)
	} else {
		if root.right == nil {
			root.right = &node{pa: root, val: k, black: false}
			return root.right
		}
		return root.right.bstInsert(k)
	}
}

func (root *node) leftRotate() *node {
	if root.right == nil {
		return root
	}

	x, y, g := root, root.right, root.pa
	if g != nil {
		if x == g.left {
			y.pa, g.left = g.left, y
		} else {
			y.pa, g.right = g, y
		}
	}
	y.left.pa, x.right = x, y.left
	y.left, x.pa = x, y

	return y
}

func (root *node) rightRotate() *node {
	if root.left == nil {
		return root
	}
	x, y, g := root, root.left, root.pa
	if g != nil {
		if x == g.right {
			g.right, y.pa = y, g
		} else {
			g.left, y.pa = y, g
		}
	}
	y.right.pa, x.left = x, y.right
	x.pa, y.right = y, x

	return y
}

func (rb *RBTree) search(val int) *node {
	cur := rb.root
	for cur != nil {
		if val == cur.val {
			return cur
		} else if val > cur.val {
			cur = cur.right
		} else {
			cur = cur.left
		}
	}
	return nil
}
