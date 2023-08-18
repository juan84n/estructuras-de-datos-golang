package lists

import "fmt"

// import "fmt"

type LinkedList struct {
	value int32
	next  *LinkedList
	prev  *LinkedList
}

var pointer = &LinkedList{}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (a *LinkedList) Push(value int32) {
	if a.value == 0 {
		a.value = value
		a.next = nil
		pointer = a
		pointer.prev = nil
	} else {
		pointer.next = &LinkedList{value: value, prev: pointer}
		pointer = pointer.next

	}
}

func (a *LinkedList) SearchGet(value int32) *LinkedList {
	auxiliar := a
	for auxiliar.value > 0 {
		if auxiliar.value == value {
			return auxiliar
		}
		auxiliar = auxiliar.next
	}

	return &LinkedList{}
}

func (a *LinkedList) Delete(value int32) {
	if a.value == value {
		a.value = a.next.value
		a.next = a.next.next
	} else if pointer.value == value {
		pointer.value = 0
		pointer.prev.next = nil
		pointer = pointer.prev
	} else {
		searched := a.SearchGet(value)
		searched.prev.next = searched.next
		searched.next.prev = searched.prev
		searched.value = 0

	}

}

func (a *LinkedList) PrintLinkedList() {
	fmt.Print("LinkedList is [", a.value)
	auxiliar := a.next
	for auxiliar != nil && auxiliar.value > 0 {
		fmt.Print(" ", auxiliar.value)
		auxiliar = auxiliar.next
	}
	fmt.Println("]")
}
