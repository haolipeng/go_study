package main

import (
	"go_study/filestore_server/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/file/upload/suc", handler.UploadFileSuccess)
	http.HandleFunc("/file/upload", handler.UploadFile)
	http.ListenAndServe(":9000", nil)
}
