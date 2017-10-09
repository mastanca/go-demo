package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
)

var (
    Router *gin.Engine
)

func main() {
    fmt.Println("Hello Demo!")
    Router = gin.Default()
    Router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    Router.Run(":8080")
}

