package main

import (
	"errors"
	"fmt"
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"go_project/pkg/request"
	"io/ioutil"
)

func main() {
	r := gin.Default()
	ginpprof.Wrapper(r)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/api/todo", func(c *gin.Context) {
		params, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(400, gin.H{
				"message": errors.New("parmas error:" + err.Error()),
			})
		}
		fmt.Println("params:", params)
		body, err := request.Q("https://www.baidu.com")
		if err != nil {
			c.JSON(400, gin.H{
				"message": errors.New("request error:" + err.Error()),
			})
		}
		c.JSON(200, gin.H{
			"message": body,
		})
	})
	r.Run("0.0.0.0:3001")
}
