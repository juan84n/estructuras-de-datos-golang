package queue

import "fmt"

type Queue struct {
	value int32
	array []int32
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(value int32) {
	q.array = append(q.array, value)
}

func (q *Queue) Search(value int32) int {
	for i := 0; i < len(q.array); i++ {
		if q.array[i] == value {
			return i
		}
	}
	return -1
}

func (q *Queue) Shift() int32 {
	shift := q.array[0]
	q.array = q.array[1:]
	return shift
}

func (q *Queue) Peek() int32 {
	peek := q.array[0]
	return peek
}

func (q *Queue) Delete(value int32) {
	index := q.Search(value)
	q.array = append(q.array[:index], q.array[index+1:]...)
}

func (q *Queue) PrintQueue() {
	fmt.Println("Queue is ", q.array)
}
