package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":    "haolipeng",
		"age":     32,
		"address": "chengdu",
	})
}
func handleHello2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":    "zhouyang",
		"age":     33,
		"address": "chengdu",
	})
}

func main() {
	addr := ":9000"
	r1 := gin.Default()
	r1.GET("/hello", handleHello)
	go func() {
		r1.Run(addr) // listen and serve on 0.0.0.0:8080
	}()
	r2 := gin.Default()
	r2.GET("/hello", handleHello2)
	err := r2.Run(":9999")
	if err != nil {
		fmt.Println("server 2 run failed")
	}
}
