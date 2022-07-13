package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://gorest.co.in/public/v1/posts?page=1")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var ar ApiResponse

	if err := json.Unmarshal(body, &ar); err != nil {
		panic(err)
	}

	fmt.Println(ar.Data)
}
