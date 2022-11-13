package tokens

type TokenCategory byte

const (
	VALUE TokenCategory = iota
	OPERATOR
	FUNCTION
	OTHER
)
