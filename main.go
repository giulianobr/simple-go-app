package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("web/html/*")
	router.Static("/resources", "./web/resources")
	router.GET("/", index)
	http.Handle("/", router)
	log.Print("Go simple app up and running ...")
	router.Run()
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
