package main

import (
	"LucasApi/hello/lib"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
    lib.Dbconf()

    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
        "message":"pong",
        })
    })
    fmt.Println("Hello, World")
    r.Run()
}
