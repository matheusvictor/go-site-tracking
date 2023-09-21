package main

import "fmt"

func menu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func main() {
	var nome string = "Matheus"
	sobrenome := "Santana" // forma "abreviada" de declarar uma variável com inferência de tipo

	var versao float32 = 1.1
	fmt.Println("Olá,", nome, sobrenome)
	fmt.Println("Este programa está na versão", versao)
	menu()

	var opcao int
	fmt.Scanf("%d", &opcao) // o valor será atributo à variável a partir de seu endereço
	// fmt.Println("Endereço da variável opcao: ", &opcao)
	fmt.Println("A opção escolhida foi", opcao)

	switch opcao {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa")
	default:
		fmt.Println("Opção inválida!")
	}

}
