package main

import (
	"fmt"

	"github.com/juan84n/estructurasdedatos/lists"
	"github.com/juan84n/estructurasdedatos/queue"
	"github.com/juan84n/estructurasdedatos/stack"
	"github.com/juan84n/estructurasdedatos/trees"
)

func main() {
	//callTree()
	//callArrayList()
	//callLinkedList()
	callQueue()
	//callStack()

}

func callArrayList() {
	arrayList := lists.NewArrayList()
	arrayList.Push(5)
	arrayList.Push(6)
	arrayList.Push(7)
	arrayList.PrintArrayList()
	index := arrayList.Search(7)
	fmt.Println("The position of 7 in array is", index)
	arrayList.Delete(7)
	arrayList.PrintArrayList()
}

func callLinkedList() {
	linkedList := lists.NewLinkedList()
	linkedList.Push(5)
	linkedList.Push(6)
	linkedList.Push(7)
	linkedList.PrintLinkedList()
	linkedSearched := linkedList.SearchGet(6)
	fmt.Println("the linkedSearched is ", linkedSearched)
	linkedList.Delete(6)
	linkedList.PrintLinkedList()
}

func callStack() {
	stack := stack.NewStack()
	stack.Push(5)
	stack.Push(6)
	stack.Push(7)
	stack.PrintStack()
	value := stack.Pop()
	fmt.Println("The Pop number in stack is ", value)
	stack.PrintStack()
}

func callQueue() {
	queue := queue.NewQueue()
	queue.Push(5)
	queue.Push(6)
	queue.Push(7)
	queue.PrintQueue()
	value := queue.Shift()
	fmt.Println("The shift number in queue is ", value)
	queue.PrintQueue()
}

func callTree() {
	tree := &trees.Node{}
	/*tree.Push(int32(10))
	tree.Push(int32(8))
	tree.Push(int32(9))
	tree.Push(int32(7))
	tree.Push(int32(6))
	tree.Push(int32(13))
	tree.Push(int32(11))
	tree.Push(int32(14))
	tree.Preorder()
	fmt.Println("")
	tree.Delete(10)
	tree.Preorder()
	fmt.Println("")
	tree.Inorder()
	fmt.Println("")
	tree.Postorder()
	fmt.Println("")*/
	var input int
	var nodeInput int
	for ok := true; ok; ok = (input != 6) {
		fmt.Println("1. Add node")
		fmt.Println("2. Loop preorden")
		fmt.Println("3. Loop inorden")
		fmt.Println("4. Find")
		fmt.Println("5. Delete")
		fmt.Println("6. Exit")

		fmt.Println("Make a choice ")
		n, err := fmt.Scanln(&input)
		if n < 1 || err != nil {
			fmt.Println("invalid input")
			break
		}

		switch input {
		case 1:
			fmt.Println("Input the node number ")
			n, err := fmt.Scanln(&nodeInput)
			if n < 1 || err != nil {
				fmt.Println("invalid input")
				break
			}
			tree.Push(int32(nodeInput))
			break
		case 2:
			tree.Preorder()
			fmt.Println("")
			break
		case 3:
			tree.Inorder()
			fmt.Println("")
			break
		case 4:
			fmt.Println("Input the node number to find ")
			n, err := fmt.Scanln(&nodeInput)
			if n < 1 || err != nil {
				fmt.Println("invalid input")
				break
			}
			node := tree.Exist(int32(nodeInput))
			if node.GetValue() > 0 {
				fmt.Println("exist", node)
			} else {
				fmt.Println("It doesnt exist")
			}
			break
		case 5:
			fmt.Println("Input the node number ")
			n, err := fmt.Scanln(&nodeInput)
			if n < 1 || err != nil {
				fmt.Println("invalid input")
				break
			}
			isDeleted := tree.Delete(int32(nodeInput))
			fmt.Println(" has been Deleted ? ", isDeleted)
			break
		default:
			fmt.Println("exit")
			break
		}
	}
}
