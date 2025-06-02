package main

import (
	"fmt"
	"net/http"
	"os"
	"unicorn-app/unicorn"
	"unicorn-app/utils"
)

func main() {
	if err := utils.LoadData(); err != nil {
		fmt.Fprintf(os.Stderr, "❌ Failed to load data files: %s\n", err)
		os.Exit(1)
	}

	go unicorn.UnicornProducer()

	unicorn.Routes()

	fmt.Println("Server running on :8888")
	http.ListenAndServe(":8888", nil)
}
