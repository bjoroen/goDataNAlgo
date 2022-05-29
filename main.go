package main

import (
	"fmt"
	"math/rand"

	B "goAlgoData.com/binTree"
)

func main() {

	tree := &B.Node{Key: 100}

	for i := 1; i < 200; i++ {
		min := 1
		max := 2000

		tree.Insert(rand.Intn(max-min+1) + min)
	}
	tree.Insert(20)

	fmt.Println(tree)


  fmt.Println(tree.Search(20))

}
