package main

import "fmt"

func normalfunction(x string) {
	fmt.Println(x)
}

func tripeArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func main() {
	//Const Declaration
	const pi float64 = 3.1416
	const pi2 float64 = 3.14163
	fmt.Println(pi, pi2)

	// Variable Declaration
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

	// Square area
	const squarebase = 10
	squareArea := squarebase * squarebase
	fmt.Println("Área Cuadrado:", squareArea)

	x := 10
	y := 50

	// Sum
	result := x + y
	fmt.Println("suma:", result)

	// subtraction
	result = x - y
	fmt.Println("Resta:", result)

	// Multiplication
	result = x * y
	fmt.Println("Multiplicación:", result)

	// Split
	result = x / y
	fmt.Println("División:", result)

	// Module
	result = x % y
	fmt.Println("Modulo:", result)

	// Incremental

	x++
	fmt.Println("incremental:", x)

	// Decremental

	x--
	fmt.Println("Decremental:", x)

	// Trapeze area

	var baseTrapecio float64 = 4
	var base2Trapecio float64 = 3
	var h float64 = 3
	areaTrapecio := (((baseTrapecio + base2Trapecio) / 2) / h)
	areaTrapecio++ // Incremental
	fmt.Println("Área Trapecio: ", areaTrapecio)

	// Rectangle area
	var bRectangulo = 30
	var aRectangulo = 20
	areaRectangulo := (bRectangulo + aRectangulo)
	fmt.Println(areaRectangulo)

	normalfunction("asd")
	tripeArgument(1, 2, "hola mundo")

	// For conditional
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	// For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// For forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}
}
