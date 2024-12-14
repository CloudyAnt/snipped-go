package algorithm

// For left to right, select the correct one from rest on right to current
func selection(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

// Sink the heaviest to the right most, then continue to bubble the rest
func bubble(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// From left to right, keep the left part ordered by bubble current
func insertion(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}

// For the rightest, move all the smaller ones to left, move itself to the index after them.
// Then, for each of the left and right part, do the same thing again
func qs(arr []int) {
	qs0(arr, 0, len(arr)-1)
}

func qs0(arr []int, l, r int) {
	if l < r {
		i := partition(arr, l, r)
		qs0(arr, l, i-1)
		qs0(arr, i+1, r)
	}
}

func partition(arr []int, l, r int) int {
	i := l - 1
	for j := l; j < r; j++ {
		if arr[j] < arr[r] {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	i++
	arr[r], arr[i] = arr[i], arr[r]
	return i
}

// Build a heap, then repeat length times: swap the top to bottom, re-heapify
func heap(arr []int) {
	mid := len(arr)/2 - 1
	for i := mid; i >= 0; i-- {
		heapify(arr, i, len(arr)) // this will put the biggest to the top
	}

	for i := len(arr) - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		heapify(arr, 0, i)
	}
}

func heapify(arr []int, i, limit int) {
	maxIndex := i
	l := i*2 + 1
	if l < limit && arr[l] > arr[maxIndex] {
		maxIndex = l
	}

	r := i*2 + 2
	if r < limit && arr[r] > arr[maxIndex] {
		maxIndex = r
	}

	if maxIndex != i {
		arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
		heapify(arr, maxIndex, limit)
	}
}
