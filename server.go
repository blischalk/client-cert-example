package main

import (
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.GET("/", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": c.Request.Header,
		})
  })
  v1 := r.Group("/api/v1")
	{
    v1.GET("/foo", func(c *gin.Context) {
      c.JSON(200, gin.H{
        "message": c.Request.Header,
      })
    })
    v1.GET("/bar", func(c *gin.Context) {
      c.JSON(200, gin.H{
        "message": c.Request.Header,
      })
    })
	}
  r.Run()
}
