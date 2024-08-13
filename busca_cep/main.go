package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Uso: go run main.go <URL>")
		os.Exit(1)
	}

	for _, cep := range os.Args[1:] {
		req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao buscar o CEP: %v\n", err)
			continue
		}
		defer req.Body.Close()

		if req.StatusCode != http.StatusOK {
			fmt.Fprintf(os.Stderr, "Erro: Status da resposta não OK (%d)\n", req.StatusCode)
			continue
		}

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler o corpo da resposta: %v\n", err)
			continue
		}

		var data ViaCep
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao converter o JSON: %v\n", err)
			continue
		}

		if data.Cep == "" {
			fmt.Printf("CEP %s não encontrado\n", cep)
			continue
		}
		
		file, err := os.Create(data.Localidade + ".json")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar o arquivo: %v\n", err)
			continue
		}

		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf("CEP: %s\nLogradouro: %s\nComplemento: %s\nUnidade: %s\nBairro: %s\nLocalidade: %s\nUF: %s\nIBGE: %s\nGIA: %s\nDDD: %s\nSIAFI: %s\n", data.Cep, data.Logradouro, data.Complemento, data.Unidade, data.Bairro, data.Localidade, data.Uf, data.Ibge, data.Gia, data.Ddd, data.Siafi))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao escrever no arquivo: %v\n", err)
			continue
		}
		fmt.Printf("CEP %s encontrado e salvo em %s.json\n", cep, data.Localidade)
	}
}
