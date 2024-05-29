package binary_tree

import "golang.org/x/exp/constraints"

type Node[T constraints.Ordered] struct {
	value T
	left  *Node[T]
	right *Node[T]
}

type BinaryTree[T constraints.Ordered] struct {
	node *Node[T]
}

func (t *BinaryTree[T]) Insert(value T) {
	n := Node[T]{
		value: value,
	}
	if t.node == nil {
		t.node = &n
		return
	}
	t.node.Insert(&n)
}

func (n *Node[T]) Insert(nn *Node[T]) {
	if nn.value < n.value {
		if n.left == nil {
			n.left = nn
			return
		}
		n.left.Insert(nn)
	} else {
		if n.right == nil {
			n.right = nn
			return
		}
		n.right.Insert(nn)
	}
}

// Remove the first occurrence of a value
func (t *BinaryTree[T]) Remove(value T) {

}
