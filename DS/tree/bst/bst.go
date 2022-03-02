package main

import "fmt"

type bst struct {
	root *node
	len  int
}

type node struct {
	pa          *node
	left, right *node
	val         int
}

func (b *bst) search(k int) *node {
	now := b.root
	for now != nil {
		if k == now.val {
			return now
		} else if k > now.val {
			now = now.right
		} else {
			now = now.left
		}
	}
	return nil
}

// ref: CLRS P298 chap12
func (b *bst) delete(k int) {
	x := b.search(k)
	if x.left == nil {
		b.transplant(x, x.right)
	} else if x.right == nil {
		b.transplant(x, x.left)
	} else { // x has two children
		y := x.successor()
		if y.pa != x {
			b.transplant(y, y.right)
			y.right = x.right
			y.right.pa = y
		}
		b.transplant(x, y)
		y.left = x.left
		y.left.pa = y
	}
}

// subroutine for delete
func (b *bst) transplant(u, v *node) {
	// replace u with v
	if u.pa == nil {
		b.root = v
	} else if u == u.pa.left {
		u.pa.left = v
	} else {
		u.pa.right = v
	}
	if v != nil {
		v.pa = u.pa
	}
}

func (root *node) treeMin() *node {
	if root == nil {
		return nil
	}
	x := root
	for x.left != nil {
		x = x.left
	}
	return x
}

func (root *node) treeMax() *node {
	if root == nil {
		return nil
	}
	x := root
	for x.right != nil {
		x = x.right
	}
	return x
}

func (x *node) predecessor() *node {
	if x.left != nil {
		return x.left.treeMax()
	}
	y := x.pa
	for y != nil && x == y.left {
		x, y = y, y.pa
	}
	return y
}

func (x *node) successor() *node {
	if x.right != nil {
		return x.right.treeMin()
	}
	y := x.pa
	for y != nil && x == y.right {
		x = y
		y = y.pa
	}
	return y
}

// Assume all keys are distinct
func (b *bst) insert(k int) {
	b.len++
	if b.root == nil {
		b.root = &node{val: k}
		return
	}
	b.root.insert(k)
}

func (root *node) insert(k int) {
	if k < root.val {
		if root.left == nil {
			root.left = &node{val: k, pa: root}
			return
		}
		root.left.insert(k)
	} else {
		if root.right == nil {
			root.right = &node{val: k, pa: root}
			return
		}
		root.right.insert(k)
	}
}

// iterative method to traverse in order with stack
func (b *bst) inorder() {
	stack := make([]*node, b.len)
	top := -1
	cur := b.root
	for cur != nil || top > -1 {
		if cur != nil { // push
			top++
			stack[top] = cur
			cur = cur.left
		} else { // pop
			pop := stack[top]
			top--
			fmt.Println(pop.val)
			cur = pop.right
		}
	}
}

func inorderRec(root *node) {
	if root == nil {
		return
	}
	inorderRec(root.left)
	fmt.Println(root.val)
	inorderRec(root.right)
}

func main() {
	b := &bst{}
	in := []int{3, 1, 8, 2, 6, 7, 5}
	// in := []int{2, 1, 4, 3}
	for _, v := range in {
		b.insert(v)
	}
	// b.delete(1)  // test whether delete is commutative
	// b.delete(2)
	// fmt.Println(b.root.val)

	// inorderRec(b.root)
	// b.inorder()

	b.delete(5)
	for i := 1; i < 9; i++ {
		tmp := b.search(i)
		if tmp == nil {
			fmt.Printf("The value %d is not found!\n", i)
		} else {
			if suc := tmp.successor(); suc != nil {
				fmt.Printf("The suc of %d is %d\n", i, tmp.successor().val)
			} else {
				fmt.Printf("The value %d has no suc!\n", i)
			}
		}
	}
}
