package som

import (
	"fmt"
	"strings"
)

func Textos()  {
	var hello string = "Olá, mundo!"
	var question string = "Como vai?"

	// Concatenação
	var meet = hello + question
	// Deixando a frase maiscula
	fmt.Println(meet)
	fmt.Println(strings.ToUpper((meet)))
	// Verificando se a palavra existe
	fmt.Println(strings.Contains(meet, "mundo"))
}