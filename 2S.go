// --------EJERCICIO2: SECUENCIAL--------
// IMPORTANDO LIBRERIAS NECESARIAS
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	//VARIABLE DEL FADE (BRILLO) ENTRE IMÁGENES / VALORES ENTRE 0 - 1 (PROBAR: 0.25 - 0.50 - 0.75)
	FADE := 0.25
	//SE REGISTRA LA RUTA DE LA PRIMERA IMÁGEN
	IMPORTAR_IMAGEN1 := "src/OG/BEAUTY_BEHIND_THE_MADNESS.jpg"
	f, err := os.Open(IMPORTAR_IMAGEN1)
	check(err)

	//SE CAPTURAN LOS VALORES DE LA 1RA IMAGEN
	img, _, err := image.Decode(f)

	//SE REGISTRA LA RUTA DE LA SEGUNDA IMÁGEN
	IMPORTAR_IMAGEN2 := "src/OG/STARBOY.jpg"
	f2, err2 := os.Open(IMPORTAR_IMAGEN2)
	check(err2)

	//SE CAPTURAN LOS VALORES DE LA 2DA IMAGEN
	img2, _, err := image.Decode(f2)

	//SE RECONOCE EL TAMAÑO DE CADA IMAGEN
	TAMAÑO := img.Bounds().Size()
	rect := image.Rect(0, 0, TAMAÑO.X, TAMAÑO.Y)
	wImg := image.NewRGBA(rect)

	start := time.Now()
	// loop though all the x

	for ALTO := 0; ALTO < TAMAÑO.Y; ALTO++ {
		for ANCHO := 0; ANCHO < TAMAÑO.X; ANCHO++ {
			pixel := img.At(ANCHO, ALTO)
			pixel2 := img2.At(ANCHO, ALTO)

			OG_COLOR1 := color.RGBAModel.Convert(pixel).(color.RGBA)
			OG_COLOR2 := color.RGBAModel.Convert(pixel2).(color.RGBA)

			// CONVIRTIENDO ANÁLISIS RGB A VALORES FLOAT (IMÁGEN 1 Y 2)
			RED_CHANNEL := float64(OG_COLOR1.R)
			GREEN_CHANNEL := float64(OG_COLOR1.G)
			BLUE_CHANNEL := float64(OG_COLOR1.B)

			RED_CHANNEL2 := float64(OG_COLOR2.R)
			GREEN_CHANNEL2 := float64(OG_COLOR2.G)
			BLUE_CHANNEL2 := float64(OG_COLOR2.B)

			// REALIZANDO MEZCLA FADE CON LA FÓRMULA ESTABLECIDA EN EL EJERCICIO2
			RED_CHANNEL3 := uint8(FADE*RED_CHANNEL + (1-FADE)*RED_CHANNEL2)
			GREEN_CHANNEL3 := uint8(FADE*GREEN_CHANNEL + (1-FADE)*GREEN_CHANNEL2)
			BLUE_CHANNEL3 := uint8(FADE*BLUE_CHANNEL + (1-FADE)*BLUE_CHANNEL2)
			COLOR := color.RGBA{
				R: RED_CHANNEL3, G: GREEN_CHANNEL3, B: BLUE_CHANNEL3, A: OG_COLOR1.A,
			}
			wImg.Set(ANCHO, ALTO, COLOR)
		}
	}

	//FINALIZAR CRONÓMETRO E IMPRIMIR EL TIEMPO
	elapsed := time.Since(start)
	print("Se guardó correctamente la nueva imágen en la carpeta: 'OUTPUT'")
	log.Printf("\nTIEMPO | EJERCICIO2 | SECUENCIAL: %s", elapsed)

	//DETERMINANDO CARPETA DONDE SE GUARDARÁ EL RESULTADO
	imgOut := "src/OUT/"

	//CREAR EL RESULTADO
	ext := filepath.Ext(IMPORTAR_IMAGEN1)
	name := strings.TrimSuffix(filepath.Base("OUTPUT"), ext)
	newImagePath := fmt.Sprintf("%s/%s_SECUENCIAL-EJ2%s", filepath.Dir(imgOut), name, ext)
	fg, err := os.Create(newImagePath)
	defer fg.Close()
	check(err)
	err = jpeg.Encode(fg, wImg, nil)
	check(err)
}

// POR SI HAY ERRORES
func check(err error) {
	if err != nil {
		panic(err)
	}
}
