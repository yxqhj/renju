package main

import (
	"renju/gobangIcon"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
)

func main() {
	a := app.New()
	//a.Settings().SetTheme(theme.LightTheme())
	win := a.NewWindow("qianhejun")

	win.SetContent(gobangIcon.NewChessboard())

	win.Resize(fyne.NewSize(540, 540))
	win.SetPadded(false)
	win.SetFixedSize(true)
	win.CenterOnScreen()
	win.ShowAndRun()
}
