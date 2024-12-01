package main

import (
	"log"
	"net/http"

	"github.com/infernus01/fileService/pkg/filehandler"
)

func main() {
	http.HandleFunc("/list", filehandler.HandleListFile)
	http.HandleFunc("/add", filehandler.HandleAddFile)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}
