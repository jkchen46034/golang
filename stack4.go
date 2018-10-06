// Implements a thread safe stack

package main

import (
	"fmt"
	"sync"
)

type Stack struct {
	mutex sync.Mutex
	stack []int
}

func NewStack() *Stack {
	return &Stack{stack: make([]int, 0)}
}

func (s *Stack) Push(val ...int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.stack = append(s.stack, val...)
}

func (s *Stack) Pop() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	length := len(s.stack)
	val := s.stack[length-1]
	s.stack = s.stack[0 : length-1]
	return val
}

func (s *Stack) thread(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 200000; i++ {
		s.Push(1, 2, 3, 4, 5)
		s.Pop()
		s.Pop()
		s.Pop()
		s.Pop()
		s.Pop()
	}
	fmt.Println("done")
}

func main() {
	var wg sync.WaitGroup
	s := NewStack()
	s.Push(8, 9, 10, 11, 12, 13)
	fmt.Println(s.stack)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go s.thread(&wg)
	}
	wg.Wait()
	fmt.Println(s.stack)
}

/*
$ go run stack4.go
[8 9 10 11 12 13]
done
done
done
done
done
done
done
done
done
done
[8 9 10 11 12 13]
*/
