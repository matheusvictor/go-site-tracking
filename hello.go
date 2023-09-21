package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const quantidadeMonitoramento = 2
const delay = 60

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
	fmt.Println("2 - Cadastrar novo site para monitoramento")
	fmt.Println("3 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func testarSite(site string) {
	response, erro := http.Get(site)

	if erro != nil {
		fmt.Println("Ocorreu um erro:", erro)
	}

	if response.StatusCode == 200 {
		fmt.Println("Site", site, "foi acessado com sucesso!")
	} else {
		fmt.Println("Site", site, "está com problemas! Status Code:", response.StatusCode)
	}
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	// Slices são como arrays dinâmicos
	sites := lerSitesDoArquivo()

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

func abrirArquivo() *os.File {

	// arquivo, erro := ioutil.ReadFile("sites.txt")
	arquivo, erro := os.Open("sites.txt")

	if erro != nil {
		fmt.Println("Ocorreu um erro:", erro)
	}

	return arquivo
}

func criarArquivo() {
	os.Create("teste.txt")
}

func lerSitesDoArquivo() []string {

	var sites []string
	arquivo := abrirArquivo()

	leitor := bufio.NewReader(arquivo)
	for {
		linha, erro := leitor.ReadString('\n') // aspas simples para definir o byte
		linha = strings.TrimSpace(linha)
		fmt.Println(linha)

		sites = append(sites, linha)

		if erro == io.EOF { // identifica final do arquivo
			break
		}
	}

	arquivo.Close()

	return sites
}
