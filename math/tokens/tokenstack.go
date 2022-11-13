package tokens

type TokenStack struct {
	tokens []Token
}

// Empty returns if the stack size is 0
func (t TokenStack) Empty() bool {
	return len(t.tokens) == 0
}

// Push adds a Token to the top of the stack
func (t *TokenStack) Push(token Token) {
	t.tokens = append(t.tokens, token)
}

// Pop removes the item at the top of the stack and returns it
func (t *TokenStack) Pop() *Token {
	length := len(t.tokens)

	if length == 0 {
		return nil
	}

	last := t.tokens[length-1]

	t.tokens = t.tokens[:length-1]

	return &last
}

// Peek returns the item at the top of the stack
func (t *TokenStack) Peek() *Token {
	length := len(t.tokens)

	if length == 0 {
		return nil
	}

	last := t.tokens[length-1]

	return &last
}

func TS() TokenStack {
	return TokenStack{}
}
