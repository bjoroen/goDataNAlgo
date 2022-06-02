package heaps

import (
	"fmt"
	"testing"
)

func TestHeapCreat(t *testing.T) {

	m := &MaxHeap{}

	buildHeap := []int{8, 12, 10, 20}

	for _, v := range buildHeap {
		m.Insert(v)
	}

	want := 4
	got := len(m.array)

	if got != want {
		t.Errorf("nhaa bitch")
	}
}

func TestHeapExtract(t *testing.T) {
	m := &MaxHeap{}
	buildHeap := []int{13, 2, 50, 100, 52, 23, 54, 23, 654, 32, 53}

	for _, v := range buildHeap {
		m.Insert(v)
	}

	want := 11
	got := len(m.array)

	if got != want {
		t.Errorf("Not correct bitch")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(m)
		m.Extract()
	}

	want = 6
	got = len(m.array)

	if got != want {
		t.Errorf("Gotthca bitch")
	}
}
