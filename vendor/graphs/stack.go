package graphs

import "errors"

type stack struct {
	size  int
	index int
	mem   []string
}

func newStack() *stack {
	return &stack{
		size:  0,
		index: 0,
		mem:   make([]string, 0),
	}
}

func (s *stack) Push(v string) {
	s.mem = append(s.mem, v)
	s.size++
}

func (s *stack) Pop() (string, error) {
	if s.size == 0 {
		return "", errors.New("Stack: Pop operation on empty stack")
	}

	result := s.mem[s.size-1]
	s.mem = s.mem[:s.size-1]
	s.size--
	return result, nil
}

func (s stack) IsEmpty() bool {
	return s.size == 0
}
