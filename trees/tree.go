package trees

import (
	"fmt"
)

type Node struct {
	value     int32
	leftNode  *Node
	rightNode *Node
}

func (t *Node) GetValue() int32 {
	return t.value
}
func (t *Node) GetLeftNode() *Node {
	return t.leftNode
}
func (t *Node) GetRightNode() *Node {
	return t.rightNode
}

func (t *Node) Push(value int32) {
	if t.value == 0 {
		t.value = value
		t.leftNode = &Node{}
		t.rightNode = &Node{}
	} else {
		if value < t.value {
			t.leftNode.Push(value)
		} else {
			t.rightNode.Push(value)
		}
	}
}

func (t *Node) Preorder() {
	if t.value > 0 {
		fmt.Print(t.value, " ")
		t.leftNode.Preorder()
		t.rightNode.Preorder()
	}

}

func (t *Node) Inorder() {
	if t.value > 0 {
		t.leftNode.Inorder()
		fmt.Print(t.value, " ")
		t.rightNode.Inorder()
	}
}

func (t *Node) Postorder() {
	if t.value > 0 {
		t.leftNode.Postorder()
		t.rightNode.Postorder()
		fmt.Print(t.value, " ")
	}
}

func (t *Node) Exist(value int32) *Node {
	if t.value == 0 {
		return &Node{}
	}
	if t.value == value {
		return t
	} else if value < t.value {
		return t.leftNode.Exist(value)
	} else {
		return t.rightNode.Exist(value)
	}

}

func (t *Node) Delete(value int32) bool {
	if t.value == 0 {
		return false
	}
	exist := t.Exist(value)
	if exist.leftNode.value == 0 && exist.rightNode.value == 0 {
		exist.value = 0
		exist.leftNode = nil
		exist.rightNode = nil
		return true
	}
	if exist.leftNode.value > 0 && exist.rightNode.value == 0 {
		replacement := replaceLeftChildren(exist)
		exist.value = replacement.value
		exist.leftNode = replacement.leftNode
		return true
	}
	if exist.leftNode.value == 0 && exist.rightNode.value > 0 {
		replacement := replaceRightChildren(exist)
		exist.value = replacement.value
		exist.rightNode = replacement.rightNode
	}
	if exist.leftNode.value > 0 && exist.rightNode.value > 0 {
		replacement := replaceLeftChildren(exist.rightNode)
		existReplacement := t.Exist(replacement.value)
		existReplacement.value = 0
		existReplacement.leftNode = nil
		existReplacement.rightNode = nil
		exist.value = replacement.value
		return true
	}
	return false
}

func replaceLeftChildren(tree *Node) Node {
	auxiliar := Node{value: tree.value, leftNode: tree.leftNode}
	if auxiliar.leftNode.value == 0 {
		return auxiliar
	}

	flag := false
	for flag == false {
		if auxiliar.leftNode.value > 0 {
			auxiliar = Node{value: auxiliar.leftNode.value, leftNode: auxiliar.leftNode.leftNode}
		} else {
			flag = true
		}
	}
	return auxiliar
}

func replaceRightChildren(tree *Node) Node {
	auxiliar := Node{value: tree.value, rightNode: tree.rightNode}
	if auxiliar.rightNode.value == 0 {
		return auxiliar
	}
	flag := false
	for flag == false {
		if auxiliar.rightNode.value > 0 {
			auxiliar = Node{value: auxiliar.rightNode.value, rightNode: auxiliar.rightNode.rightNode}
		} else {
			flag = true
		}
	}
	return auxiliar
}
