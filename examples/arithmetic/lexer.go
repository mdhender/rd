// Copyright (c) 2023 Michael D Henderson
// Copyright (c) 2018 Shivam Mamgain
// SPDX-License-Identifier: MIT
//

package main

import (
	"fmt"

	"github.com/alecthomas/chroma"
	"github.com/mdhender/rd"
)

var lexer = chroma.MustNewLexer(
	&chroma.Config{
		Name: "Arithmetic Expressions",
	},
	chroma.Rules{
		"root": {
			{`\s+`, chroma.Text, nil},
			{`[+\-*/]`, chroma.Operator, nil},
			{`[()]`, chroma.Punctuation, nil},
			{`\d*\.\d+`, chroma.NumberFloat, nil},
			{`\d+`, chroma.NumberInteger, nil},
		},
	},
)

func Lex(expr string) (tokens []rd.Token, err error) {
	iter, err := lexer.Tokenise(nil, expr)
	if err != nil {
		return nil, err
	}
	for _, token := range iter.Tokens() {
		switch token.Type {
		case chroma.Error:
			return nil, fmt.Errorf("invalid token: %v", token)
		case chroma.Operator, chroma.Punctuation, chroma.NumberFloat, chroma.NumberInteger:
			tokens = append(tokens, token.Value)
		}
	}
	return tokens, nil
}
