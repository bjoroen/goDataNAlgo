package linkedlist

import "fmt"

type Node struct {
	data int
	next *Node
}

type linkedList struct {
	head   *Node
	length int
}

func (l *linkedList) prepend(n *Node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
	fmt.Println("\nAdded value: ", l.head.data)
}

func (l linkedList) printListData() []int {
	toPrint := l.head
	var slic []int

	for l.length != 0 {
		fmt.Printf("\nValue in List %d ", toPrint.data)
		slic = append(slic, toPrint.data)
		toPrint = toPrint.next
		l.length--
	}

	return slic
}

func (l *linkedList) delteWithValue(value int) {

	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	toDelete := l.head

	for toDelete.next.data != value {
		if toDelete.next.next == nil {
			return
		}
		toDelete = toDelete.next
	}
	fmt.Println("\nDeleted value",toDelete.next.data)
	toDelete.next = toDelete.next.next
	l.length--

}
 
