package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default if PORT not set
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from DeepLX on Railway!")
	})

	fmt.Println("Starting server on port " + port)
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		panic(err)
	}
}
