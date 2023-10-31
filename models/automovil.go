package models

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"math/rand"
	"time"
)

type Automovil struct {
	Parking         *Parking
	Identificador   int
	posicionParking int
	modelo          *canvas.Image
}

// m -> modelo

func CrearAutomovil(p *Parking, m *canvas.Image) *Automovil {
	return &Automovil{
		Parking: p,
		modelo:  m,
	}
}

func (a *Automovil) MoverCarro() {
	a.Parking.EspaciosParking <- true
	a.Parking.M.Lock()

	for i := 0; i < len(a.Parking.EspaciosParking); i++ {
		if !a.Parking.LugaresParking[i].Ocupado {
			a.modelo.Move(fyne.NewPos(a.Parking.LugaresParking[i].PosicionX, a.Parking.LugaresParking[i].PosicionY))
			a.modelo.Refresh()

			//Ocupa el lugar
			a.posicionParking = i
			a.Parking.LugaresParking[i].Ocupado = true
			break
		}
	}

	fmt.Println("Automovil numero -> #", a.Identificador, " Acaba de entrar")
	time.Sleep(300 * time.Millisecond)
	// Desbloqueo de Mutex
	a.Parking.M.Unlock()

	TiempoEsperaTurno := rand.Intn(5-1+1) + 1
	time.Sleep(time.Duration(TiempoEsperaTurno) * time.Second)

	// Volvemos a bloquear el Mutex
	a.Parking.M.Lock()

	//Liberemos espacios de automoviles
	<-a.Parking.EspaciosParking
	a.Parking.LugaresParking[a.posicionParking].Ocupado = false
	a.modelo.Move(fyne.NewPos(0, 80))
	a.modelo.Refresh()
	fmt.Println("Automovil numero -> #", a.Identificador, " Acaba de salir")
	time.Sleep(300 * time.Millisecond)

	//Desbloquemos el Mutex
	a.Parking.M.Unlock()
}
