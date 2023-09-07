package lists

import "fmt"

type ArrayList struct {
	value int32
	array []int32
}

func NewArrayList() *ArrayList {
	return &ArrayList{}
}

func (a *ArrayList) Push(value int32) {
	a.array = append(a.array, value)
}

func (a *ArrayList) Search(value int32) int {
	for i := 0; i < len(a.array); i++ {
		if a.array[i] == value {
			return i
		}
	}
	return -1
}

func (a *ArrayList) Get(index int) int32 {
	return a.array[index]
}

func (a *ArrayList) Delete(value int32) {
	index := a.Search(value)
	a.array = append(a.array[:index], a.array[index+1:]...)
}

func (a *ArrayList) Size() int {
	return len(a.array)
}

func (a *ArrayList) PrintArrayList() {
	fmt.Println("ArrayList is ", a.array)
}

func equals[C comparable](a []C, value C) int {
	for i := 0; i < len(a); i++ {
		if a[i] == value {
			return i
		}
	}
	return -1
}
