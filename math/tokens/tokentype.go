package tokens

type TokenType byte

// Different tokens that are in the calculator
const (
	// values
	NUMBER TokenType = iota
	E
	PI

	// operations
	ADD
	SUB
	MUL
	DIV
	EXP

	// functions
	LOG // base 10
	LN
	SIN
	COS
	TAN
	ARCSIN
	ARCCOS
	ARCTAN
	SQRT

	// other
	LPAREN
	RPAREN
)
