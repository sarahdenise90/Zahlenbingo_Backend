package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type payloadModel struct {
	Name string
}

func testEndpoint(response http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var payload payloadModel
	err := decoder.Decode(&payload)
	if err != nil {
		panic(err)
	}
	fmt.Println(payload.Name)
	defer request.Body.Close()
	response.Header().Set("Content-Type", "text/plain")
	io.WriteString(response, "Hello Go World!\n")
}

func main() {
	fmt.Println("startup")
	http.HandleFunc("/test2", testEndpoint)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
			panic(err)
	}
}
