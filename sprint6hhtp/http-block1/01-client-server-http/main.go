package main

import (
	"fmt"
	"io"
	"net/http"
)

func mainHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "invalid method", http.StatusMethodNotAllowed)
		return
	}
	res.Write([]byte("Hello, Kenshin!"))
}

func samuraiHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "invalid method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "invalid body", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	res.WriteHeader(http.StatusCreated)
	res.Write([]byte("Samurai created!"))
	fmt.Printf("Status: %d", http.StatusCreated)
	fmt.Print("Body: %s", body)

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", mainHandler)
	mux.HandleFunc("/samurai", samuraiHandler)
	fmt.Println("Server was launched on http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Printf("Server was not launched: %v", err)
	}

}
