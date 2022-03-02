package rb

const (
	red uint8 = iota
	black
)

type node struct {
	val   int
	color uint8
	pa    *node
	left  *node
	right *node
}

func (n *node) un() *node {
	gp := n.pa.pa
	if n.pa != gp.left {
		return gp.left
	}
	return gp.right
}

func insert(val int, root *node) *node {
	n := &node{
		color: red,
		val:   val,
	}
	if root == nil {
		return insert1(n)
	}

	pa := root
	for root != nil {
		pa = root
		if n.val > pa.val {
			root = root.right
		} else {
			root = root.left
		}
	}
	n.pa = pa
	return insert1(n)
}

func insert1(n *node) *node {
	if n.pa == nil {
		n.color = black
		return n
	}
	if n.pa.color == red {
		return insert2(n)
	}

	return rt(n)
}

func insert2(n *node) *node {
	if un := n.un(); un.color == red {
		un.color, n.pa.color = black, black
		gp := n.pa.pa
		gp.color = red
		return insert1(gp)
	}
	return insert3(n)
}

// uncle is black
func insert3(n *node) *node {
	pa := n.pa
	if n == pa.right && pa == pa.pa.left {
		rotateLeft(n)
		n = n.left
	} else if n == pa.left && pa == pa.pa.right {
		rotateRight(n)
		n = n.right
	}
	return insert4(n)
}

func insert4(n *node) *node {
	pa, gp := n.pa, n.pa.pa
	pa.color, gp.color = black, red
	if pa == gp.left { // left left case
		rotateRight(pa)
	} else {
		rotateLeft(pa)
	}
	return rt(pa)
}

func rotateRight(n *node) {
	if n == nil || n.pa == nil {
		return
	}
	gp, fa, y := n.pa.pa, n.pa, n.right

	fa.left, y.pa = y, fa  // 1. n's parent "fosters" the right child of n as its left child
	n.right, fa.pa = fa, n // 2. n switch roles with its parent
	n.pa = gp              // 3. n connect with grandpa
	if gp.left == fa {
		gp.left = n
	} else {
		gp.right = n
	}
}

func rotateLeft(n *node) {
	if n == nil || n.pa == nil {
		return
	}
	gp, fa, x := n.pa.pa, n.pa, n.left

	fa.right, x.pa = x, fa
	n.left, fa.pa = fa, n
	n.pa = gp
	if gp.left == fa {
		gp.left = n
	} else {
		gp.right = n
	}
}

func rt(n *node) *node {
	for n.pa != nil {
		n = n.pa
	}
	return n
}

// rbtree is a self-balanced BST

// 5 Properties(Rules) of rbtree
// 1. Each node is either black or red
// 2. Root is black
// 3. The children of a red node are all black (No contiguous red nodes are allowed)
// 4. All leafs are black (including nil)
// 5. The path of a node to each of its leaf has the same number of black nodes

// Insertion for rbtree
// 1. Insert the new node n the same way as BST
// 2. Color the new node red

// 3. func insert1(n):
//      if n.pa == nil:
//        color n black
//      elif n.pa is red:
//        insert2(n)

// 4. func insert2(n):
//      if n.uncle is red:
//        color n.uncle and n.pa black
//        color n.gp red
//        insert1(n.pa)
//      else:
//        insert3(n)

// 5. func insert3(n):
//		if left-right case:
//		  rotateLeft(n)
//		  n = n.left
//		  insert4(n)
//      elif right-letf case:
//		  rotateRight(n)
//		  n = n.right
//		  insert(n)

// 6. func insert4(n):
//		recolor n.pa and n.gp
//		if left-left case:
//		  rotateRight(n.pa)
//		else:
//		  rotateLeft(n.pa)
