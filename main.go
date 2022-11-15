package main

import (
	"fmt"
	"image/color"
	"strconv"

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
		current.Text = tokens.SPrintArr(currTokens)
		current.Refresh()
	}

	update := func(token tokens.Token) {
		if currNum != "" {
			val, err := strconv.ParseFloat(currNum, 64)
			if err != nil {
				panic(err)
			}
			add(tokens.T(tokens.NUMBER, val))
		}
		add(token)
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
		delBtn := widget.NewButton("DEL", func() {
			length := len(currTokens)
			if length > 0 {
				currTokens = currTokens[:length-1]
				currNum = ""
			}
			refresh()
		})

		basic.Add(clearBtn)
		basic.Add(lpBtn)
		basic.Add(rpBtn)
		basic.Add(delBtn)

		for i := 1; i <= 9; i++ {
			str := fmt.Sprintf("%d", i)
			basic.Add(widget.NewButton(str, func() {
				currNum += str
				current.Text = tokens.SPrintArr(currTokens) + currNum
				current.Refresh()
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
