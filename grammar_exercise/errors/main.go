package main

//在应用程序中，我们使用github.com/pkg/errors处理应用错误
//注意在公共库中，我们一般不使用这个，为什么呢？
//TODO:工程化的错误处理实践并未敲代码实践
import (
	"fmt"
	"os"
)

func main() {
	err := os.ErrExist
	errWrap := fmt.Errorf("...%w...", err) //文件已经存在的错误
	fmt.Println(errWrap)
}
