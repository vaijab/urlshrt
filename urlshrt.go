package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type gogl struct {
	Kind    string
	Id      string
	LongUrl string
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Please specify a URL to shorten.")
		os.Exit(0)
	}

	longUrl := os.Args[len(os.Args)-1]
	apiUrl := "https://www.googleapis.com/urlshortener/v1/url"
	body := strings.NewReader(fmt.Sprintf(`{"longUrl": "%s"}`, longUrl))

	req, err := http.NewRequest("POST", apiUrl, body)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()

	var result gogl
	err = json.Unmarshal(b, &result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", result.Id)
}
