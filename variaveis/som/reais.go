package som

import "fmt"

func CalcularCirculo() {
	var floatNumber float32 = 1.1
	var pi float64 = 3.14
	var raio float64 = 2.5
	var area = pi * raio * raio

	fmt.Println("FloatNumber: ", floatNumber)
	fmt.Println("Área do círculo é: ", area)
}