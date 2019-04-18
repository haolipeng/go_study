package handler

import (
	"fmt"
	"go_study/filestore_server/meta"
	"go_study/filestore_server/util"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
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

		//创建文件元结构
		//fix me,Location
		fMeta := meta.FileMeta{
			FileName: fileHeader.Filename,
			Location: fileHeader.Filename,
			UploadAt: time.Now().Format("2006-01-02 15:04:05"),
		}

		//创建缓存文件
		newFile, err := os.Create(fileHeader.Filename)
		if err != nil {
			fmt.Printf("Failed to create file %s\n", err.Error())
			return
		}

		defer newFile.Close()

		//拷贝数据
		fMeta.FileSize, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("written size %d is not equal to file size %d", writtenSize, fileHeader.Size)

		}

		//计算sha1值
		newFile.Seek(0, 0)
		fMeta.FileSha1 = util.FileSha1(newFile)

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
