package gosort

// GO sort pkg implementation
// simpler ver.
// applies to []int only

// Mysort is the only exported func for now
func Mysort(a []int) {

}

func median(a []int, m1, m0, m2 int) {
	if a[m0] > a[m1] {
		a[m0], a[m1] = a[m1], a[m0]
	} // now a[m0] <= a[m1]

	if a[m1] > a[m2] {
		a[m1], a[m2] = a[m2], a[m1]
		if a[m1] < a[m0] {
			a[m0], a[m1] = a[m1], a[m0]
		}
	}
}

func partition(a []int, lo, hi int) (midlo, midhi int) {
	m := int(uint(lo+hi) >> 1)
	if hi-lo > 40 {
		p := (hi - 1 - lo) / 8
		median(a, lo, lo+p, lo+2*p)
		median(a, m, m-p, m+p)
		median(a, hi-1, hi-1-p, hi-1-2*p)
	}
	median(a, lo, m, hi-1)

	pivot := a[lo]
	midlo, midhi = lo, lo
	for j := lo + 1; j < hi; j++ {
		if a[j] < pivot {
			midhi, midlo = midhi+1, midlo+1
			a[j], a[midhi] = a[midhi], a[j]
			a[midlo], a[midhi] = a[midhi], a[midlo]
		} else if a[j] == pivot {
			midhi++
			a[j], a[midhi] = a[midhi], a[j]
		}
	}
	a[lo], a[midlo] = a[midlo], a[lo]
	return
}

func quicksort(a []int, l, r int) {
	// quicksort
	if r-l > 12 {
		ml, mh := partition(a, l, r)

		if ml-l < r-mh { // left side smaller, solve left side first
			quicksort(a, l, ml)
			l = mh
		} else {
			quicksort(a, mh, r)
			r = ml
		}
	}

	// shell sort
	for i := l; i+6 < r; i++ {
		if a[i] > a[i+6] {
			a[i], a[i+6] = a[i+6], a[i]
		}
		insertionsort(a, l, r)
	}

}

// insertion sort a[lo:hi]
func insertionsort(a []int, lo, hi int) {
	i := lo + 1
	for i < hi {
		for j := i; j > 0 && a[j] < a[j-1]; j-- {
			a[j], a[j-1] = a[j-1], a[j]
		}
		i++
	}
}

func siftDown(a []int, lo, hi, offset int) {
	now := lo
	for 2*now+1 < hi {
		child := 2*now + 1
		if child+1 < hi && a[child+offset] < a[child+1+offset] {
			child++
		}

		if a[child+offset] <= a[now+offset] {
			return
		}

		a[child+offset], a[now+offset] = a[now+offset], a[child+offset]
		now = child
	}
}

// create the var offset to transform the
// "absolute" position(p to q) to "relative" position(0 to q-p)
// in order to apply heap sort
func heapSort(a []int, p, q int) {
	offset := p
	lo, hi := 0, q-p

	for i := (hi - 1) / 2; i >= lo; i-- {
		siftDown(a, i, hi, offset)
	}

	for j := hi - 1; j > 0; j-- {
		a[lo+offset], a[j+offset] = a[j+offset], a[lo+offset]
		siftDown(a, lo, j, offset)
	}
}
