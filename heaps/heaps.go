package heaps

type MaxHeap struct {
	array  []int
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)

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
	return 2*i + 2
}

//Swap keys in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}
