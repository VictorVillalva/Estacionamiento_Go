package poison

import (
	"Proyecto2/models"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
	"math/rand"
	"time"
)

func MetodoPoison(a int, parking *models.Parking) {
	parking.EspaciosParking <- true
	//Creamos la imagen del coche
	for i := 0; i < a; i++ {
		automovilImagen := canvas.NewImageFromURI(storage.NewFileURI("./assets/Ferrari.png"))
		automovilImagen.Resize(fyne.NewSize(100, 100)) //Size de imagen
		automovilImagen.Move(fyne.NewPos(120, 500))

		automovilNuevo := models.CrearAutomovil(parking, automovilImagen)
		automovilNuevo.Identificador = i + 1

		//Colocamos un coche nuevo
		parking.ColocarAutomovil <- automovilImagen
		time.Sleep(time.Millisecond * 200)
		go automovilNuevo.MoverCarro()
		time.Sleep(time.Duration(rand.ExpFloat64() * float64(time.Second)))

	}
}
