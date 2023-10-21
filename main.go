//// @ei7hacker /// EIke 
//#~/go/bin
package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// Função para verificar a existência de um endpoint em todas as URLs
func scanEndpointOnAllURLs(baseURLs []string, endpoint string) {
	for _, baseURL := range baseURLs {
		fullURL := strings.TrimRight(baseURL, "/") + "/" + strings.TrimLeft(endpoint, "/")
		response, err := http.Head(fullURL)
		if err == nil && response.StatusCode == 200 {
			fmt.Printf("[%d] %s\n", response.StatusCode, fullURL)
		}
	}
}

func main() {
	sitesFile := flag.String("s", "", "Arquivo de lista de URLs")
	wordlistFile := flag.String("w", "", "Arquivo de lista de endpoints")
	flag.Parse()

	if *sitesFile == "" || *wordlistFile == "" {
		fmt.Println("É necessário especificar ambos os arquivos de entrada.")
		return
	}

	// Ler a lista de URLs dos arquivos
	sites, err := os.Open(*sitesFile)
	if err != nil {
		fmt.Printf("Erro ao abrir o arquivo de sites: %v\n", err)
		return
	}
	defer sites.Close()

	scanner := bufio.NewScanner(sites)
	var baseURLs []string
	for scanner.Scan() {
		baseURLs = append(baseURLs, scanner.Text())
	}

	wordlist, err := os.Open(*wordlistFile)
	if err != nil {
		fmt.Printf("Erro ao abrir o arquivo de wordlist: %v\n", err)
		return
	}
	defer wordlist.Close()

	scanner = bufio.NewScanner(wordlist)
	var endpoints []string
	for scanner.Scan() {
		endpoints = append(endpoints, scanner.Text())
	}

	for _, endpoint := range endpoints {
		scanEndpointOnAllURLs(baseURLs, endpoint)
	}
}

