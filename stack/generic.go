// Copyright (c) 2023 Michael D Henderson
// SPDX-License-Identifier: MIT
//

package stack

// Stack implements a generic stack structure.
type Stack[T any] struct {
	data []T
}

// New returns a new stack for type T.
func New[T any]() *Stack[T] {
	return &Stack[T]{nil}
}

// IsEmpty returns true if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// Peek returns the top element on the stack.
// If the stack is empty, returns a zero-value T and false.
func (s *Stack[T]) Peek() (tos T, ok bool) {
	if len(s.data) == 0 {
		return tos, false
	}
	return s.data[len(s.data)-1], true
}

// Pop returns the top element on the stack.
// If the stack is empty, returns a zero-value T and false.
func (s *Stack[T]) Pop() (tos T, ok bool) {
	if tos, ok = s.Peek(); ok {
		tos = s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1]
	}
	return tos, ok
}

// Push adds a new element to the stack.
func (s *Stack[T]) Push(n T) {
	s.data = append(s.data, n)
}

// Size returns the number of elements in the stack.
func (s *Stack[T]) Size() int {
	return len(s.data)
}
