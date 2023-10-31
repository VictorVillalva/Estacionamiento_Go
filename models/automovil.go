package models

import (
	"fyne.io/fyne/v2/canvas"
	"math/rand"
)

type Automovil struct {
	Parking       *Parking
	Identificador int
	modelo        *canvas.Image
}

// m -> modelo
func CrearAutomivil(p *Parking, m *canvas.Image) *Automovil {
	return &Automovil{
		Parking: p,
		modelo:  m,
	}
}

func (a *Automovil) MoverCarro() {
	a.Parking.EspaciosParking <- true
	a.Parking.M.Lock()
	x := float32(rand.Intn(650 - 150 + 1))
	y := float32(rand.Intn(300 - 50 + 1))

}
