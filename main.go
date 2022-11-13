package main

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	app := app.New()
	window := app.NewWindow("Box Layout")

	current := canvas.NewText("Hello", color.White)
	result := canvas.NewText("There", color.White)
	display := container.New(layout.NewVBoxLayout(), current, result)

	text4 := canvas.NewText("centered", color.White)
	centered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), text4, layout.NewSpacer())
	window.SetContent(container.New(layout.NewVBoxLayout(), display, centered))
	window.ShowAndRun()
}
