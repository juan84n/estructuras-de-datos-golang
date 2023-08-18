package stack

import "fmt"

type Stack struct {
	value int32
	array []int32
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(value int32) {
	s.array = append(s.array, value)
}

func (s *Stack) Search(value int32) int {
	for i := 0; i < len(s.array); i++ {
		if s.array[i] == value {
			return i
		}
	}
	return -1
}

func (s *Stack) Pop() int32 {
	pop := s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]
	return pop
}

func (s *Stack) Peek() int32 {
	peek := s.array[len(s.array)-1]
	return peek
}

func (s *Stack) Delete(value int32) {
	index := s.Search(value)
	s.array = append(s.array[:index], s.array[index+1:]...)
}

func (s *Stack) PrintStack() {
	fmt.Print("LinkedList is [", s.array[len(s.array)-1])
	for i := len(s.array) - 2; i >= 0; i-- {
		fmt.Print(" ", s.array[i])
	}
	fmt.Println("]")
}
