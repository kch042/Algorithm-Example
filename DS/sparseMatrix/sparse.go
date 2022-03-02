package main

type sparseMatrix struct {
	first *node
}

type node struct {
	val  int
	next *node
	row  *element
	col  *element
}

type element struct {
	r, c, val int
	down      *element
	right     *element
}

func newSM() sparseMatrix {
	ret := sparseMatrix{}
	ret.first.next = ret.first
	return ret
}

// checks if the sm[r][c] has entry
func (sm *sparseMatrix) elementExist(r, c int) (bool, *element) {
	for now := sm.first; now.next != sm.first; now = now.next {
		if now.next.val == r {
			for cnow := now.next.col; cnow.right != now.next.col; cnow = cnow.right {
				if cnow.right.c > c {
					return false, nil
				} else if cnow.right.c == c {
					return true, cnow.right
				}
			}
		} else if now.next.val > r {
			break
		}
	}
	return false, nil
}

func (sm *sparseMatrix) add(r, c, val int) {
	if ex, ptr := sm.elementExist(r, c); ex {
		ptr.val += val
	} else {
		for now := sm.first; now.next != sm.first; now = now.next {

		}
	}
}

func (sm *sparseMatrix) printCol(c int) {

}

func (sm *sparseMatrix) printRow(r int) {

}

func main() {

}
