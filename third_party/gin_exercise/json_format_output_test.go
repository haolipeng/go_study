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

//普通json格式输出
func TestJsonFormatOutput(t *testing.T) {
	r := gin.Default()
	r.GET("/users/haolipeng", func(c *gin.Context) {
		c.JSON(http.StatusOK, Student{
			Name: "haolipeng",
			Age:  32,
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
