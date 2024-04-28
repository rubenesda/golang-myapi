package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func handleStatus(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Server in running\n")
}

func main() {
	port := ":8082"

	http.HandleFunc("/status", handleStatus)

	err := http.ListenAndServe(port, nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed\n")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}