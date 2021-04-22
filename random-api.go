package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
)

func randNum(min, max int64) int64 {
	//Como rand.Int solo entrega números de entre [0,max) se resta a max
	dif := big.NewInt(max - min)

	//Int returns a uniform random value in [0, max). It panics if max <= 0.
	n, err := rand.Int(rand.Reader, dif)
	if err != nil {
		panic(err)
	}

	//se vuelve a sumar el minimo para neutralizar la resta del inicio.
	return n.Int64() + min
}

type helloWorldResponse struct {
	NumeroRandom int64 `json:"message"`
}

type helloRandomRequest struct {
	Minimo int64 `json:"minimo"`
	Maximo int64 `json:"maximo"`
}

func main() {
	port := 8080

	http.HandleFunc("/hellorandom", helloRandomHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloRandomHandler(w http.ResponseWriter, r *http.Request) {
	var request helloRandomRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	num_min := request.Minimo
	num_max := request.Maximo

	num_rand := randNum(num_min, num_max)

	response := helloWorldResponse{NumeroRandom: num_rand}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}
