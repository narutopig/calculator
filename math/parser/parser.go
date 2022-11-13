package parser

import "github.com/narutopig/calculator/math/tokens"

func Shunt(input []tokens.Token) tokens.TokenStack {
	output := tokens.TS()
	operator := tokens.TS()

	length := len(input)

	for i := 0; i < length; i++ {
		token := input[i]

		if token.Type == tokens.NUMBER {
			output.Push(token)
		} else if token.Type == tokens.E { // TODO: change to be for all constants
			if i+1 < length && input[i+1].Type == tokens.LPAREN {
				// next token is left paren, probably function
				operator.Push(token)
			} else {
				// regular identifier
				output.Push(token)
			}
		} else if token.Category() == tokens.OPERATOR {
			o1 := token
			o2 := operator.Peek()
			// assumes all operators are left-associative
			for o2.Type != tokens.LPAREN && o2.Precedence() >= o1.Precedence() {
				top := operator.Pop()
				if top == nil {
					panic("missing expression on right side")
				}

				output.Push(*top)

				o2 = operator.Peek()
			}
			operator.Push(o1)
		} else if token.Type == tokens.LPAREN {
			operator.Push(token)
		} else if token.Type == tokens.RPAREN {
			for operator.Peek().Type != tokens.LPAREN {
				if operator.Empty() {
					panic("mismatched parentheses")
				}
				output.Push(*operator.Pop())
			}
			if operator.Peek().Type != tokens.LPAREN {
				panic("mismatched parentheses")
			}
			operator.Pop()
			if operator.Peek().Category() == tokens.FUNCTION {
				output.Push(*operator.Pop())
			}
		}
	}

	for !operator.Empty() {
		front := operator.Peek()
		if front.Type == tokens.LPAREN || front.Type == tokens.RPAREN {
			panic("mismatched parentheses")
		}

		output.Push(*operator.Pop())
	}

	return output
}
