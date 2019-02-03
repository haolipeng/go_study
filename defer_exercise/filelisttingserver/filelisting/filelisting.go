package filelisting

//显示文件列表
import (
	"net/http"
	"os"
	"io/ioutil"
	"strings"
	"fmt"
)

const prefix = "/list/"

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	//字符串中判断是否以prefix开头
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userError(fmt.Sprintf("path %s must start with %s", request.URL.Path, prefix))
	}

	path := request.URL.Path[len("/list/"):]

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)

	return nil
}
