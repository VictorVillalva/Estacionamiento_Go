package models

import (
	"fyne.io/fyne/v2/canvas"
	"sync"
)

type Parking struct {
	EspaciosParking chan bool
	ColocarCarro    chan *canvas.Image
	M               sync.Mutex
}

func CrearParking(nP int) *Parking {
	return &Parking{
		EspaciosParking: make(chan bool, nP+1),
		ColocarCarro:    make(chan *canvas.Image, 1),
	}
}
