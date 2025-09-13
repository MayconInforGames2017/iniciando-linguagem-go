package slicesGo

import "fmt"

func GavetasSlicesGo() {
	var gavetas []string
	gavetas = append(gavetas, "Copos", "Pratos", "Colher")
	fmt.Println("Adicionando objetos na gaveta ", gavetas)
	fmt.Println("Exibindo objeto na gaveta 3 ", gavetas[2])
}