package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Item struct {
	IsProductive bool   `json:"IsProductive"`
	Start        string `json:"Start"`
	Stop         string `json:"Stop"`
	Title        string `json:"Title"`
}

func main() {
	// API request
	apiURL := "http://localhost:3000/"
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Erro ao fazer a solicitação HTTP:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler o corpo da resposta:", err)
		return
	}
	// fmt.Println("JSON recebido:", string(body))

	//  JSON Decoding
	var items []Item
	err = json.Unmarshal(body, &items)
	if err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		return
	}

	// Logic Here
	var totalTimeIsProductive, totalTimeNoProductive time.Duration
	var totalTime, lastStartTime, tempoDecorrido time.Duration
	var lastIsProductive bool
	for i, item := range items {
		initialTime, err := time.ParseDuration(item.Start)
		if err != nil {
			fmt.Printf("Erro ao analisar o tempo inicial para %s: %s\n", item.Title, err)
			continue
		}

		if i == 0 {
			lastStartTime = initialTime
			lastIsProductive = item.IsProductive
		} else {
			tempoDecorrido = initialTime - lastStartTime
			totalTime += tempoDecorrido
			lastStartTime = initialTime
			if lastIsProductive {
				totalTimeIsProductive += tempoDecorrido
			} else {
				totalTimeNoProductive += tempoDecorrido
			}
			lastIsProductive = item.IsProductive
		}
	}
	fmt.Println("Productive Time:", totalTimeIsProductive)
	fmt.Println("No Productive Time:", totalTimeNoProductive)
	fmt.Println("Total Time:", totalTime)
}
