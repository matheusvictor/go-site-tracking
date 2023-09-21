package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const quantidadeMonitoramento = 2
const delay = 60

func exibirIntroducao() {
	var nome string = "Matheus"
	sobrenome := "Santana" // forma "abreviada" de declarar uma variável com inferência de tipo

	var versao float32 = 1.1
	fmt.Println("Olá,", nome, sobrenome, "!")
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

func testarSite(site string) {
	response, _ := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Println("Site", site, "foi acessado com sucesso!")
	} else {
		fmt.Println("Site", site, "está com problemas! Status Code:", response.StatusCode)
	}
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	// Slices são como arrays dinâmicos
	sites := []string{
		"https://www.alura.com.br",
		"https://www.google.com.br",
		"https://random-status-code.herokuapp.com/"}

	for index := 1; index <= quantidadeMonitoramento; index++ {
		fmt.Println("Executando", index, "ª rotina de monitoramento...")
		//  index, site := range sites
		for index := 0; index < len(sites); index++ {
			testarSite(sites[index])
		}
		fmt.Println("Aguardando para testar novamente...")
		time.Sleep(delay * time.Second)
	}

}

func main() {
	exibirIntroducao()

	for {
		exibirMenu()
		opcao := lerOpcao()
		switch opcao {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Opção inválida! Encerrando o programa.")
			os.Exit(-1)
		}
	}
}
