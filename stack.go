// Copyright (c) 2023 Michael D Henderson
// Copyright (c) 2018 Shivam Mamgain
// SPDX-License-Identifier: MIT
//

package rd

type stack struct {
	stack []ele
}

func (s *stack) isEmpty() bool {
	return len(s.stack) == 0
}

// peek will panic if stack is empty
func (s *stack) peek() ele {
	return s.stack[len(s.stack)-1]
}

// pop will panic if stack is empty
func (s *stack) pop() ele {
	top := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return top
}

func (s *stack) push(e ele) {
	s.stack = append(s.stack, e)
}
