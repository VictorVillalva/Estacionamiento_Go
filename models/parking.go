package models

import (
	"fyne.io/fyne/v2/canvas"
	"sync"
)

// Estructura del lugar
type Lugar struct {
	PosicionX float32
	PosicionY float32
	Ocupado   bool
}

// Estructura del estacionamiento
type Parking struct {
	EspaciosParking  chan bool
	LugaresParking   []Lugar
	ColocarAutomovil chan *canvas.Image
	M                sync.Mutex
}

func CrearParking(nP int) *Parking {
	return &Parking{
		EspaciosParking:  make(chan bool, nP+1),
		ColocarAutomovil: make(chan *canvas.Image, 100),

		//Lugares a ocupar dentro del estacionamiento
		LugaresParking: []Lugar{
			{PosicionX: 250, PosicionY: 50, Ocupado: false},
			{PosicionX: 350, PosicionY: 50, Ocupado: false},
			{PosicionX: 450, PosicionY: 50, Ocupado: false},
			{PosicionX: 550, PosicionY: 50, Ocupado: false},
			{PosicionX: 650, PosicionY: 50, Ocupado: false},

			{PosicionX: 250, PosicionY: 150, Ocupado: false},
			{PosicionX: 350, PosicionY: 150, Ocupado: false},
			{PosicionX: 450, PosicionY: 150, Ocupado: false},
			{PosicionX: 550, PosicionY: 150, Ocupado: false},
			{PosicionX: 650, PosicionY: 150, Ocupado: false},

			{PosicionX: 250, PosicionY: 300, Ocupado: false},
			{PosicionX: 350, PosicionY: 300, Ocupado: false},
			{PosicionX: 450, PosicionY: 300, Ocupado: false},
			{PosicionX: 550, PosicionY: 300, Ocupado: false},
			{PosicionX: 650, PosicionY: 300, Ocupado: false},

			{PosicionX: 250, PosicionY: 400, Ocupado: false},
			{PosicionX: 350, PosicionY: 400, Ocupado: false},
			{PosicionX: 450, PosicionY: 400, Ocupado: false},
			{PosicionX: 550, PosicionY: 400, Ocupado: false},
			{PosicionX: 650, PosicionY: 400, Ocupado: false},
		},
	}
}
