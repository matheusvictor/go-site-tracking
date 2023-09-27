package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
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
		case 1: // Iniciar monitoramento
			iniciarMonitoramento()
		case 2: // Cadastrar novo site para monitoramento
			var nomeSite string = ""
			var arquivoSites *os.File = nil

			arquivoSites = abrirArquivo("sites")
			if arquivoSites == nil {
				fmt.Println("Arquivo de sites não encontrando! Criando arquivo...")
				criarArquivo("sites")
				fmt.Println("Arquivo criado com sucesso!")
			}

			fmt.Print("Nome do site: https://www.")
			fmt.Scan(&nomeSite)
			registrarSite(nomeSite)

		case 3: // Exibir logs
			var arquivoLogs *os.File = nil
			arquivoLogs = abrirArquivo("logs")

			if arquivoLogs == nil {
				fmt.Println("Arquivo de logs não encontrando! Criando arquivo...")
				criarArquivo("logs")
				fmt.Println("Arquivo criado com sucesso!")
			}

			fmt.Println("Exibindo logs...")
			fmt.Println("================")
			exibirLogs()
			fmt.Println("================")

		case 0: // Sair do programa
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

func iniciarMonitoramento() {
	fmt.Println("Monitorando sites...")

	// Slices são como arrays dinâmicos
	sites := lerSitesDoArquivo()

	if len(sites) == 0 {
		fmt.Println("Não há sites para serem monitorados!")
		return
	}

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

func testarSite(site string) {

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	response, erro := client.Get(site)

	if os.IsTimeout(erro) {
		fmt.Printf("Ocorreu um erro de timeout: %v\n", erro)
		registrarLog(site, false)
		return
	} else if erro != nil {
		fmt.Println("Ocorreu um erro:", erro)
	}

	if response.StatusCode == 200 {
		fmt.Println("Site", site, "foi acessado com sucesso!")
		registrarLog(site, true)
	} else {
		fmt.Println("Site", site, "está com problemas! Status Code:", response.StatusCode)
		registrarLog(site, false)
	}
}

func lerSitesDoArquivo() []string {

	var sites []string
	arquivo := abrirArquivo("sites")

	if arquivo != nil {
		leitor := bufio.NewReader(arquivo)
		for {
			linha, erro := leitor.ReadString('\n') // aspas simples para definir o byte
			linha = strings.TrimSpace(linha)

			// se a primeira linha for uma string vazia, não há sites para monitorar
			if len(linha) == 0 {
				break
			}

			sites = append(sites, linha)

			if erro == io.EOF { // identifica final do arquivo
				break
			}
		}
		arquivo.Close()
	}
	return sites
}

func abrirArquivo(nomeArquivo string) *os.File {
	arquivo, erro :=
		os.OpenFile(fmt.Sprintf("%s.%s", nomeArquivo, "txt"), os.O_RDWR|os.O_APPEND, 0666)

	if erro != nil {
		fmt.Println("Ocorreu um erro:", erro.Error())
		return nil
	}
	return arquivo
}

func criarArquivo(nomeArquivo string) {
	novoArquivo, erro := os.OpenFile(
		fmt.Sprintf("%s.%s", nomeArquivo, "txt"),
		os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if erro != nil {
		fmt.Println("Ocorreu um erro:", erro)
		return
	}
	fmt.Println("Arquivo", novoArquivo, "criado com sucesso!")
}

func registrarSite(site string) {
	arquivoSites := abrirArquivo("sites")

	if !strings.Contains(site, ".com") {
		fmt.Println("Não contém .com")
		arquivoSites.WriteString("https://www." + site + ".com\n")
	} else {
		arquivoSites.WriteString("https://www." + site + "\n")
	}

	arquivoSites.Close()
}

func registrarLog(site string, status bool) {
	var arquivoLogs *os.File = nil
	nomeArquivo := "logs"

	arquivoLogs = abrirArquivo(nomeArquivo)
	if arquivoLogs == nil {
		criarArquivo(nomeArquivo)
	}

	arquivoLogs = abrirArquivo(nomeArquivo)
	arquivoLogs.WriteString("Site: " + site + " [Online: " + strconv.FormatBool(status) + "] [" + time.Now().Format("02/01/2006 15:04:05") + "] \n")
	arquivoLogs.Close()
}

func exibirLogs() {
	arquivoLogs, erro := os.ReadFile("logs.txt")

	if erro != nil {
		fmt.Println(erro)
	}

	fmt.Println(string(arquivoLogs))
}
