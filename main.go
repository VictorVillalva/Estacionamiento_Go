package main

import (
	"Proyecto2/scenes"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Parking de millonarios")

	w.CenterOnScreen()
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(800, 600))
	scenes.NuevaEscena(w)
	w.ShowAndRun()
}
