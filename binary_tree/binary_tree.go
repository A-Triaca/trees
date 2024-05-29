package binary_tree

import (
	"golang.org/x/exp/constraints"
)

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
	if t.node == nil {
		return
	}
	t.node = t.node.Remove(value)
}

func (n *Node[T]) Remove(value T) *Node[T] {
	if n == nil {
		return n
	}

	if value < n.value {
		n.left = n.left.Remove(value)
	} else if value > n.value {
		n.right = n.right.Remove(value)
	} else {
		if n.left == nil {
			return n.right
		} else if n.right == nil {
			return n.left
		}

		n.value = n.right.MinValue()

		n.right = n.right.Remove(n.value)
	}
	return n
}

func (n *Node[T]) MinValue() T {
	min := n.value
	for n.left != nil {
		min = n.left.value
		n = n.left
	}
	return min

}
