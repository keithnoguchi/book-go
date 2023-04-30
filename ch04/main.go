package main

import (
	"errors"
	"fmt"
)

func main() {
	s := Stack{}
	s.push(1)
	s.push(2)
	s.push(3)
	s.push(4, 5, 6)
	fmt.Println(s)
	for {
		if data, err := s.pop(); err == nil {
			fmt.Println(data)
		} else {
			break
		}
	}
	fmt.Println(s)
}

type Stack struct {
	stack []int
}

func (s *Stack) pop() (int, error) {
	n := len(s.stack)
	if n == 0 {
		return 0, errors.New("empty")
	}
	data := s.stack[n-1]
	s.stack = s.stack[:n-1]
	return data, nil
}

func (s *Stack) push(data ...int) {
	s.stack = append(s.stack, data...)
}

func (s Stack) String() string {
	return fmt.Sprintf("%v", s.stack)
}
