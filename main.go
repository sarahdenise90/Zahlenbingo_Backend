package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	io.WriteString(response, "Hello World!")
}

func main() {
	fmt.Println("startup")
	http.HandleFunc("/test2", testEndpoint)
	http.ListenAndServe(":8082", nil)
}
