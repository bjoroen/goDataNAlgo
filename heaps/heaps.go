package heaps

import "fmt"

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)

}

func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	lastElement := len(h.array) - 1

	//if 0 return
	if len(h.array) == 0 {
		fmt.Println("Cant extract because array length is 0")
		return -1
	}

	h.array[0] = h.array[lastElement]
	h.array = h.array[:lastElement]

	return extracted
}

// maxHeapify Down, will heapify from the top to the bottom
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	//Loop while index has at least one child
	for l <= lastIndex {

		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)

		} else {
			return
		}
	}

}

// maxHeapifyUp will heapify from bottom to the top
func (h *MaxHeap) maxHeapifyUp(index int) {

	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

//Get the Parent index
func parent(i int) int {
	return (i - 1) / 2
}

//Get the left child index
func left(i int) int {
	return 2*i + 1
}

//Get the right child index
func right(i int) int {
	return 2*i + 1
}

//Swap keys in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}
