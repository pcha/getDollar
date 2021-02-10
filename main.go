package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const endpoint = "http://ws.geeklab.com.ar/dolar/get-dolar-json.php"

type dollarCotization struct {
	libre string `json: "libre"`
	blue string `json: "blue"`
}

func main() {
	res, err :=http.Get(endpoint)

	if err != nil {
		panic(err)
	}

	fmt.Println(res)

	decoder := json.NewDecoder(res.Body)
	dc := new(dollarCotization)
	decoder.Decode(&dc)

	fmt.Println(dc)
}