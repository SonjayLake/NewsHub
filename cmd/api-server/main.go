package main

import (
	"log"
	"net/http"
)


func main(){
	mux := http.NewServeMux()
	
	mux.HandleFunc("/",root)
	mux.HandleFunc("/healthcheck",healthCheck)

	log.Print("Listening on port 9000")
	err := http.ListenAndServe(":9000",mux)
	log.Fatal(err)
	
}