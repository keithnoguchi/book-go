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

type Stack struct {
	slice []int
}

func (s Stack) String() string {
	return fmt.Sprintf("%v", s.slice)
}

func (s *Stack) push(data ...int) {
	s.slice = append(s.slice, data...)
}

func (s *Stack) pop() (int, error) {
	n := len(s.slice)
	if n == 0 {
		return 0, errors.New("empty")
	}
	data := s.slice[n-1]
	s.slice = s.slice[:n-1]
	return data, nil
}

func (s *Stack) remove(pos uint) error {
	if int(pos) >= len(s.slice) {
		return errors.New("out of range")
	}
	copy(s.slice[pos:], s.slice[pos+1:])
	s.slice = s.slice[:len(s.slice)-1]
	return nil
}
