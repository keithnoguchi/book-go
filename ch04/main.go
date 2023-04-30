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
	s.remove(0)
	for {
		if data, err := s.pop(); err == nil {
			fmt.Println(data)
		} else {
			break
		}
	}
	fmt.Println(s)
}

type Stack []int

func (s *Stack) push(data ...int) {
	*s = append(*s, data...)
}

func (s *Stack) pop() (int, error) {
	n := len(*s)
	if n == 0 {
		return 0, errors.New("empty")
	}
	data := (*s)[n-1]
	*s = (*s)[:n-1]
	return data, nil
}

func (s *Stack) remove(pos uint) error {
	if int(pos) >= len(*s) {
		return errors.New("out of range")
	}
	copy((*s)[pos:], (*s)[pos+1:])
	*s = (*s)[:len(*s)-1]
	return nil
}
