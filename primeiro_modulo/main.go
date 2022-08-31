package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
  "bufio"
  "io"
  "strings"
  "strconv"
  "io/ioutil"
)

const (
	monitoramentos = 3
	delay          = 5
)


func main() {
	exibeIntroducao()
  for {
    exibeMenu()
    comando := leComando()

    switch comando {
    case 1:
      fmt.Println("Monitorando...")
      iniciarMonitoramento()
    case 2:
      fmt.Println("Exibindo os nomes...")
      exibeNomes()
    case 3:
      fmt.Println("Imprimir logs...")
      imprimeLogs()
    case 0:
      fmt.Println("Saindo do programa")
      os.Exit(0)
    default:
      fmt.Println("Não conheço este comando")
      os.Exit(-1)
    }
  }
}

func exibeIntroducao() {
	var nome string
	versao := 1.2
	fmt.Println("Por gentileza, informe seu nome:")
	fmt.Scan(&nome)
	fmt.Println("Olá Sr(a)", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Nomes")
  fmt.Println("3- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
  fmt.Println("")

	return comandoLido
}

func iniciarMonitoramento() {
  // criando o slice a partir da função leSitesDoArquivo()
  sites := leSitesDoArquivo()

  if len(sites) > 0 {
    for i := 0; i < monitoramentos; i++ {
      for i, site := range sites {
        fmt.Println("Testando site", i, ":", site)
        testaSite(site)
      }
      time.Sleep(delay * time.Second)
      fmt.Println("")
    }
  } else {
    fmt.Println("Não localizamos os sites")
    fmt.Println("")
  }
}


func leSitesDoArquivo() []string {
  var sites []string
  arquivo, err := os.Open("sites.txt")

  if err != nil {
      fmt.Println("Ocorreu um erro:", err)
  }

  leitor := bufio.NewReader(arquivo)
  for {
      linha, err := leitor.ReadString('\n')
      linha = strings.TrimSpace(linha)
      sites = append(sites, linha)
      if err == io.EOF {
          break
      }
  }

  arquivo.Close()
  return sites
}

func testaSite(site string) {
  resp, err := http.Get(site)

  if err != nil {
      fmt.Println("Ocorreu um erro:", err)
  }

  if resp.StatusCode == 200 {
      fmt.Println("Site:", site, "foi carregado com sucesso!")
      registraLog(site, true)
  } else {
      fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
      registraLog(site, false)
  }
}

func registraLog(site string, status bool) {
  arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

  if err != nil {
      fmt.Println("Ocorreu um erro:", err)
  }

  arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + 
  " - online: " + strconv.FormatBool(status) + "\n")
  arquivo.Close()
}

func imprimeLogs() {
  arquivo, err := ioutil.ReadFile("log.txt")

  if err != nil {
      fmt.Println("Ocorreu um erro:", err)
  }

  fmt.Println(string(arquivo))
}

func exibeNomes() {
	nomes := []string{"Corinthians", "Karen", "Princesa"}
	fmt.Println("O meu slice tem", len(nomes), "itens")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")

	nomes = append(nomes, "Pipoca", "Emily")
	fmt.Println("O meu slice tem", len(nomes), "itens")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")

	for i := 0; i < len(nomes); i++ {
    fmt.Println("")
		fmt.Println("Eu amo a", nomes[i])
	}
}