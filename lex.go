package main

import (
	"fmt"
	"unicode"
)

// type yyLexer interface {
//     Lex(lval *yySymType) int
//     Error(s string)
// }

type InputLex struct {
	s   string
	pos int
}

func (l *InputLex) Lex(lval *InputSymType) int {
	var c rune = ' '
	for c == ' ' {
		if l.pos == len(l.s) {
			return 0
		}
		c = rune(l.s[l.pos])
		l.pos = l.pos + 1
	}

	if unicode.IsDigit(c) || unicode.IsLower(c) {
		lval.val = string(c)
		fmt.Println("----val---", string(c))
		return CHARACTER
	}

	return int(c)

}

func (l *InputLex) Error(s string) {
	fmt.Printf("syntax error: %s\n", s)
}
