package main

import (
	"elma_hw/internal/services"
	"net/http"
)

func main() {

	server := services.New()
	err := http.ListenAndServe("localhost:3000", server.BuildRoutes())
	if err != nil {
		panic(err)
	}

}
