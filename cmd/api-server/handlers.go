package main

import (
	"fmt"
	"net/http"
)

func healthCheck(writer http.ResponseWriter, reader *http.Request){
	fmt.Println("Health check...")
}

func root(writer http.ResponseWriter, reader *http.Request){
	fmt.Println("Accessing root route...")
}