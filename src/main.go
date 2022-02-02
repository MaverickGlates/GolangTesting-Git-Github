package main

import "fmt"

func main() {
	//Declaración de constantes
	const pi float64 = 3.1416
	const pi2 float64 = 3.14163
	fmt.Println(pi, pi2)

	// Declaración de variables
	base := 12
	var altura int = 14
	var area int
	fmt.Println(base, altura, area)

	// Zero value
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Área de un cuadrado
	const squarebase = 10
	squareArea := squarebase * squarebase
	fmt.Println("Área Cuadrado:", squareArea)

	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("suma:", result)

	// Resta
	result = x - y
	fmt.Println("Resta:", result)

	// Multiplicación
	result = x * y
	fmt.Println("Multiplicación:", result)

	// División
	result = x / y
	fmt.Println("División:", result)

	// Modulo
	result = x % y
	fmt.Println("Modulo:", result)

	// Incremental

	x++
	fmt.Println("incremental:", x)

	// Decremental

	x--
	fmt.Println("Decremental:", x)

	// Área de un trapecio

	var baseTrapecio float64 = 4
	var base2Trapecio float64 = 3
	var h float64 = 3
	areaTrapecio := (((baseTrapecio + base2Trapecio) / 2) / h)
	areaTrapecio++ // Incremental
	fmt.Println("Área Trapecio: ", areaTrapecio)

	//Área de un rectangulo

	var bRectangulo = 10
	var aRectangulo = 20
	areaRectangulo := (bRectangulo + aRectangulo)
	fmt.Println(areaRectangulo)

}
