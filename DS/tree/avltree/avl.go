package avl

import "fmt"

// REVIEW AVL Tree
// 	Given a node n
// 	1. height: max(n.left.height, n.right.height) + 1
//     nil child has height -1
// 	2. balance factor(bf): n.right.height - n.left.height
//     valid bf = -1 or 0 or 1

// README
// Simple implementation of avl tree
//
// Whenever we want to insert a node
// We traverse the tree to find the valid position to insert
// during which we push the nodes we visit into a stack
// for the reason explained after
//
// After we have inserted the node and set its height 0 (since it's a leaf node)
// we traverse back up to the root node
// during which we do two things:
//   for each node we visited:
//     1. update the height of that node
//        (since the lower nodes have the correct height,
//        we can easily calculate the height by its two children node's height)
//     2. if the node is not balanced:
//          do the rotation (LR, LL, RL, RR cases)
//     return the root node of that subtree after balance
// How do we do that?
//   Remember that when we traverse down the tree, we had made a stack
//   Now we just need to pop the stack as we traverse back up
//
// To better understand
// 1. Draw avl tree
// 2. see the ref pseudocode

type node struct {
	val    int
	left   *node
	right  *node
	height int
}

// test function
func preorder(n *node) {
	if n == nil {
		return
	}
	preorder(n.left)
	preorder(n.right)
	fmt.Printf("val: %d, Height: %d\n", n.val, n.height)
}

func insert(i int, root *node) *node {
	newNode := &node{val: i}
	if root == nil {
		return &node{val: i}
	}

	// update height, also check and balance the tree in a bottom-up manner
	stack := ins(newNode, root)
	for i := len(stack) - 1; i >= 0; i-- {
		now := stack[i]
		if bNode := balance(now); bNode != now { // means that node is inbalanced and thus rotated
			if i == 0 { // root node is rotated and some node is the new root
				root = bNode
			}
			pa := stack[i-1]
			if now == pa.left {
				pa.left = bNode
			} else {
				pa.right = bNode
			}
		}
	}
	return root
}

// balance returns the new root node after balance
// ref: https://www.cs.swarthmore.edu/~brody/cs35/f14/Labs/extras/08/avl_pseudo.pdf
func balance(n *node) *node {
	if n == nil {
		return nil // no need to balance, done
	}

	// update the height of current node
	n.height = height(n)

	lh, rh := lrh(n)
	bf := rh - lh
	if bf > 1 { // right-heavy
		if llh, rrh := lrh(n.right); rrh >= llh { // RR case -> left rotation
			return leftRotate(n)
		}
		// RL case
		return rightLeftRotate(n)

	} else if bf < -1 {
		if llh, rrh := lrh(n.left); llh >= rrh { // LL case -> left rotation
			return rightRotate(n)
		}
		// LR case
		return leftRightRotate(n)
	}

	return n // at this point, the node didn't rotate
}

// returns the new root of the subtree after rotation
func rightRotate(n *node) *node {
	ret := n.left
	n.left = ret.right
	ret.right = n

	// update height of two nodes: n and ret
	// note that the order of updating matters: n first, then ret (bottom-up manner)
	n.height = height(n)
	ret.height = height(ret)
	return ret
}

func leftRotate(n *node) *node {
	ret := n.right
	n.right = ret.left
	ret.left = n

	n.height = height(n)
	ret.height = height(ret)
	return ret
}

func rightLeftRotate(n *node) *node {
	n.right = rightRotate(n.right)
	return leftRotate(n)
}

func leftRightRotate(n *node) *node {
	n.left = leftRotate(n.left)
	return rightLeftRotate(n)
}

func height(n *node) int {
	return max(lrh(n)) + 1
}

// lrh returns the height of left and right subtrees
func lrh(n *node) (int, int) {
	lh, rh := -1, -1
	if n.left != nil {
		lh = n.left.height
	}
	if n.right != nil {
		rh = n.right.height
	}
	return lh, rh
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ins is the normal bst insert func
// ins returns a stack of nodes it traverse through
func ins(newNode, root *node) []*node {
	var stack []*node
	cur := root
	for cur != nil {
		stack = append(stack, cur)
		if cur.val < newNode.val {
			if cur.right == nil {
				cur.right = newNode
				cur = nil // break
			} else {
				cur = cur.right
			}
		} else {
			if cur.left == nil {
				cur.left = newNode
				cur = nil // break
			} else {
				cur = cur.left
			}
		}
	}
	return stack
}
