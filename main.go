package main

import (
	"fmt"
	"net/http"
	"unicorn/unicorn"
)

func main() {
	go unicorn.UnicornProducer()

	unicorn.Routes()

	fmt.Println("Server running on :8888")
	http.ListenAndServe(":8888", nil)
}
