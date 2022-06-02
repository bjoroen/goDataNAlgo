package linkedlist

import (
	U "goAlgoData.com/utils"
	"testing"
)

func TestPrepend(t *testing.T) {

	list := linkedList{}
	node1 := &Node{data: 48}
	node2 := &Node{data: 30}
	node3 := &Node{data: 123}

	list.prepend(node1)
	list.prepend(node2)
	list.prepend(node3)

	want := 3
	got := list.length

	if got != want {
		t.Errorf("we got %q, but wanted %q", got, want)

	}
}

func TestPrintListData(t *testing.T) {
	list := linkedList{}

	node1 := &Node{data: 48}
	node2 := &Node{data: 30}
	node3 := &Node{data: 123}

	list.prepend(node1)
	list.prepend(node2)
	list.prepend(node3)

	want := true

	got := U.Contains(list.printListData(), 123)

	if got != want {
		t.Errorf("we got %t, but wanted %t", got, want)
	}

}

func TestDeleteWithValue(t *testing.T) {

	list := linkedList{}
	node1 := &Node{data: 48}
	node2 := &Node{data: 18}
	node3 := &Node{data: 26}
	node4 := &Node{data: 118}

	list.prepend(node1)
	list.prepend(node2)
	list.prepend(node3)
	list.prepend(node4)
	

	list.delteWithValue(26)

	want := 3
	got := list.length

	if got != want {
		t.Errorf("wanted %q got %q", want, got )
	}
	

	 


}
