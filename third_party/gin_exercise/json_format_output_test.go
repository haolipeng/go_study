package gin_exercise

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  int         `json:"status"`
}

//普通json格式输出
func TestJsonFormatOutput(t *testing.T) {
	r := gin.Default()
	s1 := Student{
		Name: "haolipeng",
		Age:  32,
	}
	s2 := Student{
		Name: "zhouyang",
		Age:  33,
	}
	var stuList []Student
	stuList = append(stuList, s1, s2)

	r.GET("/users/haolipeng", func(c *gin.Context) {
		c.JSON(http.StatusOK, &Response{
			Data:    stuList,
			Message: "success",
			Status:  http.StatusOK,
		})
	})
	r.Run(":8080")
}

//json数组格式输出
var classmates = []Student{{Name: "haolipeng", Age: 32}, {Name: "pujunlong", Age: 36}, {Name: "lidong", Age: 33}}

func TestJsonArrayOutput(t *testing.T) {
	r := gin.Default()
	r.GET("/users", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, classmates)
	})
	r.Run(":8080")
}
