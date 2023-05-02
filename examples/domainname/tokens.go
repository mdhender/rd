// Copyright (c) 2023 Michael D Henderson
// Copyright (c) 2018 Shivam Mamgain
// SPDX-License-Identifier: MIT
//

package main

type Token int

const (
	Illegal Token = iota
	EOF
	Whitespace
	Dot
	Hyphen
	Letter
	Digit
)

func (t Token) String() string {
	switch t {
	case 2:
		return " "
	case 3:
		return "."
	case 4:
		return "-"
	}
	return ""
}
