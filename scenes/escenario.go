package scenes

import (
	"Proyecto2/models"
	"Proyecto2/poison"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
)

type Escenario struct {
	window    fyne.Window
	contenido *fyne.Container
}

func (e *Escenario) RenderAutomovil(a *models.Parking) {
	for {
		autoImage := a.ColocarAutomovil
		e.contenido.Add(autoImage)
		e.window.Canvas().Refresh(e.contenido)

	}
}
func (e *Escenario) IniciarSimulacion() {
	n := models.CrearParking(20)
	go poison.MetodoPoison(100, n)
	go e.RenderAutomovil(n)

}
func (e *Escenario) Renderizado() {
	bgImagen := canvas.NewImageFromURI(storage.NewFileURI("./assets/Parking.png"))
	bgImagen.Resize(fyne.NewSize(800, 600))
	bgImagen.Move(fyne.NewPos(0, 0))

	e.contenido = container.NewWithoutLayout(bgImagen)
	e.window.SetContent(e.contenido)
	e.IniciarSimulacion()
}
func NuevaEscena(window fyne.Window) *Escenario {
	escena := &Escenario{window: window}
	escena.Renderizado()
	return escena
}
