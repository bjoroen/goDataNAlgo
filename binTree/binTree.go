package bintree

//Node
type Node struct {
	Key   int
	left  *Node
	right *Node
}


//Insert
func (n *Node) Insert(k int) {

	if n.Key < k {
		//move right
		if n.right == nil {
			n.right = &Node{Key: k}
		}
		n.right.Insert(k)
	}
	if n.Key > k {
		//move left
		if n.left == nil {
			n.left = &Node{Key: k}
		}
		n.left.Insert(k)
	}
}

//Search

func (n *Node) Search(k int) bool {
  

	if n == nil {
		return false
	}

	if n.Key < k {
		return n.right.Search(k)
	}

	if n.Key > k {
		return n.left.Search(k)
	}

	return true

}
