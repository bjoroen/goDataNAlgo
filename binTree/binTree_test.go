package bintree

import (
	"testing"
)

func TestConstructor(t *testing.T) {

	tree := &Node{Key: 100}
	got := tree.Key
	want := 100

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}

func TestInsert(t *testing.T) {

	tree := &Node{Key: 100}
	tree.Insert(200)

	got := tree.right.Key
	want := 200

	if got != want {
		t.Errorf("got %q, want %q", got, want)

	}
}

func TestSearch(t *testing.T){
	tree := &Node{Key: 100}
	tree.Insert(50)

	got := tree.Search(50)
	want := true

	if got != want {
		t.Errorf("got %t, want %t", got, want)
		
	}
}
