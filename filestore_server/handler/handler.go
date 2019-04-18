package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

const FILE_PATH = "/tmp"

func UploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//返回上传index.html页面
		data, err := ioutil.ReadFile("D:/goProject/src/go_study/filestore_server/static/view/index.html")
		if err != nil {
			io.WriteString(w, "Internel server error!")
			return
		}
		io.WriteString(w, string(data))

	} else if r.Method == "POST" {
		//接收数据流
		file, fileHeader, err := r.FormFile("file")
		fmt.Println("file size is ", fileHeader.Size)
		if err != nil {
			fmt.Printf("Failed to get data,err:%s\n", err.Error())
			return
		}
		defer file.Close()

		//创建缓存我额济纳
		newFile, err := os.Create(fileHeader.Filename)
		if err != nil {
			fmt.Printf("Failed to create file %s\n", err.Error())
			return
		}

		defer newFile.Close()

		//拷贝数据
		writtenSize, err := io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("written size %d is not equal to file size %d", writtenSize, fileHeader.Size)

		}

		//重定向
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)

	} else {
		fmt.Println("不支持的文件操作类型")
	}
}

func UploadFileSuccess(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" || r.Method == "POST" {
		io.WriteString(w, "upload success!")
	} else {
		io.WriteString(w, "not supported file operation")
	}
}
