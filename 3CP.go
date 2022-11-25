// --------EJERCICIO3: CONCURRENCIA Y PARALELISMO--------
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
	"sync"
	"time"
)

func main() {

	txt, err := os.Create("CANAL_ROJO.txt")
	if err != nil {
		fmt.Printf("error", err)
		return
	}
	txt1, err := os.Create("CANAL_VERDE.txt")
	if err != nil {
		fmt.Printf("error", err)
		return
	}
	txt2, err := os.Create("CANAL_AZUL.txt")
	if err != nil {
		fmt.Printf("error", err)
		return
	}
	//VARIABLE DEL FADE (BRILLO) ENTRE IMÁGENES / VALORES ENTRE 0 - 1 (PROBAR: 0.25 - 0.50 - 0.75)
	//SE REGISTRA LA RUTA DE LA PRIMERA IMÁGEN
	IMPORTAR_IMAGEN1 := "src/OG/ASTROWORLD.jpg"
	f, err := os.Open(IMPORTAR_IMAGEN1)
	check(err)

	//SE CAPTURAN LOS VALORES DE LA IMAGEN
	img, _, err := image.Decode(f)

	TAMAÑO := img.Bounds().Size()
	rect := image.Rect(0, 0, TAMAÑO.X, TAMAÑO.Y)
	wImg := image.NewRGBA(rect)

	//GOROUTINE
	wg := new(sync.WaitGroup)

	//TEMPORIZADOR: INICIAR
	start := time.Now()

	//CICLO QUE RECORRE TODO Y (ALTO)
	for ALTO := 0; ALTO < TAMAÑO.Y; ALTO++ {

		//GOROUTINE
		wg.Add(1)
		ALTO := ALTO
		go func() {

			// CICLO QUE RECORRE TODO X (ANCHO)
			for ANCHO := 0; ANCHO < TAMAÑO.X; ANCHO++ {
				pixel := img.At(ANCHO, ALTO)
				OG_COLOR := color.RGBAModel.Convert(pixel).(color.RGBA)

				// CONVIRTIENDO ANÁLISIS RGB A VALORES FLOAT (IMÁGEN 1 Y 2)
				RED_CHANNEL := float64(OG_COLOR.R)
				GREEN_CHANNEL := float64(OG_COLOR.G)
				BLUE_CHANNEL := float64(OG_COLOR.B)

				_, err = txt.WriteString(fmt.Sprintf("%d\n", uint8(RED_CHANNEL)))
				if err != nil {
					fmt.Printf("error", err)
				}
				_, err = txt1.WriteString(fmt.Sprintf("%d\n", uint8(GREEN_CHANNEL)))
				if err != nil {
					fmt.Printf("error", err)
				}
				_, err = txt2.WriteString(fmt.Sprintf("%d\n", uint8(BLUE_CHANNEL)))
				if err != nil {
					fmt.Printf("error", err)
				}
			}
			defer wg.Done()
		}()
	}
	wg.Wait()

	//FINALIZAR CRONÓMETRO E IMPRIMIR EL TIEMPO
	elapsed := time.Since(start)
	print("Se guardaron correctamente los txt con información de cada canal RGB en la carpeta: 'OUTPUT'")
	log.Printf("\nTIEMPO | EJERCICIO3 | GOROUTINE: %s", elapsed)

	//DETERMINANDO CARPETA DONDE SE GUARDARÁ EL RESULTADO
	imgOut := "src/OUT/"

	//CREAR EL RESULTADO
	ext := filepath.Ext(IMPORTAR_IMAGEN1)
	name := strings.TrimSuffix(filepath.Base("OUTPUT"), ext)
	newImagePath := fmt.Sprintf("%s/%s_CONCURRENCIA-EJ3%s", filepath.Dir(imgOut), name, ext)
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
