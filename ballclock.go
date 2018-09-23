// This function implments a ballclock, brute force

package main

import "fmt"

type stack struct {
	s           []int
	maxCapacity int
}

func NewStack(max int) *stack {
	return &stack{make([]int, 0), max}
}

func (s *stack) PushNTilt(val int) (bool, int, []int) {
	s.Push(val)
	if !s.IsFull() {
		return false, 0, nil
	}

	val1 := s.Pop()
	slice := make([]int, 0)
	for !s.IsEmpty() {
		val = s.Pop()
		slice = append(slice, val)
	}

	return true, val1, slice
}

func (s *stack) IsEmpty() bool {
	return len(s.s) == 0
}

func (s *stack) IsFull() bool {
	return len(s.s) == s.maxCapacity
}

func (s *stack) Push(val int) {
	s.s = append(s.s, val)
}

func (s *stack) Pop() int {
	l := len(s.s)
	if l == 0 {
		panic("Exception occurred trying to pop an emtpy stack")
	}

	val := s.s[l-1]
	s.s = s.s[:l-1]
	return val
}

type queue struct {
	stack
}

func NewQueue(max int) *queue {
	s := NewStack(max)
	for ball := 0; ball < s.maxCapacity; ball++ {
		s.Push(ball)
	}
	return &queue{*s}
}

func (q *queue) Push(v int) {
	q.stack.Push(v)
}

func (q *queue) Pop() int {
	l := len(q.stack.s)
	if l == 0 {
		panic("Exception occurred trying to pop an empty queue")
	}

	res := q.stack.s[0]
	q.stack.s = q.stack.s[1:l]
	return res
}

func (q *queue) PushMany(slice []int) {
	for _, v := range slice {
		q.Push(v)
	}
}

func (q *queue) IsTheSame() bool {
	if len(q.stack.s) != q.stack.maxCapacity {
		return false
	}

	for ball := 0; ball < q.stack.maxCapacity; ball++ {
		if q.stack.s[ball] != ball {
			return false
		}
	}
	return true
}

func main() {

	stack1M := NewStack(5)
	stack5M := NewStack(12)
	stack60M := NewStack(12)

	numBalls := 30
	qContainer := NewQueue(numBalls)

	for n := 0; ; n++ {
		ball := qContainer.Pop()
		tilt, ballOverflow, slice := stack1M.PushNTilt(ball)
		if tilt {
			qContainer.PushMany(slice)
			tilt, ballOverflow, slice = stack5M.PushNTilt(ballOverflow)
			if tilt {
				qContainer.PushMany(slice)
				tilt, ballOverflow, slice := stack60M.PushNTilt(ballOverflow)
				if tilt {
					qContainer.PushMany(slice)
					qContainer.Push(ballOverflow)
				}
			}
		}
		if (n+1)%(12*60) == 0 && qContainer.IsTheSame() {
			fmt.Println(numBalls, "balls cycle after", (n+1)/(24*60), "days.")
			break
		}
	}
}

/*
$go run ballclock.go
30 balls cycle after 15 days.
*/
