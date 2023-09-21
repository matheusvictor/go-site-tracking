package main

import "fmt"

func exibirIntroducao() {
	var nome string = "Matheus"
	sobrenome := "Santana" // forma "abreviada" de declarar uma variável com inferência de tipo

	var versao float32 = 1.1
	fmt.Println("Olá,", nome, sobrenome)
	fmt.Println("Este programa está na versão", versao)
}

func lerOpcao() int {
	var opcao int
	fmt.Scan(&opcao) // o valor será atributo à variável a partir de seu endereço e inferência de tipo
	return opcao
}

func exibirMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func main() {
	exibirIntroducao()
	exibirMenu()
	opcao := lerOpcao()

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
