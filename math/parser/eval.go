package parser

import (
	"fmt"
	"math"

	"github.com/narutopig/calculator/math/tokens"
)

// Eval returns the value of the expression given in RPN
func Eval(ts tokens.TokenStack) (float64, error) {
	toks := ts.Reverse()
	stack := fs()

	for !toks.Empty() {
		token := toks.Pop()

		if token.Type == tokens.NUMBER {
			stack.push(token.Value)
		} else if token.Type == tokens.E {
			stack.push(math.E)
		} else if token.Type == tokens.PI {
			stack.push(math.Pi)
		} else {
			// operator

			right, err := stack.pop()
			if err != nil {
				return 0, err
			}

			left, err := stack.pop()
			if err != nil {
				return 0, err
			}

			if token.Type == tokens.ADD {
				stack.push(left + right)
			} else if token.Type == tokens.SUB {
				stack.push(left - right)
			} else if token.Type == tokens.MUL {
				stack.push(left * right)
			} else if token.Type == tokens.DIV {
				if right == 0 {
					return 0, fmt.Errorf("Error: Division by zero")
				}
				stack.push(left / right)
			} else if token.Type == tokens.EXP {
				stack.push(math.Pow(right, left))
			}
		}
	}

	return stack.pop()
}
