// This function implements a ballclock, brute force

package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	s        []int
	capacity int
}

func NewStack(max int) *Stack {
	return &Stack{make([]int, 0), max}
}

func (s *Stack) Push(vals ...int) error {
	if len(s.s)+len(vals) > s.capacity {
		return errors.New("stack.Push() overflow")
	}
	s.s = append(s.s, vals...)
	return nil
}

func (s *Stack) Pop() (int, error) {
	l := len(s.s)
	if l == 0 {
		return -1, errors.New("stack.Pop() underflow")
	}

	val := s.s[l-1]
	s.s = s.s[:l-1]
	return val, nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.s) == 0
}

func (s *Stack) IsFull() bool {
	return len(s.s) >= s.capacity
}

func (s *Stack) Drop(val int) (bool, int, []int) {
	err := s.Push(val)
	if err != nil {
		panic(err)
	}
	if s.IsFull() == false {
		return false, -1, nil
	}

	// stack is full, now tilt
	ball, _ := s.Pop()
	slice := make([]int, 0)
	for !s.IsEmpty() {
		val, _ := s.Pop()
		slice = append(slice, val)
	}

	return true, ball, slice
}

type Queue struct {
	q        []int
	capacity int
}

func NewQueue(max int) *Queue {
	return &Queue{make([]int, 0), max}
}

func (q *Queue) Push(vals ...int) error {
	if len(q.q)+len(vals) > q.capacity {
		panic("cant push to queue.Push(), will be overflowed")
	}
	q.q = append(q.q, vals...)
	return nil
}

func (q *Queue) Pop() (int, error) {
	l := len(q.q)
	if l == 0 {
		return -1, errors.New("queue.Pop() underflow")
	}

	val := q.q[0]
	q.q = q.q[1:l]
	return val, nil
}

func (q *Queue) Fill() *Queue {
	for i := len(q.q); i < q.capacity; i++ {
		q.Push(i)
	}
	return q
}

func (q *Queue) IsCycle() bool {
	if len(q.q) != q.capacity {
		return false
	}

	for i := 0; i < q.capacity; i++ {
		if q.q[i] != i {
			return false
		}
	}
	return true
}

func main() {
	stack := []*Stack{NewStack(5), NewStack(12), NewStack(12)}
	numStacks := len(stack)

	numBalls := 30
	q := NewQueue(numBalls).Fill()

	for n := 0; ; n++ {
		ball, _ := q.Pop()
		for i := 0; i < numStacks; i++ {
			tilt, balltop, slice := stack[i].Drop(ball)
			if !tilt {
				break
			}
			q.Push(slice...)
			if i == numStacks-1 { // a pop from last stack will go into q
				q.Push(balltop)
			} else {
				ball = balltop
			}
		}
		if (n+1)%(12*60) == 0 && q.IsCycle() {
			fmt.Println(numBalls, "balls cycle after", (n+1)/(24*60), "days.")
			break
		}
	}
}

/*
$ time go run ballclock.go
30 balls cycle after 15 days.

real  0m0.600s
user  0m0.524s
sys   0m0.200s
*/
