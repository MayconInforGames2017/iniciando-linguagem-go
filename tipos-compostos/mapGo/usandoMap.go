package mapGo

import "fmt"

func MapGo() {
	var pessoas = map[string]int{}
	pessoas["Lais"] = 26
	pessoas["Pedro"] = 3
	fmt.Println(pessoas)

	// Buscando a idade de Pedro
	fmt.Println(pessoas["Pedro"])

	// Verificando se existe uma pessoa
	if idade, ok := pessoas["Pedro"]; ok {
		fmt.Println("Essa pessoa existe no map ", idade, ok)
	} else {
		fmt.Println("Pessoa não existe no map")
	}

	// Removendo uma pessoa
	delete(pessoas, "Lais")
	fmt.Println(pessoas)
}