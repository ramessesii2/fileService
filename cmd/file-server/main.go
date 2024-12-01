package main

import (
	"net/http"

	"github.com/infernus01/fileService/pkg/filehandler"
)

func main() {
	http.HandleFunc("/list", filehandler.HandleListFile)
	http.HandleFunc("/add", filehandler.HandleAddFile)
	http.ListenAndServe(":8080", nil)
}
