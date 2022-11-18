package main

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/narutopig/calculator/math/parser"
	"github.com/narutopig/calculator/math/tokens"
)

func main() {
	app := app.New()
	window := app.NewWindow("Calculator")

	currTokens := make([]tokens.Token, 0)
	currNum := ""

	display := container.New(layout.NewVBoxLayout())

	// needed to update stuff
	current := canvas.NewText("", color.White)
	result := canvas.NewText("", color.White)

	add := func(token tokens.Token) {
		currTokens = append(currTokens, token)
	}

	refresh := func() {
		current.Refresh()
	}

	update := func(token tokens.Token) {
		if currNum != "" {
			if currNum[len(currNum)-1] == []byte(".")[0] {
				currNum += "0"
			}
			val, err := strconv.ParseFloat(currNum, 64)
			if err != nil {
				panic(err)
			}
			add(tokens.T(tokens.NUMBER, val))
			currNum = ""
		}
		add(token)
		current.Text += token.Stringify()
		refresh()
	}

	{
		display.Add(current)
		display.Add(result)
	}

	basic := container.New(layout.NewGridLayout(4))
	{
		clearBtn := widget.NewButton("AC", func() {
			currTokens = make([]tokens.Token, 0)
			currNum = ""
			current.Text = ""
			refresh()
		})
		lpBtn := widget.NewButton("(", func() {
			update(tokens.T(tokens.LPAREN, 0))
		})
		rpBtn := widget.NewButton(")", func() {
			update(tokens.T(tokens.RPAREN, 0))
		})
		delBtn := widget.NewButton("÷", func() {
			update(tokens.T(tokens.DIV, 0))
		})

		basic.Add(clearBtn)
		basic.Add(lpBtn)
		basic.Add(rpBtn)
		basic.Add(delBtn)

		{
			for i := 7; i <= 9; i++ {
				str := fmt.Sprintf("%d", i)
				basic.Add(widget.NewButton(str, func() {
					currNum += str
					current.Text += str
					current.Refresh()
				}))
			}
			multBtn := widget.NewButton("x", func() {
				update(tokens.T(tokens.MUL, 0))
			})
			basic.Add(multBtn)

			for i := 4; i <= 6; i++ {
				str := fmt.Sprintf("%d", i)
				basic.Add(widget.NewButton(str, func() {
					currNum += str
					current.Text += str
					current.Refresh()
				}))
			}
			subBtn := widget.NewButton("-", func() {
				update(tokens.T(tokens.SUB, 0))
			})
			basic.Add(subBtn)

			for i := 1; i <= 3; i++ {
				str := fmt.Sprintf("%d", i)
				basic.Add(widget.NewButton(str, func() {
					currNum += str
					current.Text += str
					current.Refresh()
				}))
			}
			addBtn := widget.NewButton("+", func() {
				update(tokens.T(tokens.ADD, 0))
			})
			basic.Add(addBtn)

			str := "0"
			basic.Add(widget.NewButton(str, func() {
				if currNum != "0" {
					currNum += "0"
					current.Text += str
					current.Refresh()
				}
			}))

			basic.Add(widget.NewButton("", func() {
			}))

			basic.Add(widget.NewButton(".", func() {
				if !strings.Contains(currNum, ".") {
					currNum += "."
					current.Text += "."
					current.Refresh()
				}
			}))

			basic.Add(widget.NewButton("=", func() {
				if currNum != "" {
					if currNum[len(currNum)-1] == []byte(".")[0] {
						currNum += "0"
					}
					val, err := strconv.ParseFloat(currNum, 64)
					if err != nil {
						panic(err)
					}
					add(tokens.T(tokens.NUMBER, val))
				}

				val, err := parser.Eval(parser.Shunt(currTokens))
				if err != nil {
					panic(err)
				}
				currTokens = make([]tokens.Token, 0)
				current.Text = ""
				current.Refresh()
				result.Text = round(val)
				result.Refresh()
			}))
		}
	}

	advanced := container.New(layout.NewGridLayout(4))
	{
		square := widget.NewButton("a^2", func() {
			update(tokens.T(tokens.EXP, 0))
			update(tokens.T(tokens.NUMBER, 2))
		})
		exp := widget.NewButton("a^b", func() {
			update(tokens.T(tokens.EXP, 0))
		})
		log := widget.NewButton("log", func() {
			update(tokens.T(tokens.LOG, 0))
		})
		ln := widget.NewButton("ln", func() {
			update(tokens.T(tokens.LN, 0))
		})
		pi := widget.NewButton("π", func() {
			update(tokens.T(tokens.PI, 0))
		})
		e := widget.NewButton("e", func() {
			update(tokens.T(tokens.E, 0))
		})
		sin := widget.NewButton("sin", func() {
			update(tokens.T(tokens.SIN, 0))
		})
		cos := widget.NewButton("cos", func() {
			update(tokens.T(tokens.COS, 0))
		})
		tan := widget.NewButton("tan", func() {
			update(tokens.T(tokens.TAN, 0))
		})
		asin := widget.NewButton("asin", func() {
			update(tokens.T(tokens.ARCSIN, 0))
		})
		acos := widget.NewButton("acos", func() {
			update(tokens.T(tokens.ARCCOS, 0))
		})
		atan := widget.NewButton("atan", func() {
			update(tokens.T(tokens.ARCTAN, 0))
		})

		advanced.Add(square)
		advanced.Add(exp)
		advanced.Add(log)
		advanced.Add(ln)
		advanced.Add(pi)
		advanced.Add(e)
		advanced.Add(sin)
		advanced.Add(cos)
		advanced.Add(tan)
		advanced.Add(asin)
		advanced.Add(acos)
		advanced.Add(atan)
	}

	buttons := container.New(layout.NewHBoxLayout())
	buttons.Add(advanced)
	buttons.Add(basic)

	window.SetContent(container.New(layout.NewVBoxLayout(), display, buttons))
	window.ShowAndRun()
}

func round(num float64) string {
	return strconv.FormatFloat(num, 'f', -1, 64)
}
