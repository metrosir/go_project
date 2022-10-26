package main

import (
	"fmt"
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"go_project/pkg/round_robin"
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
		//params, err := c.Params.Get("v")
		//if !err {
		//	c.JSON(400, gin.H{
		//		"message": errors.New("parmas error"),
		//	})
		//}
		//fmt.Println("params:", params)
		//body := request.Q("https://www.baidu.com")
		//if err != nil {
		//	c.JSON(400, gin.H{
		//		"message": errors.New("request error:" + err.Error()),
		//	})
		//}
		fmt.Println(round_robin.CallPHP())
		body, err := round_robin.CallPHP()
		if err != nil {
			c.JSON(200, gin.H{
				"error":   err.Error(),
				"message": body,
			})
		} else {
			c.JSON(200, gin.H{
				"error":   "",
				"message": body,
			})
		}
	})
	r.Run("0.0.0.0:3001")
}
