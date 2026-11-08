package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	sendGet()
	sendPost()
}
func sendGet() {
	get, err := http.Get("http://localhost:8080/hello")
	if err != nil {
		fmt.Print("\ninvalid request")
		return
	}
	defer get.Body.Close()
	body, err := io.ReadAll(get.Body)
	if err != nil {
		fmt.Print("\ninvalid body")
		return
	}
	fmt.Printf("Status: %d\n", get.StatusCode)
	fmt.Printf("Type of request: GET\n")
	fmt.Printf("Body: %s\n", body)
}

func sendPost() {
	data := `{"name":"Kenshin","weapon":"katana"}`

	post, err := http.Post("http://localhost:8080/samurai", "application/json", strings.NewReader(data))
	if err != nil {
		fmt.Print("\ninvalid request")
		return
	}
	defer post.Body.Close()
	body, err := io.ReadAll(post.Body)

	fmt.Printf("\nStatus: %d\n", post.StatusCode)
	fmt.Printf("Type of request: POST\n")
	fmt.Printf("Body: %s\n", body)

}
