// Copyright (c) 2023 Michael D Henderson
// Copyright (c) 2018 Shivam Mamgain
// SPDX-License-Identifier: MIT
//

package main

import (
	"testing"

	"github.com/mdhender/rd"
	"github.com/mdhender/rd/examples/arithmetic/parser"
)

func TestArithmeticExpressionsGrammar(t *testing.T) {
	tokens := []rd.Token{"2.8", "+", "(", "3", "-", ".733", ")", "/", "23"}
	parseTree, debugTree, err := parser.Parse(tokens)
	if err != nil {
		t.Error("parsing failed")
	}

	expectedDebugTree := `Expr(true)
├─ Term(true)
│  ├─ Factor(true)
│  │  ├─ 2.8 ≠ (
│  │  ├─ 2.8 ≠ -
│  │  └─ Number(true)
│  └─ Term'(true)
│     ├─ + ≠ *
│     └─ + ≠ /
└─ Expr'(true)
   ├─ + = +
   └─ Expr(true)
      ├─ Term(true)
      │  ├─ Factor(true)
      │  │  ├─ ( = (
      │  │  ├─ Expr(true)
      │  │  │  ├─ Term(true)
      │  │  │  │  ├─ Factor(true)
      │  │  │  │  │  ├─ 3 ≠ (
      │  │  │  │  │  ├─ 3 ≠ -
      │  │  │  │  │  └─ Number(true)
      │  │  │  │  └─ Term'(true)
      │  │  │  │     ├─ - ≠ *
      │  │  │  │     └─ - ≠ /
      │  │  │  └─ Expr'(true)
      │  │  │     ├─ - ≠ +
      │  │  │     ├─ - = -
      │  │  │     └─ Expr(true)
      │  │  │        ├─ Term(true)
      │  │  │        │  ├─ Factor(true)
      │  │  │        │  │  ├─ .733 ≠ (
      │  │  │        │  │  ├─ .733 ≠ -
      │  │  │        │  │  └─ Number(true)
      │  │  │        │  └─ Term'(true)
      │  │  │        │     ├─ ) ≠ *
      │  │  │        │     └─ ) ≠ /
      │  │  │        └─ Expr'(true)
      │  │  │           ├─ ) ≠ +
      │  │  │           └─ ) ≠ -
      │  │  └─ ) = )
      │  └─ Term'(true)
      │     ├─ / ≠ *
      │     ├─ / = /
      │     └─ Term(true)
      │        ├─ Factor(true)
      │        │  ├─ 23 ≠ (
      │        │  ├─ 23 ≠ -
      │        │  └─ Number(true)
      │        └─ Term'(true)
      │           ├─ <no tokens left> ≠ *
      │           └─ <no tokens left> ≠ /
      └─ Expr'(true)
         ├─ <no tokens left> ≠ +
         └─ <no tokens left> ≠ -
`
	got := debugTree.String()
	if got != expectedDebugTree {
		t.Errorf("invalid debug tree. expected: %s\ngot: %s\n", expectedDebugTree, got)
	}

	expectedParseTree := `Expr
├─ Term
│  └─ Factor
│     └─ Number
│        └─ 2.8
└─ Expr'
   ├─ +
   └─ Expr
      └─ Term
         ├─ Factor
         │  ├─ (
         │  ├─ Expr
         │  │  ├─ Term
         │  │  │  └─ Factor
         │  │  │     └─ Number
         │  │  │        └─ 3
         │  │  └─ Expr'
         │  │     ├─ -
         │  │     └─ Expr
         │  │        └─ Term
         │  │           └─ Factor
         │  │              └─ Number
         │  │                 └─ .733
         │  └─ )
         └─ Term'
            ├─ /
            └─ Term
               └─ Factor
                  └─ Number
                     └─ 23
`
	got = parseTree.String()
	if got != expectedParseTree {
		t.Errorf("invalid parse tree. want: %s\ngot: %s\n", expectedParseTree, got)
	}
}
