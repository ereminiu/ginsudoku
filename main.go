package main

import (
	"log"

	"github.com/ereminiu/ginsudoku/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	r.GET("/", handlers.HomeHandler)
	r.GET("/entergrid", handlers.EnterGridHandler)
	r.GET("/gengrid", handlers.GenGridHandler)
	r.GET("/getsol", handlers.SolveHandler)
	r.POST("/sendgrid", handlers.ReadHandler)

	log.Println("Server started at port :8080")
	r.Run(":1337")
}
