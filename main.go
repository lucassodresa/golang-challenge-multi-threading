package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

func logApiError(err error, apiUrl string) {
	errMessage := fmt.Sprintf("Error: %s from %s", err, apiUrl)
	fmt.Println(errMessage)
	log.Fatal(err)
}

func getFromApi(ctx context.Context, apiUrl string, ch chan<- string) {
	url, err := url.Parse(apiUrl)
	if err != nil {
		logApiError(err, apiUrl)
	}

	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, "GET", apiUrl, nil)
	if err != nil {
		logApiError(err, apiUrl)
	}

	res, err := client.Do(req)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return
		}
		logApiError(err, apiUrl)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		logApiError(err, apiUrl)
	}

	ch <- fmt.Sprintf("## Provided by: %s ##\n%s", url.Hostname(), body)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a postcode")
	}
	postcode := os.Args[1]

	addressInfo := make(chan string)
	ctx, cancel := context.WithCancel(context.Background())

	go getFromApi(ctx, "https://brasilapi.com.br/api/cep/v1/"+postcode, addressInfo)
	go getFromApi(ctx, "https://viacep.com.br/ws/"+postcode+"/json/", addressInfo)

	select {

	case res := <-addressInfo:
		fmt.Println(res)
		cancel()
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout")
	}
}
