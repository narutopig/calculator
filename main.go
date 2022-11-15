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
	test()
	app := app.New()
	window := app.NewWindow("Calculator")

	currTokens := make([]tokens.Token, 0)
	currNum := ""

	display := container.New(layout.NewVBoxLayout())

	// needed to update stuff
	current := canvas.NewText("", color.White)
	result := canvas.NewText("There", color.White)

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

			basic.Add(widget.NewButton("C", func() {
				length := len(currTokens)
				if length > 0 {
					currTokens = currTokens[:length-1]
					currNum = ""
					current.Text = ""
				}
				refresh()
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
				for _, t := range currTokens {
					fmt.Println(t)
				}
				val, err := parser.Eval(parser.Shunt(currTokens))
				if err != nil {
					panic(err)
				}
				result.Text = fmt.Sprintf("%f", val)
				result.Refresh()
			}))
		}
	}

	window.SetContent(container.New(layout.NewVBoxLayout(), display, basic))
	window.ShowAndRun()
}

func test() {
	t := []tokens.Token{tokens.T(tokens.NUMBER, 3), tokens.T(tokens.ADD, 0), tokens.T(tokens.NUMBER, 4)}

	shunt := parser.Shunt(t)

	fmt.Println(parser.Eval(shunt))
}
