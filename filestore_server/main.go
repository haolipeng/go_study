package main

import (
	"go_study/filestore_server/handler"
	"net/http"
)

func main() {
	//路由规则
	http.HandleFunc("/file/upload/suc", handler.UploadFileSuccess)
	http.HandleFunc("/file/upload", handler.UploadFile)
	http.HandleFunc("/file/meta", handler.FileQuery)
	http.HandleFunc("/file/download", handler.DownloadFile)
	http.ListenAndServe(":9000", nil)
}
