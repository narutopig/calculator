package tokens

type Token struct {
	Type  TokenType
	Value float64
}

func T(t TokenType, v float64) Token {
	return Token{Type: t, Value: v}
}

func (t Token) Category() TokenCategory {
	tt := t.Type

	if tt >= LPAREN {
		return OTHER
	} else if tt >= LOG {
		return FUNCTION
	} else if tt >= ADD {
		return OPERATOR
	}
	return VALUE
}

func (t Token) Precedence() int {
	tt := t.Type

	if tt == ADD || tt == SUB {
		return 0
	} else if tt == MUL || tt == DIV {
		return 1
	} else if tt == EXP {
		return 2
	}
	return -1
}
