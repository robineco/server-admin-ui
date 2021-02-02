package main

import (
	"de.robinscloud.admin-backend/handler"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	fmt.Println("Hey")

	var router = gin.Default()
	router.POST("/status", handler.CheckServiceStatus)

	log.Fatal(router.Run(":4000"))
}
