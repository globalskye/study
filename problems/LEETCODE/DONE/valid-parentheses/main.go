package main

import "fmt"

func main() {
	fmt.Println(isValid("{{}}[]"))
}

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
func isValid(s string) bool {
	if len(s) == 0 || len(s) == 1 || len(s)%2 == 1 {
		return false
	}

	var m Stack
	for i := 0; i < len(s); i++ {

		if s[i] == '[' || s[i] == '(' || s[i] == '{' {
			m.Push(s[i])
			continue
		}
		str, ok := m.Pop()
		if !ok {
			return false
		}
		if str == '[' && s[i] != ']' {
			return false
		}
		if str == '(' && s[i] != ')' {
			return false
		}
		if str == '{' && s[i] != '}' {
			return false
		}

	}
	if m.IsEmpty() {
		return true
	}
	return false
}
