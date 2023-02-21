package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	apiCepChannel := make(chan string)
	viaCepChannel := make(chan string)

	go func() {
		url := "https://cdn.apicep.com/file/apicep/88240-000.json"
		apiCepChannel <- request(url)
	}()

	go func() {
		url := "http://viacep.com.br/ws/88240000/json"
		viaCepChannel <- request(url)
	}()

	select {
	case zipInformation := <-apiCepChannel:
		fmt.Printf("Received from Api Cep: %s\n", zipInformation)
	case zipInformation := <-viaCepChannel:
		fmt.Printf("Received from Via Cep: %s\n", zipInformation)
	case <-time.After(time.Second):
		println("timeout")
	}
}

func request(url string) string {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}
