package binary_tree

import "testing"

func TestInsert(t *testing.T) {
	tree := BinaryTree[int]{}
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(1)

	if tree.node.value != 2 {
		t.Errorf("Root value not 2. Got - %d", tree.node.value)
	}
	if tree.node.left.value != 1 {
		t.Errorf("Root left value not 1. Got - %d", tree.node.value)
	}
	if tree.node.right.value != 3 {
		t.Errorf("Root right value not 3. Got - %d", tree.node.value)
	}
	if tree.node.right.right.value != 4 {
		t.Errorf("Root right right value not 4. Got - %d", tree.node.value)
	}
}
func TestRemove(t *testing.T) {
	tree := BinaryTree[int]{}
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(1)

	tree.Remove(2)

	if tree.node.value != 3 {
		t.Errorf("Root value not 1. Got - %d", tree.node.value)
	}
	if tree.node.left.value != 1 {
		t.Errorf("Root left value not 3. Got - %d", tree.node.value)
	}
	if tree.node.right.value != 4 {
		t.Errorf("Root right value not 3. Got - %d", tree.node.value)
	}
	if tree.node.right.right != nil {
		t.Errorf("Root right right not nil. Got - %d", tree.node.value)
	}

	tree.Remove(1)

	if tree.node.value != 3 {
		t.Errorf("Root value not 1. Got - %d", tree.node.value)
	}
	if tree.node.left != nil {
		t.Errorf("Root left not nil. Got - %d", tree.node.value)
	}
	if tree.node.right.value != 4 {
		t.Errorf("Root right value not 3. Got - %d", tree.node.value)
	}

	tree.Remove(4)

	if tree.node.value != 3 {
		t.Errorf("Root value not 1. Got - %d", tree.node.value)
	}
	if tree.node.right != nil {
		t.Errorf("Root right not nil. Got - %d", tree.node.value)
	}

	tree.Remove(3)

	if tree.node != nil {
		t.Errorf("Root  not nil. Got - %d", tree.node.value)
	}
}
