package main

type Stack struct {
	elements []byte
}

func (s *Stack) Push(element byte) {
	s.elements = append(s.elements, element)
}
func (s *Stack) Pop() (byte, bool) {
	if len(s.elements) == 0 {
		return 0, false
	}

	element := s.elements[len(s.elements)-1]

	s.elements = s.elements[:len(s.elements)-1]
	return element, true
}

func (s *Stack) Peek() (byte, bool) {
	if len(s.elements) == 0 {
		return 0, false
	}
	return s.elements[len(s.elements)-1], true
}
func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}
