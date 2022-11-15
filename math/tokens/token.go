package tokens

import "fmt"

// Token represents a token in a exprsesion
type Token struct {
	Type  TokenType
	Value float64
}

// T returns a new Token with the given TokenType and float64 value
func T(t TokenType, v float64) Token {
	return Token{Type: t, Value: v}
}

// Category returns which category the token falls under
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

// Precedence returns what an operator's precedence is
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

func (t Token) Stringify() string {
	switch t.Type {
	case NUMBER:
		return fmt.Sprintf("%f", t.Value)
	case E:
		return "e"
	case PI:
		return "π"
	case ADD:
		return "+"
	case SUB:
		return "-"
	case MUL:
		return "·"
	case DIV:
		return "÷"
	case EXP:
		return "^"
	case LOG:
		return "log"
	case LN:
		return "ln"
	case SIN:
		return "sin"
	case COS:
		return "cos"
	case TAN:
		return "tan"
	case ARCSIN:
		return "asin"
	case ARCCOS:
		return "acos"
	case ARCTAN:
		return "atan"
	case SQRT:
		return "√"
	case LPAREN:
		return "("
	case RPAREN:
		return ")"
	default:
		return ""
	}
}

// SPrintArr prints a list of tokens in a readable format
func SPrintArr(tokens []Token) string {
	res := ""

	for _, t := range tokens {
		res += t.Stringify()
	}

	return res
}

func (t Token) String() string {
	return fmt.Sprintf("Token{Type:%s,Value:%f}", t.Type, t.Value)
}
