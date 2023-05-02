// Copyright (c) 2023 Michael D Henderson
// Copyright (c) 2018 Shivam Mamgain
// SPDX-License-Identifier: MIT
//

package parser

import (
	"fmt"
	"regexp"

	"github.com/mdhender/rd"
	. "github.com/mdhender/rd/examples/arithmetic/tokens"
)

const Grammar = `
	Expr   = Term Expr'
	Expr'  = "+" Expr | "-" Expr | ε
	Term   = Factor Term'
	Term'  = "*" Term | "/" Term | ε
	Factor = "(" Expr ")" | "-" Factor | Number
`

var numberRegex = regexp.MustCompile(`^(\d*\.\d+|\d+)$`)

func Expr(b *rd.Builder) (ok bool) {
	defer b.Enter("Expr").Exit(&ok)

	return Term(b) && ExprPrime(b)
}

func ExprPrime(b *rd.Builder) (ok bool) {
	defer b.Enter("Expr'").Exit(&ok)

	if b.Match(Plus) {
		return Expr(b)
	}
	if b.Match(Minus) {
		return Expr(b)
	}
	if b.CheckOrNotOK(CloseParen, 1) {
		b.Skip()
		return true
	}
	return false
}

func Term(b *rd.Builder) (ok bool) {
	defer b.Enter("Term").Exit(&ok)

	return Factor(b) && TermPrime(b)
}

func TermPrime(b *rd.Builder) (ok bool) {
	defer b.Enter("Term'").Exit(&ok)

	if b.Match(Star) {
		return Term(b)
	}
	if b.Match(Slash) {
		return Term(b)
	}
	b.Skip()
	return true
}

func Factor(b *rd.Builder) (ok bool) {
	defer b.Enter("Factor").Exit(&ok)

	if b.Match(OpenParen) {
		return Expr(b) && b.Match(CloseParen)
	}
	if b.Match(Minus) {
		return Factor(b)
	}
	return Number(b)
}

func Number(b *rd.Builder) (ok bool) {
	defer b.Enter("Number").Exit(&ok)

	token, ok := b.Next()
	if !ok {
		return false
	}
	if numberRegex.MatchString(fmt.Sprint(token)) {
		b.Add(token)
		return true
	}
	return false
}

func Parse(tokens []rd.Token) (parseTree *rd.Tree, debugTree *rd.DebugTree, err error) {
	b := rd.NewBuilder(tokens)
	if ok := Expr(b); ok && b.Err() == nil {
		return b.ParseTree(), b.DebugTree(), nil
	}
	return nil, b.DebugTree(), b.Err()
}
