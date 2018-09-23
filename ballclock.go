// This function implements a ballclock, brute force

package main

import (
	"errors"
	"fmt"
)

type stack struct {
	s        []int
	capacity int
}

func NewStack(max int) *stack {
	return &stack{make([]int, 0), max}
}

func (s *stack) Push(val int) error {
	if len(s.s) >= s.capacity {
		return errors.New("stack.Push() overflow")
	}
	s.s = append(s.s, val)
	return nil
}

func (s *stack) Pop() (int, error) {
	l := len(s.s)
	if l == 0 {
		return -1, errors.New("stack.Pop() underflow")
	}

	val := s.s[l-1]
	s.s = s.s[:l-1]
	return val, nil
}

func (s *stack) IsEmpty() bool {
	return len(s.s) == 0
}

func (s *stack) IsFull() bool {
	return len(s.s) >= s.capacity
}

func (s *stack) PushTilt(val int) (bool, int, []int) {
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

type queue struct {
	q        []int
	capacity int
}

func NewQueue(max int) *queue {
	return &queue{make([]int, 0), max}
}

func (q *queue) Push(vals ...int) error {
	if len(q.q)+len(vals) > q.capacity {
		panic("cant push to queue.Push(), will be overflowed")
	}
	for _, val := range vals {
		q.q = append(q.q, val)
	}
	return nil
}

func (q *queue) Pop() (int, error) {
	l := len(q.q)
	if l == 0 {
		return -1, errors.New("queue.Pop() underflow")
	}

	val := q.q[0]
	q.q = q.q[1:l]
	return val, nil
}

func (q *queue) Fill() {
	for i := len(q.q); i < q.capacity; i++ {
		q.Push(i)
	}
}

func (q *queue) IsCycle() bool {
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
	numBalls := 30

	stack1M := NewStack(5)
	stack5M := NewStack(12)
	stack60M := NewStack(12)
	q := NewQueue(numBalls)

	q.Fill()

	for n := 0; ; n++ {
		ball, _ := q.Pop()
		tilt, ball, slice := stack1M.PushTilt(ball)
		if tilt {
			q.Push(slice...)
			tilt, ball, slice := stack5M.PushTilt(ball)
			if tilt {
				q.Push(slice...)
				tilt, ball, slice := stack60M.PushTilt(ball)
				if tilt {
					q.Push(slice...)
					q.Push(ball)
				}
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

real  0m0.692s
user  0m0.568s
sys   0m0.300s
*/
