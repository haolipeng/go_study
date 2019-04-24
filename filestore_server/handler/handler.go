package handler

import (
	"encoding/json"
	"fmt"
	"go_study/filestore_server/meta"
	"go_study/filestore_server/util"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

//UploadFile:文件上传
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
			fmt.Printf("written size %d is not equal to file size %d", fMeta.FileSize, fileHeader.Size)

		}

		//计算sha1值
		newFile.Seek(0, 0)
		fMeta.FileSha1 = util.FileSha1(newFile)
		fmt.Printf("file sha1 is %s\n", fMeta.FileSha1)

		//添加到文件元集合中
		//meta.UpdateFileMeta(fMeta)
		meta.UpdateFileMeta2DB(fMeta)

		//重定向
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)

	} else {
		fmt.Println("不支持的文件操作类型")
	}
}

//UploadFileSuccess:文件上传成功
func UploadFileSuccess(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" || r.Method == "POST" {
		io.WriteString(w, "upload success!")
	} else {
		io.WriteString(w, "not supported file operation")
	}
}

//QueryFile:文件元信息查询
func QueryFile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fileSha1 := r.Form["filehash"][0]

	fMeta, bExist := meta.GetFileMetaFromDb(fileSha1)
	if !bExist {
		//向http头部写入状态码
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("query file not exist!"))
		return
	}

	//fix me GetFileMetaFromDb may return nil
	data, err := json.Marshal(fMeta)
	if err != nil {
		//向http头部写入状态码
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//返回json后的数据
	w.Write(data)
}

//DownloadFile 下载文件
func DownloadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fileSha1 := r.Form.Get("filehash")
	fMeta, bExist := meta.GetFileMetaFromDb(fileSha1)
	if !bExist {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("file is not exist in fileMeta map"))
		return
	}

	//通过元文件地址来打开文件
	f, err := os.Open(fMeta.Location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	//fix me if file is too big ,ReadAll is not good
	data, err := ioutil.ReadAll(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//添加http头，让浏览器能识别出是文件下载
	w.Header().Set("Content-Type", "application/octect-stream")
	w.Header().Set("content-disposition", "attachment;filename=\""+fMeta.FileName+"\"")

	w.Write(data)
}

//DeleteFile:删除文件
func DeleteFile(w http.ResponseWriter, r *http.Request) {
	//获取要删除的元素
	r.ParseForm()
	fileSha1 := r.Form.Get("filehash")

	//fix me 判断文件是否存在
	//获取文件元信息
	fMeta, bExist := meta.GetFileMetaFromDb(fileSha1)
	if !bExist {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("The File to be deleted is not found in fileMap"))
		return
	}

	localFileExist, _ := util.PathExists(fMeta.Location)
	if false == localFileExist {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("The File to be deleted is not found in local disk"))
		return
	}

	//删除磁盘中文件
	err := os.Remove(fMeta.Location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Delete file failed, in os.Remove function"))
	}

	meta.RemoveFileMeta(fileSha1)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("file have removed!"))
}

func UpdateFileMeta(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	opType := r.Form.Get("op")
	fileSha1 := r.Form.Get("filehash")
	newFileName := r.Form.Get("filename")

	if opType != "update" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	fMeta, bExist := meta.GetFileMetaFromDb(fileSha1)
	if false == bExist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fMeta.FileName = newFileName
	//fix me,location is not change

	//meta.UpdateFileMeta(fMeta)
	_ = meta.UpdateFileMeta2DB(fMeta)

	//return json data to client
	data, err := json.Marshal(fMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
	w.WriteHeader(http.StatusOK)
}
