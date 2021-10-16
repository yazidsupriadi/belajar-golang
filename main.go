package main

import (
	"log"
	"net/http"
	"webgolang/handler"
)

func main() {
	mux := http.NewServeMux()

	//list route fungsi
	mux.HandleFunc("/", handler.HomeFunc)
	mux.HandleFunc("/hello", handler.HelloFunc)
	mux.HandleFunc("/test", handler.TestFunc)
	mux.HandleFunc("/product", handler.ProductFunc)
	mux.HandleFunc("/get-post", handler.GetPost)
	mux.HandleFunc("/form", handler.Form)
	mux.HandleFunc("/proses", handler.Proses)
	log.Println("starting web on 9090")
	//inisiasi port
	err := http.ListenAndServe(":9090", mux)
	log.Fatal(err)
}
