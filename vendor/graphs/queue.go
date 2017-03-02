package graphs

import "errors"

type queue struct {
	size int
	mem  []string
}

func newQueue() *queue {
	return &queue{
		size: 0,
		mem:  make([]string, 0),
	}
}

func (s *queue) Push(v string) {
	s.mem = append(s.mem, v)
	s.size++
}

func (s *queue) Pop() (string, error) {
	if s.size == 0 {
		return "", errors.New("queue: Pop operation on empty queue")
	}

	result := s.mem[0]
	s.mem = s.mem[1:]
	s.size--
	return result, nil
}

func (s queue) IsEmpty() bool {
	return s.size == 0
}
