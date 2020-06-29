package pat

import (
	"unicode/utf8"
)

type Matcher func(ru rune) bool

type Expr []Matcher

func (expr Expr) Match(str string) bool {
	if len(expr) == 0 {
		return true
	}
	var i int
	var l = len(expr)
	for offset := 0; len(str) > 0 && i < l; str = str[offset:] {
		var ru, n = utf8.DecodeRuneInString(str)
		if n == 0 {
			return len(str) == 0 && i == l-1
		}
		var m = expr[i]
		switch {
		case m(ru):
			offset = n
		case i < l-1 && expr[i+1](ru):
			i++
		default:
			return len(str) == 0
		}
	}
	return i == l-1
}

func Or(left, right Matcher) Matcher {
	return func(ru rune) bool {
		return left(ru) || right(ru)
	}
}

func And(left, right Matcher) Matcher {
	return func(ru rune) bool {
		return left(ru) && right(ru)
	}
}
