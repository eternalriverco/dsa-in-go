package sliceimp

import (
	"fmt"
	"testing"
)

type stack []interface{}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) len() int {
	return len(*s)
}

func (s *stack) pop() interface{} {
	if s.isEmpty() {
		return nil
	}
	out := (*s)[s.len()-1]
	*s = (*s)[:s.len()-1]
	return out
}

func (s *stack) push(v interface{}) {
	*s = append(*s, v)
}

func (s *stack) printAll() {
	for _, v := range *s {
		fmt.Println(v)
	}
}

func TestSliceImp(t *testing.T) {
	var s stack
	s.push(1)
	s.push(2)
	s.push(3)
	s.push(4)
	s.push(5)
	s.pop()
	s.pop()
	s.push(6)
	s.printAll()
}
